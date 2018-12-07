package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/wallet"
	"github.com/vitelabs/go-vite/wallet/entropystore"
)

var WalletManager *wallet.Manager

func init() {
	Setup()
}
func Setup() {
	walletCfg := &wallet.Config{
		DataDir:        setting.ViteSetting.WalletDataDir,
		MaxSearchIndex: setting.ViteSetting.MaxSearchIndex,
	}

	WalletManager = wallet.New(walletCfg)
	WalletManager.Start()

	if setting.ViteSetting.EntropyFile != "" {

		if err := WalletManager.AddEntropyStore(setting.ViteSetting.EntropyFile); err != nil {
			logging.Error(fmt.Sprintf("node.walletManager.AddEntropyStore error: %v", err))
			return
		}

		entropyStoreManager, err := WalletManager.GetEntropyStoreManager(setting.ViteSetting.EntropyFile)

		if err != nil {
			logging.Error(fmt.Sprintf("node.walletManager.GetEntropyStoreManager error: %v", err))
			return
		}

		//unlock
		if err := entropyStoreManager.Unlock(setting.ViteSetting.Password); err != nil {
			logging.Error(fmt.Sprintf("entropyStoreManager.Unlock error: %v", err))
			return
		}

	}
}

func NewMnemonicAndEntropyStore(passphrase string) (mnemonic string, em *entropystore.Manager, err error) {

	return WalletManager.NewMnemonicAndEntropyStore(passphrase)
}

func GetPrimaryAddr() types.Address {
	entropyStoreManager, _ := WalletManager.GetEntropyStoreManager(setting.ViteSetting.EntropyFile)

	return entropyStoreManager.GetPrimaryAddr()
}

func SignData(addr types.Address, data []byte) (signedData, pubkey []byte, err error) {

	entropyStoreManager, _ := WalletManager.GetEntropyStoreManager(setting.ViteSetting.EntropyFile)

	return entropyStoreManager.SignData(addr, data)
}
