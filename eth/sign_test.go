package eth

//
//import (
//	"encoding/hex"
//	"fmt"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/vitelabs/go-vite/common/types"
//	"testing"
//)
//
//func TestSign(t *testing.T) {
//	//eth key
//	key, err := crypto.GenerateKey()
//	if err != nil {
//		fmt.Println("Error: ", err.Error())
//	}
//
//	//eth address
//	eth_address := crypto.PubkeyToAddress(key.PublicKey).Hex()
//	fmt.Printf("address[%d][%v]\n", len(eth_address), eth_address)
//
//	//eth pri
//	privateKey := hex.EncodeToString(key.D.Bytes())
//	fmt.Printf("privateKey[%d][%v]\n", len(privateKey), privateKey)
//
//	//vite address
//	addresses, _, _ := types.CreateAddress()
//	vite_addr := addresses.Hex()
//	fmt.Printf("publicKey[%d][%v]\n", len(vite_addr), vite_addr)
//
//	//source
//	source := eth_address + ":" + vite_addr
//	fmt.Println(source, source[:])
//
//	sourceByte := []byte(source[:])
//	keccak256 := crypto.Keccak256(sourceByte)
//	fmt.Println(keccak256)
//
//	//sign verify
//	signBytes, _ := crypto.Sign(keccak256, key)
//	pubByte, _ := crypto.Ecrecover(keccak256, signBytes)
//	//pubByte, _ := secp256k1.RecoverPubkey(sourceByte, signBytes)
//	signature := crypto.VerifySignature(pubByte, keccak256, signBytes[:64])
//	fmt.Println(signature)
//}
