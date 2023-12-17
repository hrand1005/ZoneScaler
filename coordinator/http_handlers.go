package coordinator

import (
	"encoding/json"
	"github.com/phuslu/log"
	"net/http"

	"github.com/hrand1005/ZoneScaler/common"
)

// SetupHTTPHandlers sets up the HTTP routes
func SetupHTTPHandlers(coordinator *Coordinator) {
	http.HandleFunc("/addNode", func(w http.ResponseWriter, r *http.Request) {
		handleNodeAddition(coordinator, w, r)
	})
	http.HandleFunc("/removeNode", func(w http.ResponseWriter, r *http.Request) {
		handleNodeRemoval(coordinator, w, r)
	})
}

func handleNodeAddition(c *Coordinator, w http.ResponseWriter, r *http.Request) {
	var node common.GameNode
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("Failed to decode node")
		return
	}
	c.AddNode(&node)
	log.Info().Str("node_id", node.ID).Msg("Node added")
}

func handleNodeRemoval(c *Coordinator, w http.ResponseWriter, r *http.Request) {
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		http.Error(w, "Node ID is required", http.StatusBadRequest)
		return
	}
	c.RemoveNode(nodeID)
	log.Info().Str("node_id", nodeID).Msg("Node removed")
}
