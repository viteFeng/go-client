package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vitelabs/go-vite-gw/setting"
	"testing"
)

func TestRpc(t *testing.T) {
	client := NewClient(setting.EthSetting.HTTPEndpoint)
	defer client.Close()
	fmt.Println(client)

	id, _ := client.NetworkID(context.Background())
	fmt.Println("NetworkID:", id)

	progress, _ := client.SyncProgress(context.Background())
	fmt.Println("SyncProgress:", progress)

	pendingBalanceAt, _ := client.PendingBalanceAt(context.Background(), common.HexToAddress("0x76C2d2DD88d0e8c3199298008f53a8a7DD6af655"))
	balanceAt, _ := client.BalanceAt(context.Background(), common.HexToAddress("0x76C2d2DD88d0e8c3199298008f53a8a7DD6af655"), nil)
	fmt.Println("PendingBalanceAt:", pendingBalanceAt, ",balanceAt:", balanceAt)

	fmt.Println("--------------")

	rawClient := NewRawDefaultClient()
	defer rawClient.Close()

	fmt.Println(GetBlockNumber())

}
