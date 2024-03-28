//go:build e2e

package e2e

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"cosmossdk.io/depinject"
	sdklog "cosmossdk.io/log"
	cometcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/app"
	"github.com/pokt-network/poktroll/testutil/testclient"
	apptypes "github.com/pokt-network/poktroll/x/application/types"
	prooftypes "github.com/pokt-network/poktroll/x/proof/types"
	sessiontypes "github.com/pokt-network/poktroll/x/session/types"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
)

var (
	addrRe   *regexp.Regexp
	amountRe *regexp.Regexp

	accNameToAddrMap     = make(map[string]string)
	accAddrToNameMap     = make(map[string]string)
	accNameToAppMap      = make(map[string]apptypes.Application)
	accNameToSupplierMap = make(map[string]sharedtypes.Supplier)

	flagFeaturesPath string
	keyRingFlag      = "--keyring-backend=test"
	chainIdFlag      = "--chain-id=poktroll"
	appGateServerUrl = "http://localhost:42069" // Keeping localhost by default because that is how we run the tests on our machines locally
)

func init() {
	addrRe = regexp.MustCompile(`address:\s+(\S+)\s+name:\s+(\S+)`)
	amountRe = regexp.MustCompile(`amount:\s+"(.+?)"\s+denom:\s+upokt`)

	flag.StringVar(&flagFeaturesPath, "features-path", "*.feature", "Specifies glob paths for the runner to look up .feature files")

	// If "APPGATE_SERVER_URL" envar is present, use it for appGateServerUrl
	if url := os.Getenv("APPGATE_SERVER_URL"); url != "" {
		appGateServerUrl = url
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	log.Printf("features path: %s", flagFeaturesPath)
	m.Run()
}

type suite struct {
	gocuke.TestingT
	// TODO_TECHDEBT: rename to `poktrolld`.
	pocketd          *pocketdBin
	scenarioState    map[string]any // temporary state for each scenario
	cdc              codec.Codec
	proofQueryClient prooftypes.QueryClient
}

func (s *suite) Before() {
	s.pocketd = new(pocketdBin)
	s.scenarioState = make(map[string]any)
	deps := depinject.Configs(
		app.AppConfig(),
		depinject.Supply(
			sdklog.NewTestLogger(s),
		),
	)
	err := depinject.Inject(deps, &s.cdc)
	require.NoError(s, err)
	s.buildAddrMap()
	s.buildAppMap()
	s.buildSupplierMap()

	flagSet := testclient.NewLocalnetFlagSet(s)
	clientCtx := testclient.NewLocalnetClientCtx(s, flagSet)
	s.proofQueryClient = prooftypes.NewQueryClient(clientCtx)
}

// TestFeatures runs the e2e tests specified in any .features files in this directory
// * This test suite assumes that a LocalNet is running
func TestFeatures(t *testing.T) {
	gocuke.NewRunner(t, &suite{}).Path(flagFeaturesPath).Run()
}

// TODO_TECHDEBT: rename `pocketd` to `poktrolld`.
func (s *suite) TheUserHasThePocketdBinaryInstalled() {
	s.TheUserRunsTheCommand("help")
}

func (s *suite) ThePocketdBinaryShouldExitWithoutError() {
	require.NoError(s, s.pocketd.result.Err)
}

func (s *suite) TheUserRunsTheCommand(cmd string) {
	cmds := strings.Split(cmd, " ")
	res, err := s.pocketd.RunCommand(cmds...)
	s.pocketd.result = res
	if err != nil {
		s.Fatalf("ERROR: error running command %s: %s", cmd, err)
	}
}

func (s *suite) TheUserShouldBeAbleToSeeStandardOutputContaining(arg1 string) {
	if !strings.Contains(s.pocketd.result.Stdout, arg1) {
		s.Fatalf("ERROR: stdout must contain '%s' but instead contains: '%s'", arg1, s.pocketd.result.Stdout)
	}
}

func (s *suite) TheUserSendsUpoktFromAccountToAccount(amount int64, accName1, accName2 string) {
	args := []string{
		"tx",
		"bank",
		"send",
		accNameToAddrMap[accName1],
		accNameToAddrMap[accName2],
		fmt.Sprintf("%dupokt", amount),
		keyRingFlag,
		chainIdFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("ERROR: error sending upokt: %s", err)
	}
	s.pocketd.result = res
}

func (s *suite) TheAccountHasABalanceGreaterThanUpokt(accName string, amount int64) {
	bal := s.getAccBalance(accName)
	if bal < int(amount) {
		s.Fatalf("ERROR: account %s does not have enough upokt: %d < %d", accName, bal, amount)
	}
	s.scenarioState[accBalanceKey(accName)] = bal // save the balance for later
}

func (s *suite) AnAccountExistsFor(accName string) {
	bal := s.getAccBalance(accName)
	s.scenarioState[accBalanceKey(accName)] = bal // save the balance for later
}

func (s *suite) TheStakeOfShouldBeUpoktThanBefore(actorType string, accName string, expectedStakeChange int64, condition string) {
	// Get previous stake
	stakeKey := accStakeKey(actorType, accName)
	prevStakeAny, ok := s.scenarioState[stakeKey]
	if !ok {
		s.Fatalf("ERROR: no previous stake found for %s", accName)
	}
	prevStake, ok := prevStakeAny.(int)
	if !ok {
		s.Fatalf("ERROR: previous stake for %s is not an int", accName)
	}

	// Get current stake
	currStake, ok := s.getStakedAmount(actorType, accName)
	if !ok {
		s.Fatalf("ERROR: no current stake found for %s", accName)
	}
	s.scenarioState[stakeKey] = currStake // save the stake for later

	// Validate the change in stake
	s.validateAmountChange(prevStake, currStake, expectedStakeChange, accName, condition)
}

func (s *suite) TheAccountBalanceOfShouldBeUpoktThanBefore(accName string, expectedStakeChange int64, condition string) {
	// Get previous balance
	balanceKey := accBalanceKey(accName)
	prevBalanceAny, ok := s.scenarioState[balanceKey]
	if !ok {
		s.Fatalf("ERROR: no previous balance found for %s", accName)
	}
	prevBalance, ok := prevBalanceAny.(int)
	if !ok {
		s.Fatalf("ERROR: previous balance for %s is not an int", accName)
	}

	// Get current balance
	currBalance := s.getAccBalance(accName)
	s.scenarioState[balanceKey] = currBalance // save the balance for later

	// Validate the change in stake
	s.validateAmountChange(prevBalance, currBalance, expectedStakeChange, accName, condition)
}

func (s *suite) TheUserShouldWaitForSeconds(dur int64) {
	time.Sleep(time.Duration(dur) * time.Second)
}

func (s *suite) TheUserStakesAWithUpoktFromTheAccount(actorType string, amount int64, accName string) {
	// Create a temporary config file
	configPathPattern := fmt.Sprintf("%s_stake_config_*.yaml", accName)
	configContent := fmt.Sprintf(`stake_amount: %d upokt`, amount)
	configFile, err := ioutil.TempFile("", configPathPattern)
	if err != nil {
		s.Fatalf("ERROR: error creating config file: %q", err)
	}
	if _, err = configFile.Write([]byte(configContent)); err != nil {
		s.Fatalf("ERROR: error writing config file: %q", err)
	}

	args := []string{
		"tx",
		actorType,
		fmt.Sprintf("stake-%s", actorType),
		"--config",
		configFile.Name(),
		"--from",
		accName,
		keyRingFlag,
		chainIdFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)

	// Remove the temporary config file
	err = os.Remove(configFile.Name())
	if err != nil {
		s.Fatalf("ERROR: error removing config file: %q", err)
	}

	if err != nil {
		s.Fatalf("ERROR: error staking %s: %s", actorType, err)
	}
	s.pocketd.result = res
}

func (s *suite) TheUserUnstakesAFromTheAccount(actorType string, accName string) {
	args := []string{
		"tx",
		actorType,
		fmt.Sprintf("unstake-%s", actorType),
		"--from",
		accName,
		keyRingFlag,
		chainIdFlag,
		"-y",
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("ERROR: error unstaking %s: %s", actorType, err)
	}
	s.pocketd.result = res
}

func (s *suite) TheAccountForIsStaked(actorType, accName string) {
	stakeAmount, found := s.getStakedAmount(actorType, accName)
	if !found {
		s.Fatalf("ERROR: account %s should be staked", accName)
	}
	s.scenarioState[accStakeKey(actorType, accName)] = stakeAmount // save the stakeAmount for later
}

func (s *suite) TheForAccountIsNotStaked(actorType, accName string) {
	_, found := s.getStakedAmount(actorType, accName)
	if found {
		s.Fatalf("ERROR: account %s should not be staked", accName)
	}
}

func (s *suite) TheForAccountIsStakedWithUpokt(actorType, accName string, amount int64) {
	stakeAmount, found := s.getStakedAmount(actorType, accName)
	if !found {
		s.Fatalf("ERROR: account %s should be staked", accName)
	}
	if int64(stakeAmount) != amount {
		s.Fatalf("ERROR: account %s stake amount is not %d", accName, amount)
	}
	s.scenarioState[accStakeKey(actorType, accName)] = stakeAmount // save the stakeAmount for later
}

func (s *suite) TheApplicationIsStakedForService(appName string, serviceId string) {
	for _, serviceConfig := range accNameToAppMap[appName].ServiceConfigs {
		if serviceConfig.Service.Id == serviceId {
			return
		}
	}
	s.Fatalf("ERROR: application %s is not staked for service %s", appName, serviceId)
}

func (s *suite) TheSupplierIsStakedForService(supplierName string, serviceId string) {
	for _, serviceConfig := range accNameToSupplierMap[supplierName].Services {
		if serviceConfig.Service.Id == serviceId {
			return
		}
	}
	s.Fatalf("ERROR: supplier %s is not staked for service %s", supplierName, serviceId)
}

func (s *suite) TheSessionForApplicationAndServiceContainsTheSupplier(appName string, serviceId string, supplierName string) {
	app, found := accNameToAppMap[appName]
	if !found {
		s.Fatalf("ERROR: application %s not found", appName)
	}
	expectedSupplier, found := accNameToSupplierMap[supplierName]
	if !found {
		s.Fatalf("ERROR: supplier %s not found", supplierName)
	}
	argsAndFlags := []string{
		"query",
		"session",
		"get-session",
		app.Address,
		serviceId,
		fmt.Sprintf("--%s=json", cometcli.OutputFlag),
	}
	res, err := s.pocketd.RunCommandOnHost("", argsAndFlags...)
	if err != nil {
		s.Fatalf("ERROR: error getting session for app %s and service %s: %s", appName, serviceId, err)
	}
	var resp sessiontypes.QueryGetSessionResponse
	responseBz := []byte(strings.TrimSpace(res.Stdout))
	s.cdc.MustUnmarshalJSON(responseBz, &resp)
	for _, supplier := range resp.Session.Suppliers {
		if supplier.Address == expectedSupplier.Address {
			return
		}
	}
	s.Fatalf("ERROR: session for app %s and service %s does not contain supplier %s", appName, serviceId, supplierName)
}

func (s *suite) TheApplicationSendsTheSupplierARequestForServiceWithData(appName, supplierName, serviceId, requestData string) {
	res, err := s.pocketd.RunCurl(appGateServerUrl, serviceId, requestData)
	if err != nil {
		s.Fatalf("ERROR: error sending relay request from app %s to supplier %s for service %s: %v", appName, supplierName, serviceId, err)
	}

	relayKey := relayReferenceKey(appName, supplierName)
	s.scenarioState[relayKey] = res.Stdout
}

func (s *suite) TheApplicationReceivesASuccessfulRelayResponseSignedBy(appName string, supplierName string) {
	relayKey := relayReferenceKey(appName, supplierName)
	stdout, ok := s.scenarioState[relayKey]
	require.Truef(s, ok, "no relay response found for %s", relayKey)
	require.Contains(s, stdout, `"result":"0x`)
}

func (s *suite) getStakedAmount(actorType, accName string) (int, bool) {
	s.Helper()
	args := []string{
		"query",
		actorType,
		fmt.Sprintf("list-%s", actorType),
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("ERROR: error getting %s: %s", actorType, err)
	}
	s.pocketd.result = res

	found := strings.Contains(res.Stdout, accNameToAddrMap[accName])
	amount := 0
	if found {
		escapedAddress := accNameToAddrMap[accName]
		re := regexp.MustCompile(`(?s)address: ([\w\d]+).*?stake:\s*amount: "(\d+)"`)
		matches := re.FindAllStringSubmatch(res.Stdout, -1)

		if len(matches) < 2 {
			s.Fatalf("ERROR: no stake amount found for %s", accName)
		}
		for _, match := range matches {
			if match[1] == escapedAddress {
				amount, err = strconv.Atoi(match[2])
				require.NoError(s, err)
				return amount, true
			}
		}
	}
	return 0, false
}

func (s *suite) buildAddrMap() {
	s.Helper()
	res, err := s.pocketd.RunCommand(
		"keys", "list", keyRingFlag,
	)
	if err != nil {
		s.Fatalf("ERROR: error getting keys: %s", err)
	}
	s.pocketd.result = res
	matches := addrRe.FindAllStringSubmatch(res.Stdout, -1)
	for _, match := range matches {
		name := match[2]
		address := match[1]
		accNameToAddrMap[name] = address
		accAddrToNameMap[address] = name
	}
}

func (s *suite) buildAppMap() {
	s.Helper()
	argsAndFlags := []string{
		"query",
		"application",
		"list-application",
		fmt.Sprintf("--%s=json", cometcli.OutputFlag),
	}
	res, err := s.pocketd.RunCommandOnHost("", argsAndFlags...)
	if err != nil {
		s.Fatalf("ERROR: error getting application list: %s", err)
	}
	s.pocketd.result = res
	var resp apptypes.QueryAllApplicationsResponse
	responseBz := []byte(strings.TrimSpace(res.Stdout))
	s.cdc.MustUnmarshalJSON(responseBz, &resp)
	for _, app := range resp.Applications {
		accNameToAppMap[accAddrToNameMap[app.Address]] = app
	}
}

func (s *suite) buildSupplierMap() {
	s.Helper()
	argsAndFlags := []string{
		"query",
		"supplier",
		"list-supplier",
		fmt.Sprintf("--%s=json", cometcli.OutputFlag),
	}
	res, err := s.pocketd.RunCommandOnHost("", argsAndFlags...)
	if err != nil {
		s.Fatalf("ERROR: error getting supplier list: %s", err)
	}
	s.pocketd.result = res
	var resp suppliertypes.QueryAllSuppliersResponse
	responseBz := []byte(strings.TrimSpace(res.Stdout))
	s.cdc.MustUnmarshalJSON(responseBz, &resp)
	for _, supplier := range resp.Supplier {
		accNameToSupplierMap[accAddrToNameMap[supplier.Address]] = supplier
	}
}

func (s *suite) getAccBalance(accName string) int {
	s.Helper()
	args := []string{
		"query",
		"bank",
		"balances",
		accNameToAddrMap[accName],
	}
	res, err := s.pocketd.RunCommandOnHost("", args...)
	if err != nil {
		s.Fatalf("ERROR: error getting balance: %s", err)
	}
	s.pocketd.result = res
	match := amountRe.FindStringSubmatch(res.Stdout)
	if len(match) < 2 {
		s.Fatalf("ERROR: no balance found for %s", accName)
	}
	accBalance, err := strconv.Atoi(match[1])
	require.NoError(s, err)
	return accBalance
}

// validateAmountChange validates if the balance of an account has increased or decreased by the expected amount
func (s *suite) validateAmountChange(prevAmount, currAmount int, expectedAmountChange int64, accName, condition string) {
	deltaAmount := int64(math.Abs(float64(currAmount - prevAmount)))
	// Verify if balance is more or less than before
	switch condition {
	case "more":
		if currAmount <= prevAmount {
			s.Fatalf("ERROR: account %s expected to have more upokt but: %d <= %d", accName, currAmount, prevAmount)
		}
		if deltaAmount != expectedAmountChange {
			s.Fatalf("ERROR: account %s balance expected to increase by %d, but actually increased by %d", accName, expectedAmountChange, deltaAmount)
		}
	case "less":
		if currAmount >= prevAmount {
			s.Fatalf("ERROR: account %s expected to have less upokt but: %d >= %d", accName, currAmount, prevAmount)
		}
		if deltaAmount != expectedAmountChange {
			s.Fatalf("ERROR: account %s balance expected to decrease by %d, but actually decreased by %d", accName, expectedAmountChange, deltaAmount)
		}
	default:
		s.Fatalf("ERROR: unknown condition %s", condition)
	}

}

// TODO_IMPROVE: use `sessionId` and `supplierName` since those are the two values
// used to create the primary composite key on-chain to uniquely distinguish relays.
func relayReferenceKey(appName, supplierName string) string {
	return fmt.Sprintf("%s/%s", appName, supplierName)
}

// accBalanceKey is a helper function to create a key to store the balance
// for accName in the context of a scenario state.
func accBalanceKey(accName string) string {
	return fmt.Sprintf("balance/%s", accName)
}

// accStakeKey is a helper function to create a key to store the stake
// for accName of type actorType in the context of a scenario state.
func accStakeKey(actorType, accName string) string {
	return fmt.Sprintf("stake/%s/%s", actorType, accName)
}
