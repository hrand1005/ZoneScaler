package main

import (
	"encoding/json"
	"fmt"
	"github.com/phuslu/log"
	"net/http"
	"os"
	"time"

	"github.com/hrand1005/ZoneScaler/common"
	"github.com/hrand1005/ZoneScaler/worker"
)

func main() {
	// Check if the coordinator URL is provided as an environment variable
	coordinatorURL := os.Getenv("COORDINATOR_URL")
	if coordinatorURL == "" {
		log.Fatal().Msg("COORDINATOR_URL environment variable is not set")
	}

	// Create a GameNode instance
	node := common.GameNode{
		ID:            "unique-node-id",
		Address:       "127.0.0.1:8081", // Adjust the address accordingly
		Load:          0,
		Regions:       []string{"region1", "region2"},
		IsActive:      true,
		LastHeartbeat: time.Now(),
	}

	// Register with the coordinator
	err := worker.RegisterWithCoordinator(coordinatorURL, node)
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	// Define the HTTP server and endpoint for receiving player data
	http.HandleFunc("/player", PlayerDataHandler)

	// Start the HTTP server on port 8081
	serverAddr := ":8081"
	fmt.Printf("Worker HTTP server listening on %s...\n", serverAddr)
	err = http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal().Msgf("Error starting HTTP server: %s", err)
	}
}

// PlayerDataHandler handles player data requests
func PlayerDataHandler(w http.ResponseWriter, r *http.Request) {
	var playerData common.PlayerData

	// Decode JSON request body into PlayerData struct
	err := json.NewDecoder(r.Body).Decode(&playerData)
	if err != nil {
		http.Error(w, "Failed to parse player data", http.StatusBadRequest)
		return
	}

	// Process the player data (you can add your own logic here)
	log.Info().Msgf("Received player data: %+v\n", playerData)

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	log.Info().Msg("Player data received successfully")
}
