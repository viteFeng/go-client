package vite

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	Setup()

	mnemonic, em, err := NewMnemonicAndEntropyStore("111111")
	fmt.Println(mnemonic, em, err)

}

func TestGetPrimaryAddr(t *testing.T) {
	Setup()

	addr := GetPrimaryAddr()
	fmt.Println(addr)
}
