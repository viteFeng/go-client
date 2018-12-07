package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/vitelabs/go-vite-gw/setting"
)

//https://github.com/ethereum/wiki/wiki/JSON-RPC
import (
	"github.com/ethereum/go-ethereum/rpc"
)

func NewClient(url string) *ethclient.Client {

	client, e := ethclient.Dial(url)
	if e != nil {
		panic(e)
	}

	return client
}

func NewDefaultClient() *ethclient.Client {
	return NewClient(setting.EthSetting.HTTPEndpoint)
}

func NewRawClient(url string) *rpc.Client {
	client, e := rpc.Dial(url)
	if e != nil {
		panic(e)
	}

	return client
}

func NewRawDefaultClient() *rpc.Client {
	return NewRawClient(setting.EthSetting.HTTPEndpoint)
}
