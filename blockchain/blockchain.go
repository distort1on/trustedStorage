package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"trustedStorage/keypair"
	"trustedStorage/signature"
	"trustedStorage/size"
	"trustedStorage/transaction"
)

var BlockChainIns *Blockchain

type Blockchain []*Block

type BlockHeader struct {
	Version       uint
	HashPrevBlock []byte
	MerkleRoot    []byte
	Time          int64
}

type Block struct {
	//MacigNo
	BlockSize uint
	BlockHeader
	TxCounter    uint
	Transactions []transaction.Transaction
}

func CreateBlock(ver uint, prevBlockHash []byte, transactions []transaction.Transaction) (b Block) {
	var emptyList [][]byte
	var txList [][]byte

	b.Version = ver
	b.HashPrevBlock = prevBlockHash
	b.Time = time.Now().Unix()
	b.Transactions = transactions
	b.TxCounter = uint(len(transactions))
	b.BlockSize = uint(size.Of(b))

	for _, tx := range b.Transactions {
		txList = append(txList, tx.GetTxHash())
	}
	b.MerkleRoot = GenMerkleRoot(emptyList, txList)
	return b
}

func (b *Block) GetBlockHash() []byte {
	//block header hash
	blockSumBytes := [][]byte{
		[]byte(strconv.FormatUint(uint64(b.Version), 10)),
		b.HashPrevBlock,
		b.MerkleRoot,
		[]byte(strconv.FormatInt(b.Time, 10)),
	}

	//for _, tx := range b.Transactions {
	//	blockSumBytes = append(blockSumBytes, tx.GetTxHash())
	//}

	blockHash := sha256.Sum256(bytes.Join(blockSumBytes, []byte{}))

	return blockHash[:]
}

func CheckTxAlreadyExist(tx transaction.Transaction) bool {
	txHash := tx.GetTxHash()
	for _, elB := range *BlockChainIns {
		for _, elTx := range elB.Transactions {
			if bytes.Equal(txHash, elTx.GetTxHash()) {
				return true
			}
		}
	}
	return false
}

func FindTransactionByHash(txHash []byte) (bool, int, transaction.Transaction) {
	for bI, block := range *BlockChainIns {
		for _, t := range block.Transactions {
			if bytes.Equal(t.GetTxHash(), txHash) {
				return true, bI, t
			}
		}
	}
	return false, 0, transaction.Transaction{}
}

func FindDocumentByHash(documentHash []byte) (bool, int, transaction.Transaction) {
	for bI, block := range *BlockChainIns {
		for _, t := range block.Transactions {
			if bytes.Equal(t.Data, documentHash) {
				return true, bI, t
			}
		}
	}
	return false, 0, transaction.Transaction{}
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
	//out, err := json.MarshalIndent(b, "", "\t")
	//if err != nil {
	//	panic(err)
	//}
	result := fmt.Sprintf(
		"\n\t'BLockSize' : %d\n\t'Version' : %d\n\t'HashPrevBlock' : %x\n\t'MerkleRoot' : %x\n\t'Time' : %d\n\t'TxCounter' : %d\n\t",
		b.BlockSize,
		b.Version,
		b.HashPrevBlock,
		b.MerkleRoot,
		b.Time,
		b.TxCounter,
	)
	result += fmt.Sprintf("Transactions: [\n")
	for _, el := range b.Transactions {
		result += el.ToString()
	}
	result += "]"

	//return fmt.Sprint(string(out))
	return result
}

func (bc *Blockchain) ToString() (s string) {

	for i, block := range *bc {

		s += fmt.Sprintf("\nBlock on height %v\n", i) + block.ToString() + "\n"
		//s = strings.Join([]string{s}, block.ToString())
	}
	return s
}

func (bc *Blockchain) AcceptingBlock(b *Block) error {
	//todo first verify then accept
	err := VerificationBlock(b, BlockChainIns.GetLastBlock().GetBlockHash())
	if err == nil {
		*bc = append(*bc, b)
		log.Printf("Block %x has been added to blockchain\n", b.GetBlockHash())
		return nil
	} else {
		return err
	}
}

func VerifyBlockChain(bc *Blockchain, prevBlockHash []byte) bool {

	for in, block := range *bc {
		if in == 0 {
			if !bytes.Equal(prevBlockHash, block.HashPrevBlock) {
				return false
			}
			continue
		}

		if VerificationBlock(block, (*bc)[in-1].GetBlockHash()) != nil {
			log.Println(VerificationBlock(block, (*bc)[in-1].GetBlockHash()))
			return false
		}
	}
	return true
}

