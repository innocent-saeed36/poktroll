package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/pokt-network/poktroll/x/session/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		appKeeper      types.ApplicationKeeper
		supplierKeeper types.SupplierKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	appKeeper types.ApplicationKeeper,
	supplierKeeper types.SupplierKeeper,

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

		appKeeper:      appKeeper,
		supplierKeeper: supplierKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// BeginBlocker is called at the beginning of every block.
// It fetches the block hash from the committed block ans saves its hash
// in the store.
func (k Keeper) BeginBlocker(ctx sdk.Context) {
	// ctx.BlockHeader().LastBlockId.Hash is the hash of the last block committed
	hash := ctx.BlockHeader().LastBlockId.Hash
	// ctx.BlockHeader().Height is the height of the last committed block.
	height := ctx.BlockHeader().Height
	// Block height 1 is the first committed block which uses `genesis.json` as its parent.
	// See the explanation here for more details: https://github.com/pokt-network/poktroll/issues/377#issuecomment-1936607294
	// Fallback to an empty byte slice during the genesis block.
	if height == 1 {
		hash = []byte{}
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(GetBlockHashKey(height), hash)
}
