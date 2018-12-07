package eth

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"math/big"
)

func GetContractEvent(from *big.Int, to *big.Int, addr common.Address) []types.Log {
	client := NewDefaultClient()
	defer client.Close()

	filterQuery := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
		FromBlock: from,
		ToBlock:   to,
	}

	logs, _ := client.FilterLogs(context.Background(), filterQuery)
	return logs
}

func SubscribeContractEvent(from *big.Int, to *big.Int, addr common.Address, f func(types.Log)) (ethereum.Subscription, error) {
	client := NewDefaultClient()
	defer client.Close()

	filterQuery := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
		FromBlock: from,
		ToBlock:   to,
	}

	logChan := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), filterQuery, logChan)
	if err != nil {
		logging.Fatal(err)
		return nil, err
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				logging.Fatal(err)
			case log := <-logChan:
				go f(log)
			}
		}
	}()

	return sub, nil
}
