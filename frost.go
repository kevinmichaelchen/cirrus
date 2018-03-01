package frost

import (
	"time"
	"sync"
	"log"
)

const (
	//NumNodes = 1024
	NumNodes = 10
	GrpcPort = 50051
)

var (
	HeartbeatPeriodicity = getHeartbeatPeriodicity()
)

var AvailableNodeIds = make(chan int, NumNodes)

// maps node IDs to heartbeats
// TODO use int32, since protobuf has no vanilla int?
var Heartbeats map[int]time.Time

type FrostServer struct {
}

// how long an apps can abstain from heartbeat-ing its node ID
// before we consider it stale
func getHeartbeatPeriodicity() time.Duration {
	return time.Second * 10
}

func (f *FrostServer) Run() {
	Heartbeats = make(map[int]time.Time, NumNodes)

	// how often we check for stale node IDs
	staleCheckPeriodicity := time.Second * 5

	var wg sync.WaitGroup
	wg.Add(1)
	go PruneStaleEntries(staleCheckPeriodicity)

	wg.Add(1)
	s := &GrpcServer{port: GrpcPort}
	go s.Run()

	wg.Wait()
}

// PruneStaleEntries checks for stale node IDs.
func PruneStaleEntries(sleepDuration time.Duration) {
	for {
		log.Println("Checking for stale node IDs...")
		for nodeID, heartbeatTime := range Heartbeats {
			if time.Now().After(heartbeatTime.Add(HeartbeatPeriodicity)) {
				log.Printf("Node %d has been requisitioned", nodeID)
				AvailableNodeIds <- nodeID
				delete(Heartbeats, nodeID)
			}
		}
		time.Sleep(sleepDuration)
	}
}