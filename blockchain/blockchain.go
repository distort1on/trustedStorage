package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"trustedStorage/size"
	"trustedStorage/transaction"
)

type Blockchain []*Block

type BlockHeader struct {
	Version       uint
	HashPrevBlock []byte
	//HashMerkleRoot []byte
	Time int64
}

type Block struct {
	//MacigNo
	BlockSize uint
	BlockHeader
	TxCounter    uint
	Transactions []transaction.Transaction
}

func CreateBlock(ver uint, prevBlockHash []byte, transactions []transaction.Transaction) (b Block) {
	b.Version = ver
	b.HashPrevBlock = prevBlockHash
	b.Time = time.Now().Unix()
	b.Transactions = transactions
	b.TxCounter = uint(len(transactions))
	b.BlockSize = uint(size.Of(b))

	return b
}

func (b *Block) GetBlockHash() []byte {
	blockSumBytes := [][]byte{
		[]byte(strconv.FormatUint(uint64(b.Version), 10)),
		b.HashPrevBlock,
		[]byte(strconv.FormatInt(b.Time, 10)),
		[]byte(strconv.FormatUint(uint64(b.BlockSize), 10)),
		[]byte(strconv.FormatUint(uint64(b.TxCounter), 10)),
	}

	for _, tx := range b.Transactions {
		blockSumBytes = append(blockSumBytes, tx.GetTxHash())
	}

	blockHash := sha256.Sum256(bytes.Join(blockSumBytes, []byte{}))

	return blockHash[:]
}

func InitBlockchain() *Blockchain {
	genesisBlock := Block{
		BlockHeader: BlockHeader{
			Version:       1,
			HashPrevBlock: make([]byte, byte(0)),
			Time:          time.Now().Unix(),
		},
		TxCounter:    0, //0 tx??
		Transactions: []transaction.Transaction{},
	}
	genesisBlock.BlockSize = uint(size.Of(genesisBlock))

	var bc Blockchain

	bc = append(bc, &genesisBlock)
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

	for i, block := range *bc {

		s += fmt.Sprintf("\nBlock on height %v\n", i) + block.ToString() + "\n"
		//s = strings.Join([]string{s}, block.ToString())
	}
	return s
}

func (bc *Blockchain) AcceptingBlock(b *Block) error {
	if VerificationBlock(b, (*bc)[len(*bc)-1].GetBlockHash()) {
		*bc = append(*bc, b)
		return nil
	}
	return errors.New("block cant be add to blockchain")

}

func VerificationBlock(b *Block, lastBlockchainBlockHash []byte) bool {
	if !bytes.Equal(b.HashPrevBlock, lastBlockchainBlockHash) {
		return false
	}

	for _, tx := range b.Transactions {
		if !transaction.VerifyTransaction(tx) {
			return false
		}
	}

	return true

}

func GetUserTxHistory(bc *Blockchain, userAddress []byte) (s string) {
	for in, _ := range *bc {
		for _, el := range (*bc)[in].Transactions {
			if bytes.Equal(el.SenderAddress, userAddress) {
				s += fmt.Sprintf("Tx in block %v\n", in) + el.ToString() + "\n"
			}
		}
	}
	return s
}
