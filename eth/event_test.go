package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/vitelabs/go-vite-gw/setting"
	"math/big"
	"testing"
)

func TestSubscribeContractEvent(t *testing.T) {

	i := big.Int{}
	i.SetUint64(4424177)

	fmt.Println("abi:", Erc2ViteABI)
	contractAddr := common.HexToAddress(setting.EthSetting.ContractAddr)

	event := GetContractEvent(&i, nil, contractAddr)
	for index, value := range event {

		txHash := value.TxHash.Hex()
		data := value.Data
		address := value.Address.Hex()
		blockHash := value.BlockHash.Hex()
		blockNumber := value.BlockNumber
		index2 := value.Index
		removed := value.Removed
		topics := value.Topics
		txIndex := value.TxIndex
		fmt.Println("index:", index)
		fmt.Println("txHash:", txHash)
		fmt.Println("data:", data)
		fmt.Println("address:", address)
		fmt.Println("blockHash:", blockHash)
		fmt.Println("blockNumber:", blockNumber)
		fmt.Println("index2:", index2)
		fmt.Println("removed:", removed)
		fmt.Println("topics:", topics)
		fmt.Println("txIndex:", txIndex)

		for i, v := range topics {
			fmt.Println("topics_inner:", i, ",value:", v.Hex())
			fmt.Println("topics_inner:", i, ",value:", common.BytesToAddress(v.Bytes()).Hex())
			fmt.Println("topics_inner:", i, ",value:")
		}
		fmt.Println("---------")
	}
	sub, _ := SubscribeContractEvent(&i, nil, contractAddr, func(log types.Log) {
		fmt.Println(log)
	})

	//time.Sleep(1 * time.Millisecond)

	fmt.Println(sub)
	sub.Unsubscribe()
}
