package config

import sdkerrors "cosmossdk.io/errors"

var (
	codespace                     = "gatewayconfig"
	ErrGatewayConfigEmptyContent  = sdkerrors.Register(codespace, 2100, "empty gateway staking config content")
	ErrGatewayConfigUnmarshalYAML = sdkerrors.Register(codespace, 2101, "config reader cannot unmarshal yaml content")
	ErrGatewayConfigInvalidStake  = sdkerrors.Register(codespace, 2102, "invalid stake in gateway stake config")
)