func VerificationBlock(b *Block, prevBlockHash []byte) error {
	var err error

	//check prev block hash
	if !bytes.Equal(prevBlockHash, b.HashPrevBlock) {
		err = errors.New("hash prev block is wrong")
		return err
	}

	//check merkle root
	var emptyList [][]byte
	var txList [][]byte
	for _, tx := range b.Transactions {
		txList = append(txList, tx.GetTxHash())
	}
	merkleRootToVer := GenMerkleRoot(emptyList, txList)
	if !bytes.Equal(merkleRootToVer, b.MerkleRoot) {
		err = errors.New("merkle root is wrong")
		return err
	}
	//check time
	curTime := time.Now().Unix()
	if curTime < b.Time {
		err = errors.New("time is wrong")
		return err
	}

	//if BlockChainIns.GetCurrentHeight() != blockHeight {
	//	err = errors.New("cant create block on given height")
	//	return err
	//}
	//check if block already exist
	bHash := b.GetBlockHash()
	for _, el := range *BlockChainIns {
		if bytes.Equal(el.GetBlockHash(), bHash) {
			err = errors.New("block is already exist")
			return err
		}
	}
	//check all tx
	for _, elTx := range b.Transactions {
		err = VerifyTransaction(elTx)
		if err != nil {
			return errors.New("tx is wrong")
		}
	}

	return nil
}

func VerifyTransaction(tx transaction.Transaction) error {

	var err error
	//
	//for i := 0; i < txFields.NumField(); i++ {
	//	fmt.Printf("Field: %s\tValue: %v\n", typeOfTxFields.Field(i).Name, len(txFields.Field(i).Interface()))
	//	if txFields.Field(i).Interface() == nil {
	//		err = errors.New(fmt.Sprintf("Transaction from %x incorrect\nField : %s is empty", tx.SenderAddress, typeOfTxFields.Field(i).Name))
	//		return false, err
	//	}
	//
	//}
	if len(tx.SenderAddress) != 32 {
		err = errors.New(fmt.Sprintf("Tx is incorrect\nProblem with SenderAdress size"))
		return err
	}
	if len(tx.PubKey) != 178 {
		err = errors.New(fmt.Sprintf("Tx is incorrect\nProblem with PubKey size"))
		return err
	}
	if len(tx.Data) != 32 {
		err = errors.New(fmt.Sprintf("Tx is incorrect\nProblem with Data size"))
		return err
	}

	if len(tx.Cid) != 46 {
		err = errors.New(fmt.Sprintf("Tx is incorrect\nProblem with cid size"))
		return err
	}

	txDataToVerifySignature := bytes.Join([][]byte{tx.SenderAddress, tx.Data}, []byte{})

	decodedPubKey, err := keypair.DecodePublicKey(tx.PubKey)
	if err != nil {
		return err
	}

	if CheckTxAlreadyExist(tx) {
		err = errors.New("transaction already exist in blockchain")
		return err
	}

	if signature.VerifySignature(tx.Signature, txDataToVerifySignature, decodedPubKey) {
		//log.Printf("Transaction from %x correct", tx.SenderAddress)
		return nil
	} else {
		err = errors.New(fmt.Sprintf("Signature is incorrect"))
		//log.Printf("Transaction from %x incorrect", tx.SenderAddress)
		return err
	}
}

func GetUserTxHistory(bc *Blockchain, userAddress []byte) (s string) {
	for in, _ := range *bc {
		for _, el := range (*bc)[in].Transactions {
			if bytes.Equal(el.SenderAddress, userAddress) {
				s += fmt.Sprintf("Tx in block %v\nTx hash:\n%v\n%v\n", in, hex.EncodeToString(el.GetTxHash()), el.ToString())
			}
		}
	}
	return s
}

func GenMerkleRoot(resList [][]byte, txList [][]byte) []byte {
	if len(txList) == 0 {
		res := sha256.New()
		res.Write([]byte{})
		resList = append(resList, res.Sum(nil))
		return resList[0]
	}
	if len(txList)%2 == 0 {
		for i := 0; i < len(txList); i += 2 {
			res := sha256.New()
			res.Write(bytes.Join([][]byte{txList[i], txList[i+1]}, []byte{}))
			resList = append(resList, res.Sum(nil))
			//fmt.Println(hex.EncodeToString(res.Sum(nil)))
		}
	} else {
		for i := 0; i < len(txList)-1; i += 2 {
			res := sha256.New()
			res.Write(bytes.Join([][]byte{txList[i], txList[i+1]}, []byte{}))
			//res.Write([]byte(hex.EncodeToString(txList[i]) + hex.EncodeToString(txList[i+1])))
			resList = append(resList, res.Sum(nil))
		}
		resList = append(resList, txList[len(txList)-1])
	}

	txList = resList

	if len(txList) == 1 {
		return txList[0]
	} else {
		resList = nil
		return GenMerkleRoot(resList, txList)
	}

}

func (bc *Blockchain) GetLastBlock() *Block {
	return (*bc)[len(*bc)-1]
}

func (bc *Blockchain) GetCurrentHeight() int {
	return len(*bc)
}
