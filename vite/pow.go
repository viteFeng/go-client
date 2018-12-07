package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite/common/types"
)

func GetPowNonce(address, PreHash *[]byte, difficulty *string) (*[]byte, error) {
	client := NewDefaultClient()
	defer client.Close()

	var nonce = &[]byte{}
	err := client.Call(nonce, "pow_getPowNonce",
		difficulty,
		types.DataHash(append(*address, *PreHash...)),
	)

	if err != nil {
		logging.Error("RPC call error", "rpcApi", "pow_getPowNonce", "error", err)
		return nil, fmt.Errorf("pow_getPowNonce call error, %v", err)
	}

	return nonce, nil
}
