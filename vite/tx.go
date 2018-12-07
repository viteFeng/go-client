package vite

import (
	"fmt"
	"github.com/vitelabs/go-vite-gw/pkg/logging"
	"github.com/vitelabs/go-vite-gw/setting"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/rpcapi/api"
	"strconv"
	"time"
)

func SendRawTx(sendBlock *api.AccountBlock) error {
	client := NewDefaultClient()
	defer client.Close()

	err := client.Call(nil, "tx_sendRawTx", *sendBlock)
	if err != nil {
		logging.Error("send raw tx failed", "method", "SendRawTx", "error", err)
		return fmt.Errorf("send raw tx failed, %v", err)
	} else {
		logging.Info("send raw tx success", "method", "SendRawTx", "txHash", sendBlock.Hash.Hex())
		return nil
	}
}

func SendTransferTx(toAddr, amount string, needPow bool, data []byte) bool {

	sendBlock, err := GetTransferAccountBlock(toAddr, amount, needPow, data)
	if err != nil {
		logging.Error("Gen send block error", "error", err)
		return false
	}

	err = SendRawTx(sendBlock)
	if err != nil {
		logging.Error("Send tx error", "err", err)
		return false
	}
	logging.Info("transfer success", "ToAddress", toAddr, "Amount", amount)
	return true
}

func GetTransferAccountBlock(toAddress, amount string, needPow bool, data []byte) (*api.AccountBlock, error) {
	client := NewDefaultClient()
	defer client.Close()

	latestSnapshotHash, err := GetLatestSnapshotChainHash()
	if err != nil {
		return nil, err
	}

	fd := GetPrimaryAddr()
	latestBlock, err := GetLatestBlock(&fd)
	if err != nil {
		return nil, err
	}

	td, _ := types.HexToAddress(toAddress)
	tokenId, _ := types.HexToTokenTypeId(setting.ViteSetting.TokeTypeId)

	sendBlock := &api.AccountBlock{
		AccountBlock: &ledger.AccountBlock{},
	}
	sendBlock.PrevHash = latestBlock.Hash
	sendBlock.BlockType = ledger.BlockTypeSendCall // send block
	sendBlock.AccountAddress = fd
	sendBlock.FromAddress = fd
	sendBlock.ToAddress = td
	sendBlock.TokenId = tokenId
	sendBlock.SnapshotHash = *latestSnapshotHash
	sendBlock.Timestamp = time.Now().Unix()
	sendBlock.Difficulty = &setting.ViteSetting.Difficulty
	fee := "0"
	sendBlock.Fee = &fee
	sendBlock.Data = data

	if needPow {
		fdBytes := fd.Bytes()
		preHashBytes := latestBlock.Hash.Bytes()
		nonce, err := GetPowNonce(&fdBytes, &preHashBytes, sendBlock.Difficulty)

		if err != nil {
			return nil, err
		}

		sendBlock.Nonce = *nonce
	}

	lbh, _ := strconv.Atoi(latestBlock.Height)
	height := strconv.Itoa(lbh + 1)
	sendBlock.Height = height
	sendBlock.Amount = &amount
	sendBlock.LedgerAccountBlock()

	sendBlock.Hash = sendBlock.ComputeHash()

	signedData, pubK, _ := SignData(fd, sendBlock.Hash.Bytes())

	sendBlock.Signature = signedData
	sendBlock.PublicKey = pubK
	logging.Info("gen send block", "method", "SendRawTx", "block", sendBlock)

	return sendBlock, nil
}
