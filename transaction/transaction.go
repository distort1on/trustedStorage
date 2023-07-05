package transaction

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
	"trustedStorage/account"
	"trustedStorage/keypair"
	"trustedStorage/signature"
)

type Transaction struct {
	SenderAddress []byte //x509 public key hash
	Data          []byte
	PubKey        []byte //x509 public key
	Signature     []byte //asn1 signature
	Cid           []byte //ipfc cid
	Nonce         int64  //unix time
}

func CreateTransaction(sender account.Account, data []byte) Transaction {
	var tx Transaction
	tx.SenderAddress = sender.AccountAddress
	dataHash := sha256.Sum256(data)
	tx.Data = dataHash[:]
	return tx
}

func (tx *Transaction) GetTxHash() []byte {
	txHash := sha256.Sum256(bytes.Join([][]byte{tx.SenderAddress, tx.Data, tx.PubKey, tx.Signature, tx.Cid, []byte(strconv.FormatInt(tx.Nonce, 10))}, []byte{}))
	return txHash[:]
}

func SignTransaction(tx Transaction, sender account.Account, walletIndex uint8) Transaction {
	txDataToSign := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})
	tx.Signature = signature.SignData(sender.Wallet[walletIndex].GetPrivateKey(), txDataToSign)
	tx.PubKey = keypair.EncodePublicKey(sender.Wallet[walletIndex].PubKey)
	tx.Nonce = time.Now().Unix()

	return tx
}

func (tx *Transaction) ToString() string {
	//out, err := json.MarshalIndent(tx, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	//return fmt.Sprint(string(out))

	return fmt.Sprintf(
		"\n\t\t'SenderAddress' : %x\n\t\t'Data' : %x\n\t\t'PubKey' : %x\n\t\t'Signature' : %x\n\t\t'Cid' : %x\n\t\t'Nonce' : %d\n\t\t'Tx hash' : %x\n",
		tx.SenderAddress,
		tx.Data,
		tx.PubKey,
		tx.Signature,
		tx.Cid,
		tx.Nonce,
		tx.GetTxHash(),
	)

}

func (tx *Transaction) ToBytes() []byte {
	txBytes := bytes.Join([][]byte{tx.SenderAddress, tx.Data, tx.PubKey, tx.Signature, tx.Cid, []byte(strconv.FormatInt(tx.Nonce, 10))}, []byte{})
	return txBytes
}
