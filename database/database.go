package database

import (
	"github.com/prologic/bitcask"
	"log"
	"os"
)

// "" -> null ?
func WriteToDB(data []byte, key string, nodeNumber string) {
	dirPath, _ := os.Getwd()

	db, _ := bitcask.Open(dirPath + "/database/blockchainDB_" + nodeNumber)
	defer db.Close()

	err := db.Put([]byte(key), data)
	if err != nil {
		log.Println(err)
	}
}

func GetFromDB(key string, nodeNumber string) []byte {
	dirPath, _ := os.Getwd()

	db, _ := bitcask.Open(dirPath + "/database/blockchainDB_" + nodeNumber)
	defer db.Close()

	val, err := db.Get([]byte(key))
	if err != nil {
		log.Println(err)
	}

	return val
}
