package mempool

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"trustedStorage/blockchain"
	"trustedStorage/transaction"
)

type MempoolTransactions []transaction.Transaction

var MemPoolIns *MempoolTransactions

func InitMempool() *MempoolTransactions {
	var mp MempoolTransactions
	return &mp
}

func (mp *MempoolTransactions) ToString() (s string) {

	for i, tx := range *mp {

		s += fmt.Sprintf("TRANSACTION №%v\n", i) + tx.ToString() + "\n"
		//s += fmt.Sprintf("TX - %v\n", i) + hex.EncodeToString(tx.Data) + "\n"
	}
	return s
}

func (mp *MempoolTransactions) AddTxToMempool(tx transaction.Transaction) error {

	errTx := blockchain.VerifyTransaction(tx)
	if errTx != nil {
		return errTx
	} else {
		log.Printf("Transaction %x correct and added to mempool", tx.GetTxHash())
	}

	*mp = append(*mp, tx)

	return nil
}

func (mp *MempoolTransactions) FormTransactionsList(numOfTransactions int) []transaction.Transaction {
	var txList []transaction.Transaction
	var rInd int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numOfTransactions; i++ {
		rInd = rand.Intn(len(*mp))
		txList = append(txList, (*mp)[rInd])

		(*mp)[rInd] = (*mp)[len(*mp)-1]
		*mp = (*mp)[:len(*mp)-1]
	}

	return txList
}

func (mp *MempoolTransactions) ReturnTxToMempool(txList []transaction.Transaction) {
	for _, tx := range txList {
		*mp = append(*mp, tx)
	}
}

func (mp *MempoolTransactions) RemoveExistingTxFromMempool() {
	var newMempool MempoolTransactions

	for _, tx := range *mp {
		if !blockchain.CheckTxAlreadyExist(tx) {
			//(*mp)[i] = (*mp)[len(*mp)-1]
			//*mp = (*mp)[:len(*mp)-1]
			newMempool = append(newMempool, tx)
		}
	}
	*mp = newMempool
}
