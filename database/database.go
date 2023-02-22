package database

import (
	"github.com/prologic/bitcask"
	"os"
)

// "" -> null
func WriteToDB(data []byte, key string) {
	dirPath, _ := os.Getwd()
	db, _ := bitcask.Open(dirPath + "/database/db")
	defer db.Close()
	db.Put([]byte(key), data)
}

func GetFromDB(key string) []byte {
	dirPath, _ := os.Getwd()

	db, _ := bitcask.Open(dirPath + "/database/db")
	defer db.Close()

	val, _ := db.Get([]byte(key))

	return val
}
