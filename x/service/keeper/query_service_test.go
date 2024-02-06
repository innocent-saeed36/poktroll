package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/pokt-network/poktroll/testutil/keeper"
	"github.com/pokt-network/poktroll/testutil/nullify"
	"github.com/pokt-network/poktroll/x/service/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestServiceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	msgs := createNService(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetServiceRequest
		response *types.QueryGetServiceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetServiceRequest{
				Index: msgs[0].Id,
			},
			response: &types.QueryGetServiceResponse{Service: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetServiceRequest{
				Index: msgs[1].Id,
			},
			response: &types.QueryGetServiceResponse{Service: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetServiceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Service(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestServiceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ServiceKeeper(t)
	msgs := createNService(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllServiceRequest {
		return &types.QueryAllServiceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ServiceAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Service), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Service),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ServiceAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Service), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Service),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ServiceAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Service),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ServiceAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
