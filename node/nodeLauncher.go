package node

import (
	"fmt"
	"log"
	"time"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
	"trustedStorage/p2pCommunication"
	"trustedStorage/serialization"
	"trustedStorage/settings"
	"trustedStorage/stateWorker"
	"trustedStorage/test_pr"
	"trustedStorage/transaction"
)

func BlocksCreatingProcess(memPool *mempool.MempoolTransactions) {
	//p2pCommunication.StartTime = time.Now().Unix()
	//p2pCommunication.StartHeight = blockchain.BlockChainIns.GetCurrentHeight()
	numOfTransactionsInBlock := int(settings.GetNumOfTransactionsInBlock())
	for {
		if stateWorker.GetCurrentNodeState() != "Working" {
			//time.Sleep(time.Second * 5)
			continue
		}

		if settings.GetCurrentProposer() == p2pCommunication.Node.ID().String() {
			time.Sleep(time.Duration(settings.GetConsensusTime()) * time.Second)
			mempool.MemPoolIns.RemoveExistingTxFromMempool()

			if len(*memPool) >= numOfTransactionsInBlock {
				block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), memPool.FormTransactionsList(numOfTransactionsInBlock))

				blockBytes := serialization.Serialize(block)

				if !p2pCommunication.StartConsensus(blockBytes) {
					mempool.MemPoolIns.ReturnTxToMempool(block.Transactions)
				}

			} else {
				var tt []transaction.Transaction
				block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), tt)

				blockBytes := serialization.Serialize(block)

				if !p2pCommunication.StartConsensus(blockBytes) {
					mempool.MemPoolIns.ReturnTxToMempool(block.Transactions)
				}
			}
		} else if settings.GetNextProposer() == p2pCommunication.Node.ID().String() {
			if time.Now().Unix()-p2pCommunication.StartTime > settings.GetConsensusTime()+10 && p2pCommunication.StartHeight == blockchain.BlockChainIns.GetCurrentHeight() {
				mempool.MemPoolIns.RemoveExistingTxFromMempool()

				if len(*memPool) >= numOfTransactionsInBlock {
					log.Println("Consensus should have already begun, replacing proposer")
					block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), memPool.FormTransactionsList(numOfTransactionsInBlock))
					blockBytes := serialization.Serialize(block)

					if !p2pCommunication.StartConsensus(blockBytes) {
						mempool.MemPoolIns.ReturnTxToMempool(block.Transactions)
					}
				} else {
					var tt []transaction.Transaction
					block := blockchain.CreateBlock(1, (*blockchain.BlockChainIns)[len(*blockchain.BlockChainIns)-1].GetBlockHash(), tt)

					blockBytes := serialization.Serialize(block)

					if !p2pCommunication.StartConsensus(blockBytes) {
						mempool.MemPoolIns.ReturnTxToMempool(block.Transactions)
					}
				}
				time.Sleep(time.Second * 10)
			}
		}
	}
}

func init() {
	blockchain.BlockChainIns = blockchain.InitBlockchain()
	mempool.MemPoolIns = mempool.InitMempool()
	p2pCommunication.CurConsensusMessage = &p2pCommunication.ConsensusMessage{}
	p2pCommunication.RequestQueueIns = p2pCommunication.InitQueue()
	stateWorker.SetNodeState("None")
}

var NodeNum string

func LaunchNode() {
	fmt.Scan(&NodeNum)

	settings.NodeSettingsFileName = "settings/nodeSettings" + NodeNum + ".json"
	fmt.Println(settings.NodeSettingsFileName)
	log.Println("Initializing node")

	stateWorker.SetNodeState("Launching")
	go test_pr.StartGrpc()
	go p2pCommunication.LaunchP2PPeer()
	go p2pCommunication.RequestQueueIns.StartNodeQueueProcess()
	time.Sleep(time.Second * 10)

	if !NodeAwakening() {
		askForFullChainLaunch()
	} else {
		//*blockchain.BlockChainIns = (*blockchain.BlockChainIns)[:1]
		verifyAndAskForMissingBlocks()
	}
	p2pCommunication.StartHeight = blockchain.BlockChainIns.GetCurrentHeight()
	p2pCommunication.StartTime = time.Now().Unix()

	if settings.IsMasterNode(p2pCommunication.Node.ID().String()) {
		log.Println("Starting blocks creating process")
		go BlocksCreatingProcess(mempool.MemPoolIns)
	}

	var cmd string
	for {
		_, err := fmt.Scan(&cmd)
		if err != nil {
			log.Println(err)
		}
		chooseCommand(cmd)

	}

}
