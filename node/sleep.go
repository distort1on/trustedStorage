package node

import (
	"log"
	"trustedStorage/blockchain"
	"trustedStorage/database"
	"trustedStorage/mempool"
	"trustedStorage/serialization"
)

func SleepNode() {
	log.Println("Writing data to the database")
	bcBytes := serialization.Serialize(blockchain.BlockChainIns)
	database.WriteToDB(bcBytes, "blockchain", NodeNum) //maybe write each block?

	mpBytes := serialization.Serialize(mempool.MemPoolIns)
	database.WriteToDB(mpBytes, "mempool", NodeNum)

	//currState := stateWorker.GetCurrentNodeState()
	//stateBytes := serialization.Serialize(currState)
	//database.WriteToDB(stateBytes, "current_state", NodeNum)
	//
	////todo check if in consensus phase
	//consesusBytes := serialization.Serialize(p2pCommunication.CurConsensusMessage)
	//database.WriteToDB(consesusBytes, "consensus_message", NodeNum)
}
