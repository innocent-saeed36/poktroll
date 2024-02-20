package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/pokt-network/poktroll/x/proof/types"
)

func (k Keeper) AllClaims(ctx context.Context, req *types.QueryAllClaimsRequest) (*types.QueryAllClaimsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	//claimStore := prefix.NewStore(store, types.KeyPrefix(types.ClaimKeyPrefix))

	// isCustomIndex is used to determined if we'll be using the store that points
	// to the actual Claim values, or a secondary index that points to the primary keys.
	var isCustomIndex bool
	var keyPrefix []byte
	switch filter := req.Filter.(type) {
	case *types.QueryAllClaimsRequest_SupplierAddress:
		isCustomIndex = true
		keyPrefix = types.KeyPrefix(types.ClaimSupplierAddressPrefix)
		keyPrefix = append(keyPrefix, []byte(filter.SupplierAddress)...)

	case *types.QueryAllClaimsRequest_SessionEndHeight:
		isCustomIndex = true
		heightBz := make([]byte, 8)
		binary.BigEndian.PutUint64(heightBz, filter.SessionEndHeight)

		keyPrefix = types.KeyPrefix(types.ClaimSessionEndHeightPrefix)
		keyPrefix = append(keyPrefix, heightBz...)

	case *types.QueryAllClaimsRequest_SessionId:
		isCustomIndex = false
		keyPrefix = types.KeyPrefix(types.ClaimPrimaryKeyPrefix)
		keyPrefix = append(keyPrefix, []byte(filter.SessionId)...)

	default:
		isCustomIndex = false
		keyPrefix = types.KeyPrefix(types.ClaimPrimaryKeyPrefix)
	}
	claimStore := prefix.NewStore(store, keyPrefix)

	var claims []types.Claim
	pageRes, err := query.Paginate(claimStore, req.Pagination, func(key []byte, value []byte) error {
		if isCustomIndex {
			// We retrieve the primaryKey, and need to query the actual Claim before decoding it.
			claim, claimFound := k.getClaimByPrimaryKey(ctx, value)
			if claimFound {
				claims = append(claims, claim)
			}
		} else {
			// The value is an encoded Claim.
			var claim types.Claim
			if err := k.cdc.Unmarshal(value, &claim); err != nil {
				return err
			}
			claims = append(claims, claim)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllClaimsResponse{Claims: claims, Pagination: pageRes}, nil
}

func (k Keeper) Claim(ctx context.Context, req *types.QueryGetClaimRequest) (*types.QueryGetClaimResponse, error) {
	if req == nil {
		err := types.ErrProofInvalidQueryRequest.Wrapf("request cannot be nil")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := req.ValidateBasic(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	val, found := k.GetClaim(
		ctx,
		req.SessionId,
		req.SupplierAddress,
	)
	if !found {
		err := types.ErrProofClaimNotFound.Wrapf("session ID %q and supplier %q", req.SessionId, req.SupplierAddress)
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &types.QueryGetClaimResponse{Claim: val}, nil
}