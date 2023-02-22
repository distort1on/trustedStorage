package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"trustedStorage/account"
	"trustedStorage/keypair"
	"trustedStorage/signature"
)

type Transaction struct {
	//TxVersion
	//LockTime
	//sequence
	SenderAddress   []byte
	Data            []byte
	PubKey          []byte
	Signature       []byte
	TransactionHash []byte

	//Nonce (Data1 = Data2 sender?)
}

type TransactionDataBase struct {
	TxDataBase map[string]Transaction
}

func CreateTransaction(sender account.Account, data []byte) Transaction {
	var tx Transaction
	tx.SenderAddress = sender.AccountAddress
	dataHash := sha256.Sum256(data)
	tx.Data = dataHash[:]
	return tx
}

func SignTransaction(tx Transaction, sender account.Account, walletIndex uint8) Transaction {
	txDataToSign := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})
	tx.Signature = signature.SignData(sender.Wallet[walletIndex].GetPrivateKey(), txDataToSign)
	tx.PubKey = []byte(keypair.EncodePublicKey(sender.Wallet[walletIndex].PubKey))
	txHash := sha256.Sum256(bytes.Join([][]byte{tx.SenderAddress, tx.Data, tx.Signature, tx.PubKey}, []byte{}))
	tx.TransactionHash = txHash[:]
	return tx
}

func VerifyTransaction(tx Transaction, txDB *TransactionDataBase) bool {
	//check if fields non empty

	if _, inMap := txDB.TxDataBase[string(tx.Data)]; !inMap {
		txDataToVerifySignature := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})
		if signature.VerifySignature(tx.Signature, txDataToVerifySignature, keypair.DecodePublicKey(string(tx.PubKey))) {
			return true
		}
	}

	return false
}

func (tx *Transaction) ToString() string {
	out, err := json.MarshalIndent(tx, "", "\t")
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(string(out))
}
