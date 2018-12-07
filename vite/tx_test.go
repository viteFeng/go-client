package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"testing"
)

func TestRawTx(t *testing.T) {

	setting.Setup("../../conf/app.ini")
	logging.Setup()

	accountBlock := &api.AccountBlock{}
	isSuccess := SendRawTx(accountBlock)

	fmt.Println(isSuccess)
}

func TestTransferTx(t *testing.T) {

	setting.Setup("../../conf/app.ini")
	logging.Setup()
	Setup()

	toAddr := "vite_5112ca7f7ba950f0458dccd6a6b7baf59c32b53665b81cc234"
	isSuccess := SendTransferTx(toAddr, "1"+"000000000000000000", true, []byte("1"))

	fmt.Println(isSuccess)
}
