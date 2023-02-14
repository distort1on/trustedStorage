package main

import (
	"encoding/json"
	"trustedStorage/account"
	"trustedStorage/signature"
	"trustedStorage/transaction"

	"fmt"
)

func main() {
	//fmt.Printf("%x", keypair.SHA256([]byte("test")))

	testAccount := account.GenAccount()

	//sig := signature.SignData(testAccount.Wallet[0].GetPrivateKey(), []byte("hello"))

	//res := signature.VerifySignature(sig, []byte("hello"), testAccount.Wallet[0].PubKey)

	//signature.PrintSignature(sig)
	//fmt.Println(res)

	fmt.Printf("AccountAddress: %x\n", testAccount.AccountAddress)

	tx := transaction.CreateTransaction(testAccount, []byte("hello"))
	tx = transaction.SignTransaction(tx, testAccount, 0)

	txRaw, _ := json.Marshal(tx)
	res := signature.VerifySignature(tx.Signature, txRaw, testAccount.Wallet[0].PubKey)
	tx.Print()

	fmt.Println(res)

}
