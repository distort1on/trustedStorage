package transaction

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"trustedStorage/account"
	"trustedStorage/signature"
)

type Transaction struct {
	TransactionID []byte `json:",omitempty"`
	SenderAddress []byte
	Data          []byte
	Signature     []byte `json:"-"`
	//Nonce (Data1 = Data2 sender?)
}

func CreateTransaction(sender account.Account, data []byte) Transaction {
	var tx Transaction
	tx.SenderAddress = sender.AccountAddress

	dataHash := sha256.Sum256(data)
	tx.Data = dataHash[:]

	raw, _ := json.Marshal(tx)
	txIDHash := sha256.Sum256(raw)
	tx.TransactionID = txIDHash[:]
	return tx
}

func SignTransaction(tx Transaction, sender account.Account, walletIndex uint8) Transaction {
	raw, _ := json.Marshal(tx)
	//fmt.Printf("%s", raw)
	sig := signature.SignData(sender.Wallet[walletIndex].GetPrivateKey(), raw)
	tx.Signature = sig
	return tx
}

func (tx Transaction) Print() {
	fmt.Printf("TransactionID: %x\nSenderAddress: %x\nData: %x\nSignature: %x\n", tx.TransactionID, tx.SenderAddress, tx.Data, tx.Signature)
}
