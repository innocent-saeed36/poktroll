package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/pokt-network/poktroll/x/supplier/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper       types.BankKeeper
		sessionKeeper    types.SessionKeeper
		tokenomicsKeeper types.TokenomicsKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
	sessionKeeper types.SessionKeeper,
	tokenomicsKeeper types.TokenomicsKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		bankKeeper:       bankKeeper,
		sessionKeeper:    sessionKeeper,
		tokenomicsKeeper: tokenomicsKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// TODO_TECHDEBT: Evaluate if this is still necessary after the upgrade to cosmos 0.5x
// SupplySessionKeeper assigns the session keeper dependency of this supplier
// keeper. This MUST be done as a separate step from construction because there
// is a circular dependency between the supplier and session keepers.
func (k *Keeper) SupplySessionKeeper(sessionKeeper types.SessionKeeper) {
	k.sessionKeeper = sessionKeeper
}
