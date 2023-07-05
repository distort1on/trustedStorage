package node

import (
	"fmt"
	"log"
	"strconv"
	"trustedStorage/blockchain"
	"trustedStorage/checkpoints"
	"trustedStorage/p2pCommunication"
	"trustedStorage/settings"
	"trustedStorage/stateWorker"
	"trustedStorage/test_pr"
)

func systemLaunch() {
	//launch first node in system and fill blockchain with generated data
	settings.NodeSettingsFileName = "settings/nodeSettings" + NodeNum + ".json"
	fmt.Println(settings.NodeSettingsFileName)

	log.Println("System launched, Initializing first node")

	stateWorker.SetNodeState("Working")

	test_pr.TestFillBlockchain()
	SaveNode()
}

func askForFullChainLaunch() {
	var err error
	masterNodesIdsList := settings.GetMasterNodeIds()

	log.Println("Database is empty, asking full chain from master node and compare with checkpoints")

	//sending message ask for full chain

	stateWorker.SetNodeState("Waiting_Full_Chain")

	//try to ask all master nodes

	for _, el := range masterNodesIdsList {

		if el == p2pCommunication.Node.ID().String() {
			continue
		}
		err = p2pCommunication.WaitForDelivery(el, []byte{'a'}, 10)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Waiting for data")
			break
		}
	}

	if err == nil && stateWorker.WaitForStateChanged("Waiting_Full_Chain", "Received_Full_Chain", 60) {
		log.Println("Starting to verify checkpoints")
	} else {
		log.Fatalln("Master peer afk")
	}

	if !checkpoints.VerifyCheckPoints() {
		log.Println("Hash value doesn't match")
	} else {
		log.Println("Checkpoints correct, starting normal work")
	}
	stateWorker.SetNodeState("Working")
}

func verifyAndAskForMissingBlocks() {
	var err error
	masterNodesIdsList := settings.GetMasterNodeIds()
	currentBlockchainHeight := strconv.Itoa(blockchain.BlockChainIns.GetCurrentHeight())

	log.Println("Asking if I have a full blockchain from master nodes")

	stateWorker.SetNodeState("Waiting_Missing_Blocks")
	for _, el := range masterNodesIdsList {

		if el == p2pCommunication.Node.ID().String() {
			continue
		}
		err = p2pCommunication.WaitForDelivery(el, []byte("h"+currentBlockchainHeight), 10)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}

	if err == nil && stateWorker.WaitForStateChanged("Waiting_Missing_Blocks", "Received_Missing_Blocks", 60) {
		log.Println("Blockchain is valid")
	} else {
		if settings.IsMasterNode(p2pCommunication.Node.ID().String()) {
			log.Println("Its the only one online master node, starting normal work")
			stateWorker.SetNodeState("Working")
			return
		} else {
			log.Fatalln("Master peer afk")
		}
	}
	stateWorker.SetNodeState("Working")

}
