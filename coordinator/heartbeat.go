package coordinator

import (
	"log"
	"time"
)

// StartHeartbeatChecker starts the heartbeat checking loop
func StartHeartbeatChecker(c *Coordinator) {
	for {
		time.Sleep(30 * time.Second)
		checkHeartbeat(c)
	}
}

func checkHeartbeat(c *Coordinator) {
	c.nodesMutex.Lock()
	defer c.nodesMutex.Unlock()
	for id, node := range c.nodes {
		if time.Since(node.LastHeartbeat) > time.Minute {
			log.Printf("Node %s is inactive. Removing node.\n", id)
			delete(c.nodes, id)
		}
	}
}
