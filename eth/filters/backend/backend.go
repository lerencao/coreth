package backend

import (
	"context"
	"github.com/ava-labs/coreth/core"
	"github.com/ava-labs/coreth/core/bloombits"
	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/coreth/core/vm"
	"github.com/ava-labs/coreth/ethdb"
	"github.com/ava-labs/coreth/params"
	"github.com/ava-labs/coreth/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

type Backend interface {
	ChainDb() ethdb.Database
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error)
	HeaderByHash(ctx context.Context, blockHash common.Hash) (*types.Header, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)
	GetLogs(ctx context.Context, blockHash common.Hash, number uint64) ([][]*types.Log, error)

	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainAcceptedEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription
	SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription
	SubscribeAcceptedLogsEvent(ch chan<- []*types.Log) event.Subscription

	SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription

	SubscribeAcceptedTransactionEvent(ch chan<- core.NewTxsEvent) event.Subscription

	BloomStatus() (uint64, uint64)
	ServiceFilter(ctx context.Context, session *bloombits.MatcherSession)

	// Added to the backend interface to support limiting of logs requests
	GetVMConfig() *vm.Config
	LastAcceptedBlock() *types.Block
	GetMaxBlocksPerRequest() int64
	ChainConfig() *params.ChainConfig
	CurrentHeader() *types.Header
}
