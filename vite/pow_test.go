package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/common/types"
	"testing"
)

func TestPow(t *testing.T) {
	setting.Setup("../../conf/app.ini")
	logging.Setup()

	hashes, e := GetLatestSnapshotChainHash()
	fmt.Println(hashes, e)

	addr := "vite_8c982a7513237a984acae15589cea9800274f07585bfd928db"
	address, _ := types.HexToAddress(addr)
	latestBlock, _ := GetLatestBlock(&address)

	var addrByte = address.Bytes()
	var preHashByte = latestBlock.PrevHash.Bytes()
	Nonce, err := GetPowNonce(&addrByte, &preHashByte, &setting.ViteSetting.Difficulty)
	fmt.Println(*Nonce, err)
}
