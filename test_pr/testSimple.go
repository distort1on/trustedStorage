package test_pr

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"trustedStorage/account"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
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

func Test1() {
	blockChain := blockchain.InitBlockchain()
	memPool := mempool.InitMempool()
	dirPath, _ := os.Getwd()

	txDataBase := transaction.TransactionDataBase{
		TxDataBase: make(map[string]transaction.Transaction),
	}
	testAccount := account.GenAccount()

	doc1Bytes, err := getBytesFromFile(dirPath + "/certificates/Certificate_cryptography_Illia_Popov.pdf")
	if err != nil {
		panic(err)
	}
	tx1 := transaction.CreateTransaction(testAccount, doc1Bytes)
	tx1 = transaction.SignTransaction(tx1, testAccount, 0)

	//txDataBase.TxDataBase[string(tx2.TransactionHash)] = tx2
	//fmt.Println(txDataBase.TxDataBase[string(tx2.TransactionHash)].ToString)

	err = memPool.AddTxToMempool(tx1, &txDataBase)
	if err != nil {
		panic(err)
	}

	block := blockchain.CreateBlock(1, blockChain.BlocksHistory[len(blockChain.BlocksHistory)-1].BlockHash, memPool.FormTransactionsList(1))

	err = blockChain.AddBlockToBlockchain(&block, &txDataBase)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(blockChain.ToString())

}
