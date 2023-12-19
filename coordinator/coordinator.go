package coordinator

import (
	"sync"

	"github.com/hrand1005/ZoneScaler/common"
)

// Coordinator manages the game nodes and load balancing.
type Coordinator struct {
	nodes      map[string]*common.GameNode
	nodesMutex sync.RWMutex
}

// New creates a new Coordinator instance.
func New() *Coordinator {
	return &Coordinator{
		nodes: make(map[string]*common.GameNode),
	}
}

// AddNode adds a new game node to the coordinator.
func (c *Coordinator) AddNode(node *common.GameNode) {
	c.nodesMutex.Lock()
	c.nodes[node.ID] = node
	c.nodesMutex.Unlock()
}

// RemoveNode removes a game node from the coordinator.
func (c *Coordinator) RemoveNode(nodeID string) {
	c.nodesMutex.Lock()
	delete(c.nodes, nodeID)
	c.nodesMutex.Unlock()
}

func (c *Coordinator) CopyNodes() map[string]*common.GameNode {
	c.nodesMutex.RLock()

	copiedNodes := make(map[string]*common.GameNode)
	for id, node := range c.nodes {
		copiedNode := node.Copy()
		copiedNodes[id] = copiedNode
	}
	c.nodesMutex.RUnlock()

	return copiedNodes
}
