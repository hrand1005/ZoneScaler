package main

import (
	"github.com/phuslu/log"
	"net/http"

	"github.com/hrand1005/ZoneScaler/coordinator"
)

func main() {
	c := coordinator.New()

	// Setup HTTP handlers
	coordinator.SetupHTTPHandlers(c)

	// Start heartbeat checker in a separate goroutine
	go coordinator.StartHeartbeatChecker(c)

	log.Info().Msg("Coordinator server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
