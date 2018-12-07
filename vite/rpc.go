package vite

import (
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/rpc"
)

func NewClient(url string) *rpc.Client {

	client, e := rpc.Dial(url)
	if e != nil {
		panic(e)
	}

	return client
}

func NewDefaultClient() *rpc.Client {
	return NewClient(setting.ViteSetting.HTTPEndpoint)
}
