package stateWorker

import (
	"sync"
	"time"
)

var NodeState string
var mu sync.Mutex

func GetCurrentNodeState() string {
	defer mu.Unlock()
	mu.Lock()
	return NodeState
}

func SetNodeState(newState string) {
	defer mu.Unlock()
	mu.Lock()

	//todo check if given state is correct
	NodeState = newState
}

func WaitForStateChanged(startState string, targetState string, waitTime int) bool {
	curTime := time.Now().Unix()

	for GetCurrentNodeState() == startState {
		time.Sleep(time.Second)
		if GetCurrentNodeState() == targetState {
			return true
		} else if time.Now().Unix() > curTime+int64(waitTime) {
			return false
		}
	}
	return false
}
