package test_pr

import (
	"bufio"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
	"log"
	"os"
	"trustedStorage/account"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/mempool"
	"trustedStorage/serialization"
	"trustedStorage/transaction"
)

var blockChain = blockchain.InitBlockchain()
var memPool = mempool.MempoolTransactions{}
var sh = shell.NewShell("localhost:5001")

//var txDataBase = make(transaction.TransactionDataBase)

const numOfTransactionsInBlock = 1

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

func fillMempool() {
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

		err = memPool.AddTxToMempool(tx, doc1Bytes, sh)

		if err != nil {
			panic(err)
		}

	}

	//fmt.Println(memPool.ToString())
}

func Test1() {

	fillMempool()

	for {

		if len(memPool) > numOfTransactionsInBlock {
			block := blockchain.CreateBlock(1, (*blockChain)[len(*blockChain)-1].GetBlockHash(), memPool.FormTransactionsList(numOfTransactionsInBlock))
			err := blockChain.AcceptingBlock(&block)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			break
		}

	}

	fmt.Println("Blockchain: \n" + blockChain.ToString())

	fmt.Println("Mempool: \n" + memPool.ToString())

	//fmt.Println("Transaction Database: \n")
	//for key, value := range txDataBase {
	//
	//	fmt.Printf("\nDocument: %x\n%v\n", key, value.ToString())
	//}

	//save to DB
	bcBytes := serialization.Serialize(&blockChain)
	database.WriteToDB(bcBytes, "blockchain") //maybe write each block?

	mpBytes := serialization.Serialize(&memPool)
	database.WriteToDB(mpBytes, "mempool")

}
