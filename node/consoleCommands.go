package node

import (
	"log"
	"os"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
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
		os.Exit(0)
	} else if cmd == "save" {
		SaveNode()
	} else if cmd == "system_launch" {
		systemLaunch()
	}
}
