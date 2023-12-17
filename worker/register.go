package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrand1005/ZoneScaler/common"
)

// RegisterWithCoordinator registers the worker with the coordinator
func RegisterWithCoordinator(coordinatorURL string, node common.GameNode) error {
	// Convert node to JSON
	jsonData, err := json.Marshal(node)
	if err != nil {
		return fmt.Errorf("failed to marshal worker info: %v", err)
	}

	// Make a POST request to the coordinator to register the worker
	resp, err := http.Post(coordinatorURL+"/addNode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to register with coordinator: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registration failed. Coordinator responded with status code: %d", resp.StatusCode)
	}

	fmt.Println("Successfully registered with the coordinator")
	return nil
}
