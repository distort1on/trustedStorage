package mempool

import (
	"errors"
	"trustedStorage/transaction"
)

type Mempool struct {
	MempoolTransactions []transaction.Transaction
}

func InitMempool() *Mempool {
	var m Mempool
	return &m
}

func (m *Mempool) AddTxToMempool(tx transaction.Transaction, txDB *transaction.TransactionDataBase) error {
	if transaction.VerifyTransaction(tx, txDB) {
		m.MempoolTransactions = append(m.MempoolTransactions, tx)
		return nil
	} else {
		return errors.New("tx invalid")
	}

}

func (m *Mempool) FormTransactionsList(numOfTransactions int) []transaction.Transaction {
	var txList []transaction.Transaction

	for i := 0; i < numOfTransactions; i++ {
		txList = append(txList, m.MempoolTransactions[len(m.MempoolTransactions)-1-i])
	}
	m.MempoolTransactions = m.MempoolTransactions[:len(m.MempoolTransactions)-numOfTransactions]

	return txList
}
