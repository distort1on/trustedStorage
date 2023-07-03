package p2pCommunication

import (
	"context"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
	"log"
	"strconv"
	"sync"
	"time"
	"trustedStorage/blockchain"
	"trustedStorage/mempool"
	"trustedStorage/serialization"
	"trustedStorage/settings"
	"trustedStorage/stateWorker"
	"trustedStorage/transaction"
)

var StartTime int64
var StartHeight int

func ProposeTransaction(tx transaction.Transaction) {
	log.Printf("Proposing transaction %x to other nodes\n", tx.GetTxHash())
	txBytes := serialization.Serialize(tx)
	data := append([]byte{'t'}, txBytes...)
	masterPeers := settings.GetMasterNodeIds()
	SendDataToPeersInList(masterPeers, data, 0)
}

func AddReceivedBlockToBlockchain(data []byte, bc *blockchain.Blockchain) {
	var block blockchain.Block
	err := serialization.DeSerialize(&block, data)
	if err != nil {
		log.Println(err)
	}
	err = bc.AcceptingBlock(&block)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Block on height %v created", len(*bc)-1)

	StartTime = time.Now().Unix()
	StartHeight = blockchain.BlockChainIns.GetCurrentHeight()
}

func SendDataToConnectedPeerByFullAddress(node host.Host, targetPeerAddr string, data []byte) error {
	log.Println("sending message to ", targetPeerAddr)

	addr, err := multiaddr.NewMultiaddr(targetPeerAddr)
	if err != nil {
		return err
	}
	pAddr, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}

	stream, err := node.NewStream(context.Background(), pAddr.ID, "/send/1.0.0")
	if err != nil {
		return err
	}

	_, err = stream.Write(data)
	if err != nil {
		return err
	}

	err = stream.Close()
	if err != nil {
		return err
	}
	return nil
}

func SendDataToConnectedPeerByPeerID(node host.Host, targetPeerID string, data []byte) error {
	log.Println("sending message to ", targetPeerID)

	mult, err := multihash.FromB58String(targetPeerID)
	if err != nil {
		return err
	}
	addrInfo := node.Peerstore().PeerInfo(peer.ID(mult))

	p2pAddr, err := peer.AddrInfoToP2pAddrs(&addrInfo)
	if err != nil {
		return err
	}

	addr, err := multiaddr.NewMultiaddr(p2pAddr[0].String())
	if err != nil {
		return err
	}
	pAddr, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}

	stream, err := node.NewStream(context.Background(), pAddr.ID, "/send/1.0.0")
	if err != nil {
		//if timeout != 0 {
		//	log.Println(err)
		//	result := WaitForDelivery(targetPeerID, data, timeout)
		//	if result {
		//		return nil
		//	} else {
		//		return errors.New("request wasn't delivered")
		//	}
		//}

		return err
	}

	_, err = stream.Write(data)
	if err != nil {
		return err
	}

	err = stream.Close()
	if err != nil {
		return err
	}

	return nil
}

func SendDataToConnectedPeersWithWait(node host.Host, data []byte, timeout int64) {
	connectedPeers := node.Peerstore().Peers()
	//fmt.Println(connectedPeers)
	var wg sync.WaitGroup
	var err error
	wg.Add(len(connectedPeers))

	for _, e := range connectedPeers {

		p2pAddrStruct := node.Peerstore().PeerInfo(e)

		go func() {
			if p2pAddrStruct.ID == Node.ID() {
				wg.Done()
			} else {
				//err = SendDataToConnectedPeerByPeerID(node, p2pAddrStruct.ID.String(), data)
				err = WaitForDelivery(p2pAddrStruct.ID.String(), data, timeout)
				if err != nil {
					log.Println(err)
				} else {
					log.Println("Request delivered successfully")
				}
				wg.Done()
			}

		}()
	}
	wg.Wait()

}

//func SendDataToConnectedPeers(node host.Host, data []byte) {
//
//	connectedPeers := node.Peerstore().Peers()
//	//fmt.Println(connectedPeers)
//
//	for _, e := range connectedPeers {
//
//		p2pAddrStruct := node.Peerstore().PeerInfo(e)
//		p2pAddr, err := peer.AddrInfoToP2pAddrs(&p2pAddrStruct)
//		if err != nil {
//			log.Println(err)
//		}
//		if p2pAddr[0].String() == settings.GetP2PFullAddress() {
//			continue
//		}
//		err = SendDataToConnectedPeer(node, p2pAddr[0].String(), data)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}

func SendDataToPeersInMap(peersMap map[peer.ID]bool, data []byte, timeout int64) {

	var wg sync.WaitGroup
	var err error

	wg.Add(len(peersMap))

	for key := range peersMap {

		key := key
		go func() {
			if key == Node.ID() {
				wg.Done()
			} else {
				err = WaitForDelivery(key.String(), data, timeout)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			}
		}()
	}
	wg.Wait()

}

func SendDataToPeersInList(peersList []string, data []byte, timeout int64) {
	var wg sync.WaitGroup
	var err error

	wg.Add(len(peersList))

	for _, p := range peersList {
		p := p
		go func() {
			if p == Node.ID().String() {
				wg.Done()
			} else {
				err = WaitForDelivery(p, data, timeout)
				if err != nil {
					log.Println(err)
				}
				wg.Done()
			}
		}()
	}
	wg.Wait()
}

