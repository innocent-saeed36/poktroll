package types

// DONTCOVER

import sdkerrors "cosmossdk.io/errors"

// x/tokenomics module sentinel errors
var (
	ErrTokenomicsInvalidSigner                 = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrTokenomicsAuthorityAddressInvalid       = sdkerrors.Register(ModuleName, 1101, "the provided authority address is not a valid bech32 address")
	ErrTokenomicsAuthorityAddressMismatch      = sdkerrors.Register(ModuleName, 1102, "the provided authority address does not match the on-chain governance address")
	ErrTokenomicsClaimNil                      = sdkerrors.Register(ModuleName, 1103, "provided claim is nil")
	ErrTokenomicsSessionHeaderNil              = sdkerrors.Register(ModuleName, 1104, "provided claim's session header is nil")
	ErrTokenomicsSessionHeaderInvalid          = sdkerrors.Register(ModuleName, 1105, "provided claim's session header is invalid")
	ErrTokenomicsSupplierModuleMintFailed      = sdkerrors.Register(ModuleName, 1106, "failed to mint uPOKT to supplier module account")
	ErrTokenomicsSupplierRewardFailed          = sdkerrors.Register(ModuleName, 1107, "failed to send uPOKT from supplier module account to supplier")
	ErrTokenomicsSupplierAddressInvalid        = sdkerrors.Register(ModuleName, 1108, "the supplier address in the claim is not a valid bech32 address")
	ErrTokenomicsApplicationNotFound           = sdkerrors.Register(ModuleName, 1109, "application not found")
	ErrTokenomicsApplicationModuleBurn         = sdkerrors.Register(ModuleName, 1110, "failed to burn uPOKT from application module account")
	ErrTokenomicsApplicationModuleFeeFailed    = sdkerrors.Register(ModuleName, 1111, "failed to send uPOKT from application module account to application")
	ErrTokenomicsApplicationUndelegationFailed = sdkerrors.Register(ModuleName, 1112, "failed to undelegate uPOKT from the application module to the application account")
	ErrTokenomicsApplicationAddressInvalid     = sdkerrors.Register(ModuleName, 1113, "the application address in the claim is not a valid bech32 address")
	ErrTokenomicsParamsInvalid                 = sdkerrors.Register(ModuleName, 1114, "provided params are invalid")
	ErrTokenomicsRootHashInvalid               = sdkerrors.Register(ModuleName, 1115, "the root hash in the claim is invalid")
)
