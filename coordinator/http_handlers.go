package coordinator

import (
	"encoding/json"
	"fmt"
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
		return
	}
	c.AddNode(&node)
	fmt.Fprintf(w, "Node added: %s\n", node.ID)
}

func handleNodeRemoval(c *Coordinator, w http.ResponseWriter, r *http.Request) {
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		http.Error(w, "Node ID is required", http.StatusBadRequest)
		return
	}
	c.RemoveNode(nodeID)
	fmt.Fprintf(w, "Node removed: %s\n", nodeID)
}
