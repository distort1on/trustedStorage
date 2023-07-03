package settings

import (
	"encoding/json"
	"log"
	"os"
)

type BlockConfig struct {
	NumOfTransactionsInBlock uint64
	ConsensusTime            int64
}

var blockchainSettingsFileName = "settings/blockchainSettings.json"

func GetNumOfTransactionsInBlock() uint64 {
	content, err := os.ReadFile(blockchainSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := BlockConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.NumOfTransactionsInBlock
}

func GetConsensusTime() int64 {
	content, err := os.ReadFile(blockchainSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := BlockConfig{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}
	return data.ConsensusTime
}
