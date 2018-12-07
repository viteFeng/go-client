package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/p2p"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"testing"
)

func TestRpc(t *testing.T) {
	setting.Setup("../../conf/app.ini")
	//client := NewDefaultClient()
	client := NewClient(setting.ViteSetting.WSEndpoint)

	var info = &api.SyncInfo{}
	client.Call(info, "net_syncInfo", nil)
	fmt.Println("----net_syncInfo:", info)

	var nodeInfo = &p2p.NodeInfo{}
	client.Call(nodeInfo, "net_peers", nil)
	fmt.Println("-----net_peers:", nodeInfo)

	var returns = &api.RpcAccountInfo{}
	client.Call(returns, "ledger_getAccountByAccAddr", "vite_8c982a7513237a984acae15589cea9800274f07585bfd928db")
	fmt.Println("-----ledger_getAccountByAccAddr:", returns)

	wsClient := NewClient(setting.ViteSetting.WSEndpoint)
	fmt.Println("-----client:", wsClient)

}
