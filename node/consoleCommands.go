package node

import (
	"log"
	"os"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
	"trustedStorage/p2pCommunication"
	"trustedStorage/serialization"
	"trustedStorage/settings"
	"trustedStorage/stateWorker"
)

func chooseCommand(cmd string) {
	if cmd == "print_blockchain" {
		log.Println(blockchain.BlockChainIns.ToString())
	} else if cmd == "print_mempool" {
		log.Println(mempool.MemPoolIns.ToString())
	} else if cmd == "print_nodeState" {
		log.Println(stateWorker.GetCurrentNodeState())
	} else if cmd == "shut_down" {
		SleepNode()
		os.Exit(0)
	} else if cmd == "system_launch" {
		systemLaunch()
	} else if cmd == "consensus" {

		block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), mempool.MemPoolIns.FormTransactionsList(int(settings.GetNumOfTransactionsInBlock())))

		blockBytes := serialization.Serialize(block)

		//if !p2pCommunication.StartConsensus(blockBytes) {
		//	mempool.ReturnTxToMempool(block.Transactions)
		//}
		if !p2pCommunication.StartConsensus(blockBytes) {
			mempool.MemPoolIns.ReturnTxToMempool(block.Transactions)
		}
	}
}
