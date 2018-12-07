package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/vitelabs/go-vite-gw/setting"
	"os"
)

var Erc2ViteABI *abi.ABI

func init() {
	contractAbi, e := GetContractAbi(setting.EthSetting.Erc2ViteABIPath)
	if e != nil {
		panic(e)
	}
	Erc2ViteABI = contractAbi
}

//read abi-json file
func GetContractAbi(filePath string) (*abi.ABI, error) {
	file, e := os.Open(filePath)
	if e != nil {
		return nil, e
	}
	defer file.Close()

	abi, e := abi.JSON(file)
	if e != nil {
		return nil, e
	}

	return &abi, nil
}
