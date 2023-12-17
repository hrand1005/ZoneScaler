package main

import (
	"sync"
)

// Coordinator manages the game nodes and load balancing
type Coordinator struct {
	nodes      map[string]*GameNode
	nodesMutex sync.RWMutex
}

// NewCoordinator creates a new Coordinator instance
func NewCoordinator() *Coordinator {
	return &Coordinator{
		nodes: make(map[string]*GameNode),
	}
}

// AddNode adds a new game node to the coordinator
func (c *Coordinator) AddNode(node *GameNode) {
	c.nodesMutex.Lock()
	defer c.nodesMutex.Unlock()
	c.nodes[node.ID] = node
}

// RemoveNode removes a game node from the coordinator
func (c *Coordinator) RemoveNode(nodeID string) {
	c.nodesMutex.Lock()
	defer c.nodesMutex.Unlock()
	delete(c.nodes, nodeID)
}
