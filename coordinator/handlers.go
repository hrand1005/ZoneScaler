package coordinator

import (
	"encoding/json"
	"net/http"

	"github.com/phuslu/log"

	"github.com/hrand1005/ZoneScaler/common"
)

func (c *Coordinator) AddNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node common.GameNode
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("Failed to decode node")
		return
	}
	c.AddNode(&node)
	log.Info().Str("node_id", node.ID).Msg("Node added")
}

func (c *Coordinator) RemoveNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		http.Error(w, "Node ID is required", http.StatusBadRequest)
		return
	}
	c.RemoveNode(nodeID)
	log.Info().Str("node_id", nodeID).Msg("Node removed")
}
