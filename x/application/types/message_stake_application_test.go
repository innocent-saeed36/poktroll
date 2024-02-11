package types

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/testutil/sample"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

func TestMsgStakeApplication_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgStakeApplication
		err  error
	}{
		// address related tests
		{
			name: "invalid address - nil stake",
			msg: MsgStakeApplication{
				Address: "invalid_address",
				// Stake explicitly nil
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidAddress,
		},

		// stake related tests
		{
			name: "valid address - nil stake",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				// Stake explicitly nil
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidStake,
		}, {
			name: "valid address - valid stake",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
		}, {
			name: "valid address - zero stake",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(0)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidStake,
		}, {
			name: "valid address - negative stake",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(-100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidStake,
		}, {
			name: "valid address - invalid stake denom",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "invalid", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidStake,
		}, {
			name: "valid address - invalid stake missing denom",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
				},
			},
			err: ErrAppInvalidStake,
		},

		// service related tests
		{
			name: "valid service configs - multiple services",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "svc1"}},
					{Service: &sharedtypes.Service{Id: "svc2"}},
				},
			},
		},
		{
			name: "invalid service configs - not present",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				// Services: omitted
			},
			err: ErrAppInvalidServiceConfigs,
		},
		{
			name: "invalid service configs - empty",
			msg: MsgStakeApplication{
				Address:  sample.AccAddress(),
				Stake:    &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{},
			},
			err: ErrAppInvalidServiceConfigs,
		},
		{
			name: "invalid service configs - invalid service ID that's too long",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "123456790"}},
				},
			},
			err: ErrAppInvalidServiceConfigs,
		},
		{
			name: "invalid service configs - invalid service Name that's too long",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{
						Id:   "123",
						Name: "abcdefghijklmnopqrstuvwxyzab-abcdefghijklmnopqrstuvwxyzab",
					}},
				},
			},
			err: ErrAppInvalidServiceConfigs,
		},
		{
			name: "invalid service configs - invalid service ID that contains invalid characters",
			msg: MsgStakeApplication{
				Address: sample.AccAddress(),
				Stake:   &sdk.Coin{Denom: "upokt", Amount: sdkmath.NewInt(100)},
				Services: []*sharedtypes.ApplicationServiceConfig{
					{Service: &sharedtypes.Service{Id: "12 45 !"}},
				},
			},
			err: ErrAppInvalidServiceConfigs,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
