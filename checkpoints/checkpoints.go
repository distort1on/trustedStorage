package checkpoints

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"trustedStorage/blockchain"
)

type CheckPoints struct {
	Height    []string
	BlockHash []string
}

var CheckPointsSettingsFileName = "checkpoints/checkpoints.json"

func VerifyCheckPoints() bool {
	content, err := os.ReadFile(CheckPointsSettingsFileName)
	if err != nil {
		log.Println(err)
	}
	data := CheckPoints{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Println(err)
	}

	var inNum int
	var curHash []byte
	for in, el := range data.Height {
		inNum, err = strconv.Atoi(el)
		if err != nil {
			log.Println(err)
		}
		curHash = (*blockchain.BlockChainIns)[inNum].GetBlockHash()
		if hex.EncodeToString(curHash) != data.BlockHash[in] {

			return false
		}

	}
	return true
}
