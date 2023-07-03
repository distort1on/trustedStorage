package p2pCommunication

import (
	"bytes"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"math"
	"strconv"
	"sync"
	"time"
	"trustedStorage/blockchain"
	"trustedStorage/serialization"
	"trustedStorage/settings"
	"trustedStorage/stateWorker"
)

// PBFT CONSENSUS
// todo dont send messeges to peers that already not in consensus
// todo quit time.sleep if alredy have all requests

var totalNodesNum int
var fault int

const DEFAULT_TIME_OUT = 10

var mu sync.Mutex

type ConsensusMessage struct {
	blockData          []byte
	TotalCount         int
	consensusGroup     map[peer.ID]bool
	waitingList        map[peer.ID]bool
	currentProposer    peer.ID
	consensusStartTime time.Time
}

func InitConsensusMessage() *ConsensusMessage {
	var msg ConsensusMessage
	msg.blockData = nil
	msg.TotalCount = 1
	msg.consensusGroup = make(map[peer.ID]bool)
	msg.waitingList = make(map[peer.ID]bool)
	msg.consensusGroup[Node.ID()] = true

	return &msg
}

func verifyReceivedConsensusData(blockData []byte, supposedProposer peer.ID) bool {
	if supposedProposer.String() != settings.GetCurrentProposer() {
		log.Println("Not expected proposer, verify")
		// firstly check if real proposer rly offline
		masterPeers := settings.GetMasterNodeIds()
		flag := false

		for _, el := range masterPeers {
			if el == supposedProposer.String() {
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}

		//check if real proposer did not proposed block
		if time.Now().Unix()-StartTime > settings.GetConsensusTime()+10 {
			log.Println("Expected proposer afk, proposer changed")
		}

	}

	var block blockchain.Block
	err := serialization.DeSerialize(&block, blockData)
	if err != nil {
		log.Println(err)
		return false
	}
	if err = blockchain.VerificationBlock(&block, blockchain.BlockChainIns.GetLastBlock().GetBlockHash()); err != nil {
		log.Println(err)
		return false
	}

	return true
}

var CurConsensusMessage *ConsensusMessage

func StartConsensus(data []byte) bool {
	stateWorker.SetNodeState("Starting_Consensus")
	log.Println("Initiating the consensus process")
	CurConsensusMessage = InitConsensusMessage()
	CurConsensusMessage.blockData = data
	CurConsensusMessage.currentProposer = Node.ID()
	CurConsensusMessage.consensusGroup[Node.ID()] = true
	CurConsensusMessage.consensusStartTime = time.Now()

	timeStampString := strconv.FormatInt(CurConsensusMessage.consensusStartTime.Unix(), 10)
	timeStampStringLength := strconv.Itoa(len(timeStampString))

	nData := append([]byte(timeStampStringLength), timeStampString...)
	nData = append(nData, data...)
	nData = append([]byte{'0'}, nData...)

	totalNodesNum = len(Node.Network().Peers()) + 1
	fault = int(math.Trunc(float64((totalNodesNum - 1) / 3)))
	if fault == 0 {
		log.Println("Consensus process cant be started, must be more connected peers")
		stateWorker.SetNodeState("Working")
		return false
	}

	//send 0 wait for ppr
	SendDataToConnectedPeersWithWait(Node, nData, int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT)).Seconds()))

	// wait for responses
	time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 2)))

	if len(CurConsensusMessage.consensusGroup) < 3*fault+1 {
		log.Println("Consensus process cant be started, must be more ready peers")
		stateWorker.SetNodeState("Working")
		return false
	} else {
		log.Println("Consensus started")
		consensusGroupBytes := serialization.Serialize(CurConsensusMessage.consensusGroup)

		//send ppf wait for pf

		SendDataToPeersInMap(CurConsensusMessage.consensusGroup, append([]byte("ppf"), consensusGroupBytes...), int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT*3)).Seconds()))

		time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 7)))

		if len(CurConsensusMessage.waitingList) < 2*fault {
			log.Println("Haven't received enough prepared messages, quiting consensus")
			stateWorker.SetNodeState("Working")
			return false
		} else {
			//starting commit
			stateWorker.SetNodeState("Consensus_Commit")
			log.Println("Starting commit phase")
			SendDataToPeersInMap(CurConsensusMessage.consensusGroup, append([]byte{'2'}, CurConsensusMessage.blockData...), int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT*8)).Seconds()))

			//wait for finish commits
			time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 11)))

			if CurConsensusMessage.TotalCount >= 2*fault {
				log.Println("Committing phase successful, adding block to blockchain")
				SendDataToPeersInMap(CurConsensusMessage.consensusGroup, []byte{'f'}, int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT*12)).Seconds()))
				AddReceivedBlockToBlockchain(CurConsensusMessage.blockData, blockchain.BlockChainIns)
				stateWorker.SetNodeState("Working")
				return true
			} else {
				log.Println("Haven't received enough commit messages, quiting consensus")
				stateWorker.SetNodeState("Working")
				return false
			}
		}

	}

}

