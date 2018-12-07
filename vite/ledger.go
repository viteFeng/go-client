package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/rpcapi/api"
)

func GetLatestSnapshotChainHash() (*types.Hash, error) {
	client := NewDefaultClient()
	defer client.Close()

	var latestSnapshotHash = &types.Hash{}
	err := client.Call(latestSnapshotHash, "ledger_getLatestSnapshotChainHash", nil)
	if err != nil {
		logging.Error("RPC call error", "rpcApi", "ledger_getLatestSnapshotChainHash", "error", err)
		return nil, fmt.Errorf("ledger_getLatestSnapshotChainHash call error, %v", err)
	}

	return latestSnapshotHash, nil
}

func GetLatestBlock(fd *types.Address) (*api.AccountBlock, error) {
	client := NewDefaultClient()
	defer client.Close()

	var latestBlock = &api.AccountBlock{}
	err := client.Call(latestBlock, "ledger_getLatestBlock", fd)
	if err != nil {
		logging.Error("RPC call error", "rpcApi", "ledger_getLatestBlock", "sendAddress", fd.Hex(), "error", err)
		return nil, fmt.Errorf("ledger_getLatestBlock call error, %v", err)
	}

	logging.Info("The latest block", "latestBlock", latestBlock)

	return latestBlock, nil
}

func GetAccountByAccAddr(fd types.Address) (*RpcAccountInfo, error) {
	client := NewDefaultClient()
	defer client.Close()

	var rpcAccountInfo = &RpcAccountInfo{}
	err := client.Call(rpcAccountInfo, "ledger_getAccountByAccAddr", fd)
	if err != nil {
		logging.Error("RPC call error", "rpcApi", "ledger_getAccountByAccAddr", "sendAddress", fd.Hex(), "error", err)
		return nil, fmt.Errorf("ledger_getAccountByAccAddr call error, %v", err)
	}

	logging.Info("The rpcAccountInfo", "rpcAccountInfo", *rpcAccountInfo)

	return rpcAccountInfo, nil
}

type RpcAccountInfo struct {
	AccountAddress      types.Address `json:"accountAddress"`
	TotalNumber         string        `json:"totalNumber"` // uint64
	TokenBalanceInfoMap map[string]*api.RpcTokenBalanceInfo
}

func GetBlocksByHash(addr types.Address, hash types.Hash, count int64) (*[]api.AccountBlock, error) {
	client := NewDefaultClient()
	defer client.Close()

	var accountBlocks = &[]api.AccountBlock{}
	err := client.Call(accountBlocks, "ledger_getBlocksByHash", addr, hash, count)
	if err != nil {
		logging.Error("RPC call error", "rpcApi", "ledger_getBlocksByHash", "args", addr.Hex(), hash.Hex(), count, "error", err)
		return nil, fmt.Errorf("ledger_getBlocksByHash call error, %v", err)
	}

	logging.Info("The AccountBlocks", "rpcAccountInfo", *accountBlocks)

	return accountBlocks, nil
}
