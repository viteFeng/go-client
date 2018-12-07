package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/common/types"
	"testing"
)

func TestLedger(t *testing.T) {
	setting.Setup("../../conf/app.ini")
	logging.Setup()

	hashes, e := GetLatestSnapshotChainHash()
	fmt.Println("----hashes:", hashes, e)

	addr := "vite_8c982a7513237a984acae15589cea9800274f07585bfd928db"
	address, _ := types.HexToAddress(addr)
	accountBlock, err := GetLatestBlock(&address)
	fmt.Println("-----accountBlock:", accountBlock, err)

	rpcAccountInfo, err := GetAccountByAccAddr(address)
	fmt.Println("-----rpcAccountInfo:", rpcAccountInfo, err)

	mapp := make(map[string]*string)
	s := "sss"
	mapp["s"] = &s
	fmt.Println(mapp)

	//address, _ = types.HexToAddress("vite_846dda2485a8f8ab1e95a030cb4548d0b7c9a26c6510b41a18")
	address, _ = types.HexToAddress("vite_29a3ec96d1e4a52f50e3119ed9945de09bef1d74a772ee60ff")
	hash, _ := types.HexToHash("8cd902a569a71cf747b9ae53b45498ff2beabef6cfb5de1d2f6730142a3beb72")

	fmt.Println(fmt.Sprint(111))

	block, err := GetBlocksByHash(address, hash, 15)
	fmt.Println("-----GetBlocksByHash:", block, err)
	for key, value := range *block {
		fmt.Println(key,
			",tx hash:", value.Hash.Hex(),
			",pre hash:", value.PrevHash.Hex(),
			",type:", value.BlockType,
			",data:", value.Data,
			"block:", value)
	}

}
