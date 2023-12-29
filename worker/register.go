package worker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrand1005/ZoneScaler/common"
	"github.com/phuslu/log"
)

// RegisterWithCoordinator registers the worker with the coordinator
func RegisterWithCoordinator(coordinatorURL string, node common.GameNode) error {
	jsonData, err := json.Marshal(node)
	if err != nil {
		log.Error().Msgf("failed to marshal worker info: %v", err)
		return err
	}

	addNodeURL := fmt.Sprintf("http://%v/nodes", coordinatorURL)
	resp, err := http.Post(addNodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("failed to register with coordinator")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("registration failed. Coordinator responded with status code: %d", resp.StatusCode)
		log.Error().Int("status_code", resp.StatusCode).Msg(msg)
		return fmt.Errorf(msg)
	}

	log.Info().Msg("Successfully registered with the coordinator")
	return nil
}
