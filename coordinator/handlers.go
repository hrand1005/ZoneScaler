package coordinator

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/phuslu/log"

	"github.com/hrand1005/ZoneScaler/common"
)

func (c *Coordinator) NodesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateNodeHandler(w, r)
	case http.MethodGet:
		c.ReadNodesHandler(w, r)
	case http.MethodDelete:
		c.DeleteNodeHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (c *Coordinator) CreateNodeHandler(w http.ResponseWriter, r *http.Request) {
	var node common.GameNode
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Error().Err(err).Msg("Failed to decode node")
		return
	}
	c.AddNode(&node)
	log.Info().Str("node_id", node.ID).Msg("Node added")
}

func (c *Coordinator) ReadNodesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("coordinator/templates/node.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nodes := c.CopyNodes()
	if err := tmpl.Execute(w, nodes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Coordinator) DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	nodeID := r.URL.Query().Get("id")
	if nodeID == "" {
		http.Error(w, "Node ID is required", http.StatusBadRequest)
		return
	}
	c.RemoveNode(nodeID)
	log.Info().Str("node_id", nodeID).Msg("Node removed")
}
