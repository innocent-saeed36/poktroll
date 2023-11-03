//go:generate mockgen -destination=../../internal/mocks/mockclient/events_query_client_mock.go -package=mockclient . Dialer,Connection,EventsQueryClient
//go:generate mockgen -destination=../../internal/mocks/mockclient/block_client_mock.go -package=mockclient . Block,BlockClient
//go:generate mockgen -destination=../../internal/mocks/mockclient/tx_client_mock.go -package=mockclient . TxContext
//go:generate mockgen -destination=../../internal/mocks/mockclient/cosmos_tx_builder_mock.go -package=mockclient github.com/cosmos/cosmos-sdk/client TxBuilder
//go:generate mockgen -destination=../../internal/mocks/mockclient/cosmos_keyring_mock.go -package=mockclient github.com/cosmos/cosmos-sdk/crypto/keyring Keyring
//go:generate mockgen -destination=../../internal/mocks/mockclient/cosmos_client_mock.go -package=mockclient github.com/cosmos/cosmos-sdk/client AccountRetriever

// Package client defines interfaces and types that facilitate interactions
// with blockchain functionalities, both transactional and observational. It is
// built to provide an abstraction layer for sending, receiving, and querying
// blockchain data, thereby offering a standardized way of integrating with
// various blockchain platforms.
//
// The client package leverages external libraries like cosmos-sdk and cometbft,
// but there is a preference to minimize direct dependencies on these external
// libraries, when defining interfaces, aiming for a cleaner decoupling.
// It seeks to provide a flexible and comprehensive interface layer, adaptable to
// different blockchain configurations and requirements.

package client

import (
	"context"

	comettypes "github.com/cometbft/cometbft/rpc/core/types"
	cosmosclient "github.com/cosmos/cosmos-sdk/client"
	cosmoskeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"pocket/pkg/either"
)

// TxClient provides an interface for initiating transactions on a blockchain.
// It offers a singular, streamlined method to sign and broadcast multiple messages
// as part of a transaction. This interface abstracts the actual sending of a
// transaction, while more intricate functionalities like building, encoding,
// or querying transactions are likely to be handled by other components,
// such as the TxContext.
type TxClient interface {
	SignAndBroadcast(
		ctx context.Context,
		msgs ...cosmostypes.Msg,
	) either.AsyncError
}

// TxContext provides an interface which consolidates the operational dependencies
// required to facilitate the sender side of the transaction lifecycle: build, sign, encode,
// broadcast, query (optional).
//
// TODO_IMPROVE: Avoid depending on cosmos-sdk structs or interfaces; add Pocket
// interface types to substitute:
//   - ResultTx
//   - TxResponse
//   - Keyring
//   - TxBuilder
type TxContext interface {
	// GetKeyring returns the associated key management mechanism for the transaction context.
	GetKeyring() cosmoskeyring.Keyring

	// NewTxBuilder creates and returns a new transaction builder instance.
	NewTxBuilder() cosmosclient.TxBuilder

	// SignTx signs a transaction using the specified key name. It can operate in offline mode,
	// and can overwrite any existing signatures based on the provided flags.
	SignTx(
		keyName string,
		txBuilder cosmosclient.TxBuilder,
		offline, overwriteSig bool,
	) error

	// EncodeTx takes a transaction builder and encodes it, returning its byte representation.
	EncodeTx(txBuilder cosmosclient.TxBuilder) ([]byte, error)

	// BroadcastTx broadcasts the given transaction to the network.
	BroadcastTx(txBytes []byte) (*cosmostypes.TxResponse, error)

	// QueryTx retrieves a transaction status based on its hash and optionally provides
	// proof of the transaction.
	QueryTx(
		ctx context.Context,
		txHash []byte,
		prove bool,
	) (*comettypes.ResultTx, error)
}

// BlockClient is an interface which provides notifications about newly committed
// blocks as well as direct access to the latest block via some blockchain API.
type BlockClient interface {
	// Blocks returns an observable which emits newly committed blocks.
	CommittedBlocksSequence(context.Context) BlocksObservable
	// LatestBlock returns the latest block that has been committed.
	LatestBlock(context.Context) Block
	// Close unsubscribes all observers of the committed blocks sequence observable
	// and closes the events query client.
	Close()
}

// Block is an interface which abstracts the details of a block to its minimal
// necessary components.
type Block interface {
	Height() int64
	Hash() []byte
}

// TODO_CONSIDERATION: the cosmos-sdk CLI code seems to use a cometbft RPC client
// which includes a `#Subscribe()` method for a similar purpose. Perhaps we could
// replace this custom websocket client with that.
// (see: https://github.com/cometbft/cometbft/blob/main/rpc/client/http/http.go#L110)
// (see: https://github.com/cosmos/cosmos-sdk/blob/main/client/rpc/tx.go#L114)
//
// NOTE: a branch which attempts this is available at:
// https://github.com/pokt-network/poktroll/pull/74

// EventsQueryClient is used to subscribe to chain event messages matching the given query,
type EventsQueryClient interface {
	// EventsBytes returns an observable which is notified about chain event messages
	// matching the given query. It receives an either value which contains either an
	// error or the event message bytes.
	EventsBytes(
		ctx context.Context,
		query string,
	) (EventsBytesObservable, error)
	// Close unsubscribes all observers of each active query's events bytes
	// observable and closes the connection.
	Close()
}

// Connection is a transport agnostic, bi-directional, message-passing interface.
type Connection interface {
	// Receive blocks until a message is received or an error occurs.
	Receive() (msg []byte, err error)
	// Send sends a message and may return a synchronous error.
	Send(msg []byte) error
	// Close closes the connection.
	Close() error
}

// Dialer encapsulates the construction of connections.
type Dialer interface {
	// DialContext constructs a connection to the given URL and returns it or
	// potentially a synchronous error.
	DialContext(ctx context.Context, urlStr string) (Connection, error)
}
