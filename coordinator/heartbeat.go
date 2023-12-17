package main

import (
	"github.com/phuslu/log"
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
			log.Warn().Str("node_id", id).Msg("Node is inactive. Removing node")
			delete(c.nodes, id)
		}
	}
}
