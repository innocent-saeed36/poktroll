//go:generate mockgen -destination=../../testutil/mockcrypto/ring_cache_mock.go -package=mockcrypto . RingCache
package crypto

import (
	"context"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/noot/ring-go"

	"github.com/pokt-network/poktroll/x/service/types"
)

// RingCache is used to store rings used for signing and verifying relay requests.
// It will cache rings for future use after querying the application module for
// the addresses of the gateways the application is delegated to, and converting
// them into their corresponding public key points on the secp256k1 curve.
type RingCache interface {
	RingClient

	// GetCachedAddresses returns the addresses of the applications that are
	// currently cached in the ring cache.
	GetCachedAddresses() []string
	// Start starts the ring cache, it takes a cancellable context and, in a
	// separate goroutine, listens for on-chain delegation events and invalidates
	// the cache if the redelegation event's AppAddress is stored in the cache.
	Start(ctx context.Context)
	// Stop stops the ring cache by unsubscribing from on-chain delegation events.
	// And clears the cache, so that it no longer contains any rings,
	Stop()
}

// RingClient is used to construct rings by querying the application module for
// the addresses of the gateways the application is delegated to, and converting
// them into their corresponding public key points on the secp256k1 curve.
type RingClient interface {
	// GetRingForAddress returns the ring for the given application address if
	// it exists.
	GetRingForAddress(ctx context.Context, appAddress string) (*ring.Ring, error)

	// VerifyRelayRequestSignature verifies the relay request signature against the
	// ring for the application address in the relay request.
	VerifyRelayRequestSignature(ctx context.Context, relayRequest *types.RelayRequest) error
}

// PubKeyClient is used to get the public key of an address and verify a signatures
// against given an address.
// On-chain and off-chain implementations should take care of retrieving the public
// key from the account module and verifying the signature.
type PubKeyClient interface {
	// GetPubKeyFromAddress returns the public key of the given account address if
	// it exists.
	GetPubKeyFromAddress(ctx context.Context, address string) (cryptotypes.PubKey, error)

	// VerifySignature verifies a signature against the signable bytes and the public
	// key of the given address.
	VerifySignature(ctx context.Context, address string, signature []byte, signableBz []byte) error
}
