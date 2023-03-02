package mempool

import (
	"errors"
	"fmt"
	"trustedStorage/transaction"
)

type MempoolTransactions []transaction.Transaction

func (mp *MempoolTransactions) ToString() (s string) {

	for i, tx := range *mp {

		s += fmt.Sprintf("TRANSACTION №%v\n", i) + tx.ToString() + "\n"
		//s += fmt.Sprintf("TX - %v\n", i) + hex.EncodeToString(tx.Data) + "\n"
	}
	return s
}

func (m *MempoolTransactions) AddTxToMempool(tx transaction.Transaction, txDB *transaction.TransactionDataBase) error {
	if transaction.VerifyTransaction(tx, txDB) {
		*m = append(*m, tx)
		return nil
	} else {
		return errors.New("tx invalid")
	}

}

func (m *MempoolTransactions) FormTransactionsList(numOfTransactions int) []transaction.Transaction {
	var txList []transaction.Transaction

	for i := 0; i < numOfTransactions; i++ {
		txList = append(txList, (*m)[len(*m)-1-i])
	}
	*m = (*m)[:len(*m)-numOfTransactions]

	return txList
}
