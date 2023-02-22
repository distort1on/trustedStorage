package test_pr

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/serialization"
)

var bc blockchain.Blockchain

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

func Test2() {
	bcBytes := database.GetFromDB("blockchain")
	bc = blockchain.Blockchain{}
	serialization.DeSerialize(&bc, bcBytes)

	router := gin.Default()
	router.GET("/blockchain", getBlockchain)
	router.GET("/blockchain/:height", getBlock)
	router.Run("localhost:8080")
}