func PrePrepare(data []byte, remotePeerID peer.ID) {
	stateWorker.SetNodeState("Consensus_Pre_Prepare")
	log.Println("Received consensus initiating message")

	CurConsensusMessage = InitConsensusMessage()
	CurConsensusMessage.currentProposer = remotePeerID
	CurConsensusMessage.waitingList = make(map[peer.ID]bool)

	timeStampBytesLength, err := strconv.Atoi(string(data[:2]))
	if err != nil {
		log.Println("Error during conversion")
		stateWorker.SetNodeState("Working")
		return
	}
	timeStamp := string(data[2 : 2+timeStampBytesLength])
	i, err := strconv.ParseInt(timeStamp, 10, 64)
	if err != nil {
		log.Println(err)
		stateWorker.SetNodeState("Working")
		return
	}
	CurConsensusMessage.consensusStartTime = time.Unix(i, 0)

	data = data[2+timeStampBytesLength:]
	verResult := verifyReceivedConsensusData(data, remotePeerID)

	if verResult {
		CurConsensusMessage.blockData = data

		//send ppr wait for ppf
		err := SendDataToConnectedPeerByPeerID(Node, remotePeerID.String(), append([]byte("ppr"), data...))
		if err != nil {
			log.Println(err)
			log.Println("Primary node afk, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}

		//wait for proposer response with group

		time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 3)))

		if CurConsensusMessage.waitingList[CurConsensusMessage.currentProposer] == false {
			log.Println("Primary node afk, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}

		Prepare()
	} else {
		log.Println("Message is invalid")
		stateWorker.SetNodeState("Working")
	}
	//todo if not accept incoming propose block
}

func Prepare() {
	stateWorker.SetNodeState("Consensus_Prepare")
	log.Println("Message is ok, send prepare to other nodes")
	CurConsensusMessage.waitingList = make(map[peer.ID]bool)
	delete(CurConsensusMessage.consensusGroup, CurConsensusMessage.currentProposer)
	data := append([]byte{'1'}, CurConsensusMessage.blockData...)

	SendDataToPeersInMap(CurConsensusMessage.consensusGroup, data, int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT*4)).Seconds()))

	//wait for other prepared messages (pf)
	time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 6)))

	if CurConsensusMessage.TotalCount >= 2*fault {
		log.Println("Prepare finished, starting commit phase")

		err := SendDataToConnectedPeerByPeerID(Node, CurConsensusMessage.currentProposer.String(), append([]byte("pf")))
		if err != nil {
			log.Println(err)
			log.Println("Primary node afk, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}

		time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 8)))
		if CurConsensusMessage.waitingList[CurConsensusMessage.currentProposer] == false {
			log.Println("Primary node afk1, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}

		Commit()
	} else {
		log.Println("Not enough prepared messages, quiting consensus process")
		stateWorker.SetNodeState("Working")
	}
}

func CheckPrepared(data []byte, remotePeerID peer.ID) {
	if bytes.Equal(data, CurConsensusMessage.blockData) {
		log.Println("prepared")
		mu.Lock()
		CurConsensusMessage.TotalCount++
		mu.Unlock()
		//err := SendDataToConnectedPeerByPeerID(Node, remotePeerID.String(), append([]byte("1r"), data...), 60)
		//todo err != response doesn't received
		//if err != nil {
		//	log.Println(err)
		//}

	} else {
		log.Println("Message is invalid")
	}
}

func Commit() {
	stateWorker.SetNodeState("Consensus_Commit")
	log.Println("Send commit to other nodes")
	CurConsensusMessage.TotalCount = 2 //node + proposer
	CurConsensusMessage.waitingList = make(map[peer.ID]bool)
	CurConsensusMessage.consensusGroup[CurConsensusMessage.currentProposer] = true

	data := append([]byte{'c'}, CurConsensusMessage.blockData...)

	SendDataToPeersInMap(CurConsensusMessage.consensusGroup, data, int64(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second*DEFAULT_TIME_OUT*9)).Seconds()))

	// wait for others commit messages
	time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 10)))

	if CurConsensusMessage.TotalCount >= 2*fault {
		log.Println("Committing finished, waiting for block commit")
		stateWorker.SetNodeState("Finishing")

		err := SendDataToConnectedPeerByPeerID(Node, CurConsensusMessage.currentProposer.String(), append([]byte("fc")))
		if err != nil {
			log.Println(err)
			log.Println("Primary node afk, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}

		//wait for block commit CurConsensusMessage
		time.Sleep(time.Until(CurConsensusMessage.consensusStartTime.Add(time.Second * DEFAULT_TIME_OUT * 12)))
		if CurConsensusMessage.waitingList[CurConsensusMessage.currentProposer] == false {
			log.Println("Primary node afk, quiting consensus process")
			stateWorker.SetNodeState("Working")
			return
		}
		log.Println("Consensus finished, add block to blockchain")
		AddReceivedBlockToBlockchain(CurConsensusMessage.blockData, blockchain.BlockChainIns)
		stateWorker.SetNodeState("Working")

	} else {
		log.Println("Not enough prepared messages, quiting consensus process")
		stateWorker.SetNodeState("Working")
	}
}

func CheckCommit(data []byte, remotePeerID peer.ID) {
	//todo add if from list first
	if bytes.Equal(data, CurConsensusMessage.blockData) {
		log.Println("commit")
		mu.Lock()
		CurConsensusMessage.TotalCount++
		mu.Unlock()
	} else {
		log.Println("Message is invalid")
	}
}
