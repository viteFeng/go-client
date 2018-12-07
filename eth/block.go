package eth

import (
	"github.com/vitelabs/go-vite-gw/types"
	"math/big"
)

func GetBlockNumber() (*big.Int, error) {
	client := NewRawDefaultClient()
	defer client.Close()

	var blockNumber types.Big
	e := client.Call(&blockNumber, "eth_blockNumber")
	if e != nil {
		return nil, e
	}

	return blockNumber.BigInt(), nil
}