func NodeActionDecision(data []byte, connPeerID peer.ID, fromQueue bool) {
	//todo check if received message is in consensus group on each state
	if data[0] == '0' && stateWorker.GetCurrentNodeState() == "Working" {
		//send consensus propose
		PrePrepare(data[1:], connPeerID)

	} else if string(data[:3]) == "ppr" && stateWorker.GetCurrentNodeState() == "Starting_Consensus" {
		//proposer receive answer from backup nodes
		//ResponseReceived(connPeerID.String())
		mu.Lock()
		CurConsensusMessage.consensusGroup[connPeerID] = true
		mu.Unlock()

	} else if string(data[:3]) == "ppf" && stateWorker.GetCurrentNodeState() == "Consensus_Pre_Prepare" {
		//ResponseReceived(connPeerID.String())
		//backup node receive from proposer, starting consensus, receiving consensus group
		if connPeerID == CurConsensusMessage.currentProposer {
			CurConsensusMessage.waitingList[connPeerID] = true
			err := serialization.DeSerialize(&CurConsensusMessage.consensusGroup, data[3:])
			if err != nil {
				log.Println(err)
			}
			log.Println("received group", CurConsensusMessage.consensusGroup)
		}

	} else if data[0] == '1' && stateWorker.GetCurrentNodeState() == "Consensus_Prepare" {
		_, ok := CurConsensusMessage.consensusGroup[connPeerID]
		if ok {
			CheckPrepared(data[1:], connPeerID)
		}

	} else if string(data[:2]) == "pf" && stateWorker.GetCurrentNodeState() == "Starting_Consensus" {
		_, ok := CurConsensusMessage.consensusGroup[connPeerID]
		if ok {
			mu.Lock()
			CurConsensusMessage.waitingList[connPeerID] = true
			mu.Unlock()
		}
	} else if data[0] == '2' && stateWorker.GetCurrentNodeState() == "Consensus_Prepare" {
		if CurConsensusMessage.currentProposer == connPeerID {
			mu.Lock()
			CurConsensusMessage.waitingList[connPeerID] = true
			mu.Unlock()
		}
	} else if data[0] == 'c' && stateWorker.GetCurrentNodeState() == "Consensus_Commit" {
		_, ok := CurConsensusMessage.consensusGroup[connPeerID]
		if ok {
			CheckCommit(data[1:], connPeerID)
		}
	} else if string(data[:2]) == "fc" && stateWorker.GetCurrentNodeState() == "Consensus_Commit" && Node.ID() == CurConsensusMessage.currentProposer {
		_, ok := CurConsensusMessage.consensusGroup[connPeerID]
		if ok {
			mu.Lock()
			CurConsensusMessage.TotalCount++
			mu.Unlock()
		}

	} else if data[0] == 'f' && stateWorker.GetCurrentNodeState() == "Finishing" {
		if connPeerID == CurConsensusMessage.currentProposer {
			CurConsensusMessage.waitingList[connPeerID] = true
		}

	} else if data[0] == 'a' && stateWorker.GetCurrentNodeState() == "Working" {
		bcBytes := serialization.Serialize(blockchain.BlockChainIns)
		err := SendDataToConnectedPeerByPeerID(Node, connPeerID.String(), append([]byte{'r'}, bcBytes...))
		if err != nil {
			log.Println(err)
		}
	} else if data[0] == 'r' && stateWorker.GetCurrentNodeState() == "Waiting_Full_Chain" {
		log.Println("Received full chain")
		err := serialization.DeSerialize(&blockchain.BlockChainIns, data[1:])
		if err != nil {
			log.Println(err)
		}
		stateWorker.SetNodeState("Received_Full_Chain")
	} else if data[0] == 'h' && stateWorker.GetCurrentNodeState() == "Working" {
		askBlockchainHeight, err := strconv.Atoi(string(data[1:]))
		if err != nil {
			log.Println(err)
		}

		if askBlockchainHeight != blockchain.BlockChainIns.GetCurrentHeight() {
			bcBytes := serialization.Serialize((*blockchain.BlockChainIns)[askBlockchainHeight:])
			err = SendDataToConnectedPeerByPeerID(Node, connPeerID.String(), append([]byte{'m'}, bcBytes...))
			if err != nil {
				log.Println(err)
			}
		} else {
			//height is ok
			err = SendDataToConnectedPeerByPeerID(Node, connPeerID.String(), append([]byte{'m'}))
			if err != nil {
				log.Println(err)
			}
		}

	} else if data[0] == 'm' && stateWorker.GetCurrentNodeState() == "Waiting_Missing_Blocks" {
		if len(data) == 1 {
			log.Println("My blockchain is valid, start normal working")
		} else {
			log.Println("Received missing blocks")
			var missBlocks blockchain.Blockchain

			err := serialization.DeSerialize(&missBlocks, data[1:])
			if err != nil {
				log.Println(err)
			}

			for _, block := range missBlocks {
				err = blockchain.BlockChainIns.AcceptingBlock(block)
				if err != nil {
					log.Println(err)
				}
			}

		}
		stateWorker.SetNodeState("Received_Missing_Blocks")
	} else if data[0] == 't' {
		log.Println("Received proposed transaction")
		var tx transaction.Transaction

		err := serialization.DeSerialize(&tx, data[1:])
		if err != nil {
			log.Println(err)
		}

		err = mempool.MemPoolIns.AddTxToMempool(tx)
		if err != nil {
			log.Println(err)
		}
	} else if !fromQueue {
		RequestQueueIns.AddRequest(data, connPeerID)
	}
}
