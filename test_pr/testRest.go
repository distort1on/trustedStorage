package test_pr

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strconv"
//	"trustedStorage/blockchain"
//)
//
//import (
//	"errors"
//	"trustedStorage/database"
//	"trustedStorage/serialization"
//	"trustedStorage/transaction"
//)
//
////var memPool *mempool.Mempool = mempool.InitMempool()
////var txDB = transaction.TransactionDataBase{
////	TxDataBase: make(map[string]transaction.Transaction),
////}
//
//func getBlockchain(context *gin.Context) {
//	//&blockChain?
//	context.IndentedJSON(http.StatusOK, blockChain)
//}
//
//func getBlockByHeight(height string) (*blockchain.Block, error) {
//	h, err := strconv.ParseUint(height, 10, 32)
//	if err != nil {
//		panic(err)
//	}
//
//	if uint64(len(*blockChain)) > h {
//		return (*blockChain)[h], nil
//	}
//
//	return nil, errors.New("Block not found")
//}
//
//func getBlock(context *gin.Context) {
//	height := context.Param("height")
//
//	block, err := getBlockByHeight(height)
//	if err != nil {
//		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Block not found"})
//		panic(err)
//	}
//
//	context.IndentedJSON(http.StatusOK, block)
//}
//
//func getLastBlock(context *gin.Context) {
//	block, err := getBlockByHeight(strconv.Itoa(len(*blockChain) - 1))
//	if err != nil {
//		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Block not found"})
//		panic(err)
//	}
//	context.IndentedJSON(http.StatusOK, block)
//
//}
//
//func addTX(context *gin.Context) {
//	var tx transaction.Transaction
//
//	if err := context.BindJSON(&tx); err != nil {
//		panic(err)
//	}
//	fmt.Println(tx.Data)
//
//	err := memPool.AddTxToMempool(tx, )
//	if err != nil {
//		panic(err)
//	}
//
//	//todo add goroutine
//	if len(memPool) > numOfTransactionsInBlock {
//		block := blockchain.CreateBlock(1, (*blockChain)[len(*blockChain)-1].GetBlockHash(), memPool.FormTransactionsList(numOfTransactionsInBlock))
//		err := blockChain.AcceptingBlock(&block)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
//
//	//
//
//	context.IndentedJSON(http.StatusCreated, tx)
//}
//
//func getHeight(context *gin.Context) {
//	height := len(*blockChain) - 1
//
//	context.IndentedJSON(http.StatusOK, height)
//}
//
//func RunRest() {
//	//db already contains blocks
//	bcBytes := database.GetFromDB("blockchain")
//	err := serialization.DeSerialize(&blockChain, bcBytes)
//	if err != nil {
//		panic(err)
//	}
//
//	mpBytes := database.GetFromDB("mempool")
//	err = serialization.DeSerialize(&memPool, mpBytes)
//	if err != nil {
//		panic(err)
//	}
//
//	//qrcode.WriteQrToFile(1)
//
//	router := gin.Default()
//	router.GET("/blockchain", getBlockchain)
//	router.GET("/blockchain/:height", getBlock)
//	router.GET("/blockchain/height", getHeight)
//	router.GET("/blockchain/lastBlock", getLastBlock)
//	router.POST("/blockchain", addTX)
//	err = router.Run("localhost:8080")
//	if err != nil {
//		return
//	}
//
//}
