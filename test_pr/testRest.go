package test_pr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/mempool"
	"trustedStorage/qrcode"
	"trustedStorage/serialization"
	"trustedStorage/transaction"
)

// later take from db
var bc blockchain.Blockchain
var memP *mempool.Mempool = mempool.InitMempool()
var txDB = transaction.TransactionDataBase{
	TxDataBase: make(map[string]transaction.Transaction),
}

func getBlockchain(context *gin.Context) {
	//&bc?
	context.IndentedJSON(http.StatusOK, bc)
}

func getBlockByHeight(height string) (*blockchain.Block, error) {
	h, err := strconv.ParseUint(height, 10, 32)
	if err != nil {
		panic(err)
	}

	if uint64(len(bc.BlocksHistory)) > h {
		return bc.BlocksHistory[h], nil
	}

	return nil, errors.New("Block not found")
}

func getBlock(context *gin.Context) {
	height := context.Param("height")

	block, err := getBlockByHeight(height)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Block not found"})
		panic(err)
	}

	context.IndentedJSON(http.StatusOK, block)
}

func getLastBlock(context *gin.Context) {
	block, err := getBlockByHeight(strconv.Itoa(len(bc.BlocksHistory) - 1))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Block not found"})
		panic(err)
	}
	context.IndentedJSON(http.StatusOK, block)

}

func addTX(context *gin.Context) {
	var tx transaction.Transaction
	if err := context.BindJSON(&tx); err != nil {
		panic(err)
	}

	err := memP.AddTxToMempool(tx, &txDB)
	if err != nil {
		panic(err)
	}
	context.IndentedJSON(http.StatusCreated, tx)
}

func getHeight(context *gin.Context) {
	height := len(bc.BlocksHistory) - 1

	context.IndentedJSON(http.StatusOK, height)
}

func Test2() {
	//db already contains 2 blocks
	bcBytes := database.GetFromDB("blockchain")
	bc = blockchain.Blockchain{}
	serialization.DeSerialize(&bc, bcBytes)

	qrcode.WriteQrToFile(1)

	router := gin.Default()
	router.GET("/blockchain", getBlockchain)
	router.GET("/blockchain/:height", getBlock)
	router.GET("/blockchain/height", getHeight)
	router.GET("/blockchain/lastBlock", getLastBlock)
	router.POST("/blockchain", addTX)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
