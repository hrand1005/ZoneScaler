package coordinator

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/hrand1005/ZoneScaler/common"
	"github.com/phuslu/log"
)

// Coordinator manages the game nodes and load balancing.
type Coordinator struct {
	// TODO: perhaps an http server object should be a field
	// the server could be created in New() from the config
	host      string
	port      int
	staticDir string

	// the game to be distributed and managed by the coordinator
	game            *common.GameData
	partitions      map[*common.Partition]common.PartitionStatus
	partitionsMutex sync.RWMutex
	nodes           map[string]*common.GameNode
	nodesMutex      sync.RWMutex
}

// New creates a new Coordinator instance.
func New(conf Config, g *common.GameData) *Coordinator {
	return &Coordinator{
		host:      conf.Host,
		port:      conf.Port,
		staticDir: conf.StaticDir,

		game:  g,
		nodes: make(map[string]*common.GameNode),
	}
}

func (c *Coordinator) Start() {
	http.Handle("/", http.FileServer(http.Dir(c.staticDir)))
	http.HandleFunc("/nodes", c.NodesHandler)
	// http.HandleFunc("/partitions", c.PartitionsHandler)

	go c.startHeartbeatChecker()

	log.Info().Msgf("Coordinator server started on %d", c.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", c.port), nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}

// AddNode adds a new game node to the coordinator.
func (c *Coordinator) AddNode(node *common.GameNode) {
	c.nodesMutex.Lock()
	defer c.nodesMutex.Unlock()
	c.nodes[node.ID] = node
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
