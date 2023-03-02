package test_pr

import (
	"fmt"
	"trustedStorage/database"
	"trustedStorage/serialization"
)

func Test2() {
	bcBytes := database.GetFromDB("blockchain")
	err := serialization.DeSerialize(&blockChain, bcBytes)
	if err != nil {
		panic(err)
	}

	mpBytes := database.GetFromDB("mempool")
	err = serialization.DeSerialize(&memPool, mpBytes)
	if err != nil {
		panic(err)
	}

	txDbBytes := database.GetFromDB("txDatabase")
	err = serialization.DeSerialize(&txDataBase, txDbBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println("Blockchain: \n" + blockChain.ToString())

	fmt.Println("Mempool: \n" + memPool.ToString())

	fmt.Println("Transaction Database: \n")
	for key, value := range txDataBase {

		fmt.Printf("\nDocument: %x\n%v\n", key, value.ToString())
	}

}
