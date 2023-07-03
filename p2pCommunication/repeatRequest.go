package p2pCommunication

import (
	"errors"
	"log"
	"time"
)

func WaitForDelivery(peerID string, requestData []byte, timeout int64) error {
	var err error
	startTime := time.Now().Unix()

	for {
		err = SendDataToConnectedPeerByPeerID(Node, peerID, requestData)
		if err != nil {
			log.Println("Request wasn't delivered, repeating")
			log.Println(err)
		} else {
			return nil
		}
		time.Sleep(time.Second * 4)

		if time.Now().Unix()-startTime > timeout {
			return errors.New("timeout has expired")
		}

	}
}

func RepeatDelivery(timeout int64) {

}
