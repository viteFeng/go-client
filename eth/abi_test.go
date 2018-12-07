package eth

import (
	"fmt"
	_ "github.com/vitelabs/go-vite-gw/setting"
	"testing"
)

func TestGetContractAbi(t *testing.T) {
	fmt.Println(Erc2ViteABI)
}
