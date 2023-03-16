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
	//timestamp here?
	SenderAddress []byte //x509 public key hash
	Data          []byte
	PubKey        []byte //x509 public key
	Signature     []byte //asn1 signature
	Cid           []byte
	//Nonce (Data1 = Data2 sender?)
}

func CreateTransaction(sender account.Account, data []byte) Transaction {
	var tx Transaction
	tx.SenderAddress = sender.AccountAddress
	dataHash := sha256.Sum256(data)
	tx.Data = dataHash[:]
	return tx
}

func (tx *Transaction) GetTxHash() []byte {
	txHash := sha256.Sum256(bytes.Join([][]byte{tx.SenderAddress, tx.Data, tx.PubKey, tx.Signature}, []byte{}))
	return txHash[:]
}

func SignTransaction(tx Transaction, sender account.Account, walletIndex uint8) Transaction {
	txDataToSign := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})
	tx.Signature = signature.SignData(sender.Wallet[walletIndex].GetPrivateKey(), txDataToSign)
	tx.PubKey = keypair.EncodePublicKey(sender.Wallet[walletIndex].PubKey)

	return tx
}

func VerifyTransaction(tx Transaction) bool {
	//todo check if fields non empty

	txDataToVerifySignature := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})
	if signature.VerifySignature(tx.Signature, txDataToVerifySignature, keypair.DecodePublicKey(tx.PubKey)) {
		return true
	}

	return false
}

func (tx *Transaction) ToString() string {
	out, err := json.MarshalIndent(tx, "", "\t")
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(string(out))

	//return fmt.Sprintf(
	//	"\nSenderAddress: %x\nData: %x\nPubKey: %x\nSignature: %x\nTransactionHash: %x\n",
	//	tx.SenderAddress,
	//	tx.Data,
	//	tx.PubKey,
	//	tx.Signature,
	//	tx.TransactionHash,
	//)

}
