package node

import (
	"log"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/mempool"
	"trustedStorage/serialization"
)

func NodeAwakening() bool {
	log.Println("Reading data from the database")

	var bcToTest blockchain.Blockchain
	bcBytes := database.GetFromDB("blockchain", NodeNum)
	err := serialization.DeSerialize(&bcToTest, bcBytes)
	if err != nil {
		log.Println(err)
		log.Println("Blockchain does not exist")
		return false
	}
	if !blockchain.VerifyBlockChain(&bcToTest, make([]byte, byte(0))) {
		log.Fatalln("Blockchain from db is incorrect")
	} else {
		log.Println("Blockchain from db verified and correct")
		blockchain.BlockChainIns = &bcToTest
	}

	mpBytes := database.GetFromDB("mempool", NodeNum)
	err = serialization.DeSerialize(&mempool.MemPoolIns, mpBytes)
	if err != nil {
		log.Println(err)
		return false
	}

	//var currState string
	//stBytes := database.GetFromDB("current_state", NodeNum)
	//err = serialization.DeSerialize(&currState, stBytes)
	//if err != nil {
	//	log.Println(err)
	//	return false
	//}
	//stateWorker.SetNodeState(currState)
	//
	//consensusBytes := database.GetFromDB("consensus_message", NodeNum)
	//err = serialization.DeSerialize(&p2pCommunication.CurConsensusMessage, consensusBytes)
	//if err != nil {
	//	log.Println(err)
	//	return false
	//}
	return true
}
