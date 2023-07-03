package test_pr

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"trustedStorage/account"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
	"trustedStorage/settings"
	"trustedStorage/transaction"
)

func getBytesFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return make([]byte, byte(0)), err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return make([]byte, byte(0)), err
	}

	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return make([]byte, byte(0)), err
	}
	return bs, nil
}
func TestFillBlockchain() {
	entries, err := os.ReadDir("./certificates")
	if err != nil {
		log.Fatal(err)
	}
	dirPath, _ := os.Getwd()

	for _, e := range entries {
		//?????
		//fmt.Println(e.Name())
		if e.Name() == ".DS_Store" {
			continue
		}

		tempAccount := account.GenAccount()
		doc1Bytes, err := getBytesFromFile(dirPath + "/certificates/" + e.Name())
		if err != nil {
			panic(err)
		}
		tx := transaction.CreateTransaction(tempAccount, doc1Bytes)
		tx = transaction.SignTransaction(tx, tempAccount, 0)

		err = mempool.MemPoolIns.AddTxToMempool(tx)

		if err != nil {
			panic(err)
		}
	}
	numOfTransactionsInBlock := int(settings.GetNumOfTransactionsInBlock())

	for {
		if len(*mempool.MemPoolIns) > numOfTransactionsInBlock {
			block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), mempool.MemPoolIns.FormTransactionsList(numOfTransactionsInBlock))
			err := blockchain.BlockChainIns.AcceptingBlock(&block)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			break
		}
	}

}
