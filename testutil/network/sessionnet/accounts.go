package sessionnet

//import (
//	"encoding/json"
//	"fmt"
//	"sync/atomic"
//	"testing"
//
//	"cosmossdk.io/math"
//	"github.com/cosmos/cosmos-sdk/client/flags"
//	"github.com/cosmos/cosmos-sdk/crypto/keyring"
//	testcli "github.com/cosmos/cosmos-sdk/testutil/cli"
//	sdknetwork "github.com/cosmos/cosmos-sdk/testutil/network"
//	"github.com/cosmos/cosmos-sdk/types"
//	"github.com/stretchr/testify/require"
//
//	"github.com/pokt-network/poktroll/testutil/network"
//	"github.com/pokt-network/poktroll/testutil/testkeyring"
//	apptypes "github.com/pokt-network/poktroll/x/application/types"
//	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
//	suppliertypes "github.com/pokt-network/poktroll/x/supplier/types"
//)
//
//// createKeyringAccounts populates the Keyring associated with the in-memory
//// network with memnet.numKeyringAccounts() number of pre-generated accounts.
//func (memnet *inMemoryNetworkWithSessions) createKeyringAccounts(t *testing.T) {
//	t.Helper()
//
//	if memnet.Config.Keyring == nil {
//		t.Log("Keyring not initialized, using new in-memory keyring")
//
//		// Construct an in-memory keyring so that it can be populated and used prior
//		// to network start.
//		memnet.Config.Keyring = keyring.NewInMemory(memnet.Config.CosmosCfg.Codec)
//	} else {
//		t.Log("Keyring already initialized, using existing keyring")
//	}
//
//	// Create memnet.NumKeyringAccounts() accounts in the configured keyring.
//	accts := testkeyring.CreatePreGeneratedKeyringAccounts(
//		t, memnet.Config.Keyring, memnet.Config.GetNumKeyringAccounts(t),
//	)
//
//	// Assign the memnet's pre-generated accounts to be a new pre-generated
//	// accounts iterator containing only the accounts which were also created
//	// in the keyring.
//	memnet.PreGeneratedAccounts = testkeyring.NewPreGeneratedAccountIterator(accts...)
//}
//
//func (memnet *inMemoryNetworkWithSessions) createOnChainAccounts(t *testing.T) {
//	net := memnet.GetNetwork(t)
//	supplierGenesisState := network.GetGenesisState[*suppliertypes.GenesisState](t, suppliertypes.ModuleName, memnet)
//	appGenesisState := network.GetGenesisState[*apptypes.GenesisState](t, apptypes.ModuleName, memnet)
//
//	// Initialize all the accounts
//	sequenceIndex := int32(1)
//	initSupplierAccountsWithSequence(
//		t, net,
//		&sequenceIndex,
//		supplierGenesisState.SupplierList...,
//	)
//
//	initAppAccountsWithSequence(
//		t, net,
//		&sequenceIndex,
//		appGenesisState.ApplicationList...,
//	)
//
//	// need to wait for the account to be initialized in the next block
//	require.NoError(t, net.WaitForNextBlock())
//}
//
//// TODO_IN_THIS_COMMIT: integrate with inMemoryNetworkWithSessions?
//func initSupplierAccountsWithSequence(
//	t *testing.T,
//	net *sdknetwork.Network,
//	sequenceIdx *int32,
//	supplierList ...sharedtypes.Supplier,
//) {
//	t.Helper()
//
//	for _, supplier := range supplierList {
//		supplierAddr, err := types.AccAddressFromBech32(supplier.GetAddress())
//		require.NoError(t, err)
//		InitAccountWithSequence(t, net, supplierAddr, int(*sequenceIdx))
//		atomic.AddInt32(sequenceIdx, 1)
//	}
//}
//
//func initAppAccountsWithSequence(
//	t *testing.T,
//	net *sdknetwork.Network,
//	sequenceIdx *int32,
//	appList ...apptypes.Application,
//) {
//	t.Helper()
//
//	for _, application := range appList {
//		appAddr, err := types.AccAddressFromBech32(application.GetAddress())
//		require.NoError(t, err)
//		InitAccountWithSequence(t, net, appAddr, int(*sequenceIdx))
//		atomic.AddInt32(sequenceIdx, 1)
//	}
//}
//
//// InitAccountWithSequence initializes an Account by sending it some funds from
//// the validator in the network to the address provided
//func InitAccountWithSequence(
//	t *testing.T,
//	net *sdknetwork.Network,
//	addr types.AccAddress,
//	signatureSequencerNumber int,
//) {
//	t.Helper()
//
//	val := net.Validators[0]
//	signerAccountNumber := 0
//	ctx := val.ClientCtx
//	args := []string{
//		fmt.Sprintf("--%s=true", flags.FlagOffline),
//		fmt.Sprintf("--%s=%d", flags.FlagAccountNumber, signerAccountNumber),
//		fmt.Sprintf("--%s=%d", flags.FlagSequence, signatureSequencerNumber),
//
//		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
//		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
//		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
//		fmt.Sprintf("--%s=%s", flags.FlagFees, types.NewCoins(types.NewCoin(net.Config.BondDenom, math.NewInt(10))).String()),
//	}
//	amount := types.NewCoins(types.NewCoin("stake", math.NewInt(200)))
//	responseRaw, err := testcli.MsgSendExec(ctx, val.Address, addr, amount, args...)
//	require.NoError(t, err)
//	var responseJSON map[string]interface{}
//	err = json.Unmarshal(responseRaw.Bytes(), &responseJSON)
//	require.NoError(t, err)
//	require.Equal(t, float64(0), responseJSON["code"], "code is not 0 in the response: %v", responseJSON)
//}
