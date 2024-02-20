package config_test

import (
	"testing"

	sdkerrors "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/pokt-network/poktroll/testutil/yaml"
	"github.com/pokt-network/poktroll/x/application/module/config"
	sharedtypes "github.com/pokt-network/poktroll/x/shared/types"
)

func Test_ParseApplicationConfigs(t *testing.T) {
	tests := []struct {
		desc string

		inputConfig string

		expectedError  *sdkerrors.Error
		expectedConfig *config.ApplicationStakeConfig
	}{
		// Valid Configs
		{
			desc: "valid: service staking config",

			inputConfig: `
				stake_amount: 1000upokt
				service_ids:
				  - svc1
				  - svc2
				`,

			expectedError: nil,
			expectedConfig: &config.ApplicationStakeConfig{
				StakeAmount: sdk.NewCoin("upokt", math.NewInt(1000)),
				Services: []*sharedtypes.ApplicationServiceConfig{
					{
						Service: &sharedtypes.Service{Id: "svc1"},
					},
					{
						Service: &sharedtypes.Service{Id: "svc2"},
					},
				},
			},
		},
		// Invalid Configs
		{
			desc: "invalid: empty service staking config",

			inputConfig: ``,

			expectedError: config.ErrApplicationConfigEmptyContent,
		},
		{
			desc: "invalid: no service ids",

			inputConfig: `
				stake_amount: 1000upokt
				service_ids: # explicitly omitting service ids
				`,

			expectedError: config.ErrApplicationConfigInvalidServiceId,
		},
		{
			desc: "invalid: invalid serviceId",

			inputConfig: `
				stake_amount: 1000upokt
				service_ids:
				  - sv c1
				`,

			expectedError: config.ErrApplicationConfigInvalidServiceId,
		},
		{
			desc: "invalid: no stake amount",

			inputConfig: `
				stake_amount: # explicitly omitting stake amount
				service_ids:
				  - svc1
				  - svc2
				`,

			expectedError: config.ErrApplicationConfigInvalidStake,
		},
		{
			desc: "invalid: non-positive stake amount",

			inputConfig: `
				stake_amount: 0upokt
				service_ids:
				  - svc1
				  - svc2
				`,

			expectedError: config.ErrApplicationConfigInvalidStake,
		},
		{
			desc: "invalid: unsupported stake denom",

			inputConfig: `
				stake_amount: 1000npokt
				service_ids:
				  - svc1
				  - svc2
				`,

			expectedError: config.ErrApplicationConfigInvalidStake,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			normalizedConfig := yaml.NormalizeYAMLIndentation(test.inputConfig)
			appServiceConfig, err := config.ParseApplicationConfigs([]byte(normalizedConfig))

			if test.expectedError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, test.expectedError)
				require.Contains(t, err.Error(), test.expectedError.Error())
				require.Nil(t, appServiceConfig)
				return
			}

			require.NoError(t, err)

			require.Equal(t, test.expectedConfig.StakeAmount.Amount, appServiceConfig.StakeAmount.Amount)
			require.Equal(t, test.expectedConfig.StakeAmount.Denom, appServiceConfig.StakeAmount.Denom)

			t.Logf("serviceIds: %v", test.expectedConfig.Services)
			require.Equal(t, len(test.expectedConfig.Services), len(appServiceConfig.Services))
			for i, expected := range test.expectedConfig.Services {
				require.Equal(t, expected.Service.Id, appServiceConfig.Services[i].Service.Id)
			}
		})
	}
}
