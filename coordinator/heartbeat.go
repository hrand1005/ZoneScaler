package coordinator

import (
	"time"

	"github.com/phuslu/log"
)

// startHeartbeatChecker starts the heartbeat checking loop.
func (c *Coordinator) startHeartbeatChecker() {
	for {
		time.Sleep(30 * time.Second)
		c.checkHeartbeat()
	}
}

func (c *Coordinator) checkHeartbeat() {
	c.nodesMutex.Lock()
	defer c.nodesMutex.Unlock()
	for id, node := range c.nodes {
		if time.Since(node.LastHeartbeat) > time.Minute {
			log.Warn().Str("node_id", id).Msg("Node is inactive. Removing node")
			delete(c.nodes, id)
		}
	}
}
