package p2pCommunication

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"time"
	"trustedStorage/stateWorker"
)

type Pair struct {
	data []byte
	id   peer.ID
}
type ReqQueue []Pair

var RequestQueueIns *ReqQueue

func InitQueue() *ReqQueue {
	var q ReqQueue
	return &q
}

func (q *ReqQueue) StartNodeQueueProcess() {
	log.Println("Starting queue process")
	for {
		if stateWorker.GetCurrentNodeState() != "Working" {
			time.Sleep(time.Second * 5)
			continue
		}
		if len(*q) != 0 {
			log.Println("Executing request from a queue")
			NodeActionDecision((*q)[0].data, (*q)[0].id, true)
			(*q)[0] = (*q)[len(*q)-1]
			*q = (*q)[:len(*q)-1]
		} else {
			time.Sleep(time.Second * 10)
		}
	}
}

func (q *ReqQueue) AddRequest(data_ []byte, id_ peer.ID) {
	mu.Lock()
	defer mu.Unlock()

	*q = append(*q, Pair{data: data_, id: id_})
}
