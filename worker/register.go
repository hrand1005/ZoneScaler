package worker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hrand1005/ZoneScaler/common"
	"github.com/phuslu/log"
	"net/http"
)

// RegisterWithCoordinator registers the worker with the coordinator
func RegisterWithCoordinator(coordinatorURL string, node common.GameNode) error {
	// Convert node to JSON
	jsonData, err := json.Marshal(node)
	if err != nil {
		log.Error().Msgf("failed to marshal worker info: %v", err)
		return err
	}

	// Make a POST request to the coordinator to register the worker
	resp, err := http.Post(coordinatorURL+"/addNode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Msgf("failed to register with coordinator: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("registration failed. Coordinator responded with status code: %d", resp.StatusCode)
		log.Error().Int("status_code", resp.StatusCode).Msg(msg)
		return errors.New(msg)
	}

	log.Info().Msg("Successfully registered with the coordinator")
	return nil
}
