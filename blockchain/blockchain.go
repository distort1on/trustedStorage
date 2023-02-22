package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"trustedStorage/transaction"
)

type Blockchain struct {
	BlocksHistory []*Block `json:"blocks_history"`
}

type TransactionDataBase struct {
	txDataBase map[string]transaction.Transaction
}

type Block struct {
	//MacigNo
	Version       uint   `json:"version"`
	HashPrevBlock []byte `json:"hash_prev_block"`
	//HashMerkleRoot []byte
	Time         int64                     `json:"time"`
	TxCounter    uint                      `json:"tx_counter"`
	Transactions []transaction.Transaction `json:"transactions"`
	BlockHash    []byte                    `json:"block_hash"`
}

func CreateBlock(ver uint, prevBlockHash []byte, transactions []transaction.Transaction) (b Block) {
	b.Version = ver
	b.HashPrevBlock = prevBlockHash
	b.Time = time.Now().Unix()
	b.Transactions = transactions
	b.TxCounter = uint(len(transactions))

	b.BlockHash = b.GetBlockHash()

	return b
}

func (b *Block) GetBlockHash() []byte {
	blockSumBytes := [][]byte{
		[]byte(strconv.FormatUint(uint64(b.Version), 10)),
		b.HashPrevBlock,
		[]byte(strconv.FormatInt(b.Time, 10)),
		[]byte(strconv.FormatUint(uint64(b.TxCounter), 10)),
	}

	for _, tx := range b.Transactions {
		blockSumBytes = append(blockSumBytes, tx.TransactionHash)
	}

	blockHash := sha256.Sum256(bytes.Join(blockSumBytes, []byte{}))

	return blockHash[:]
}

func InitBlockchain() *Blockchain {
	genesisBlock := Block{
		Version:       1,
		HashPrevBlock: make([]byte, byte(0)),
		Time:          time.Now().Unix(),
		TxCounter:     0, //0 tx??
		Transactions:  []transaction.Transaction{},
	}
	genesisBlock.BlockHash = genesisBlock.GetBlockHash()

	var bc Blockchain

	bc.BlocksHistory = append(bc.BlocksHistory, &genesisBlock)
	return &bc
}

func (b *Block) ToString() string {
	out, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		panic(err)
	}

	return fmt.Sprint(string(out))
}

func (bc *Blockchain) ToString() (s string) {

	for i, block := range bc.BlocksHistory {

		s += fmt.Sprintf("Block on height %v\n", i) + block.ToString() + "\n"
		//s = strings.Join([]string{s}, block.ToString())
	}
	return s
}

func (bc *Blockchain) AddBlockToBlockchain(b *Block, txDB *transaction.TransactionDataBase) error {
	if ValidateBlock(b, bc.BlocksHistory[len(bc.BlocksHistory)-1].BlockHash, txDB) {
		bc.BlocksHistory = append(bc.BlocksHistory, b)
		for _, tx := range txDB.TxDataBase {
			txDB.TxDataBase[string(tx.TransactionHash)] = tx
		}
		return nil
	}

	return errors.New("block cant be add to blockchain")

}

func ValidateBlock(b *Block, lastBlockchainBlockHash []byte, txDB *transaction.TransactionDataBase) bool {

	//verify signature
	//fmt.Printf("%x\n", b.HashPrevBlock)
	//fmt.Printf("%x\n", lastBlockchainBlockHash)
	//fmt.Println(bytes.Equal(b.HashPrevBlock, lastBlockchainBlockHash))

	if !bytes.Equal(b.HashPrevBlock, lastBlockchainBlockHash) {
		return false
	}

	for _, tx := range b.Transactions {
		if !transaction.VerifyTransaction(tx, txDB) {
			return false
		}
	}

	return true

}
