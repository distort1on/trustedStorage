package settings

import (
	"encoding/json"
	"log"
	"os"
	"trustedStorage/blockchain"
)

type NodeConfig struct {
	P2PPeerAddressFull  string
	NodeAddress         string
	RPCPeerAddress      string
	KnownPeersAddresses []string
	MasterPeersIds      []string
	NodePrivateKey      string
}

var NodeSettingsFileName = "settings/nodeSettings1.json"

func GetAllSettingsValues() *NodeConfig {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return &data
}

func GetNodePrivateKey() string {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.NodePrivateKey
}

func GetP2PNodeAddress() string {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.NodeAddress
}

func GetP2PFullAddress() string {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.P2PPeerAddressFull
}

func GetRpcNodeAddress() string {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.RPCPeerAddress
}

func GetMasterNodeIds() []string {
	content, err := os.ReadFile(NodeSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := NodeConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.MasterPeersIds
}

func IsMasterNode(NodeId string) bool {
	masterNodes := GetMasterNodeIds()

	for _, el := range masterNodes {
		if el == NodeId {
			return true
		}
	}
	return false
}

func GetCurrentProposer() string {
	masterNodes := GetMasterNodeIds()
	proposerNum := (blockchain.BlockChainIns.GetCurrentHeight() - 1) % len(masterNodes)

	return masterNodes[proposerNum]
}

func GetNextProposer() string {
	masterNodes := GetMasterNodeIds()
	nextProposerNum := (blockchain.BlockChainIns.GetCurrentHeight()-1)%len(masterNodes) + 1

	return masterNodes[nextProposerNum%len(masterNodes)]
}
