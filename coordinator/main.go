package main

import (
	"github.com/phuslu/log"
	"net/http"
)

func main() {
	coordinator := NewCoordinator()

	// Setup HTTP handlers
	SetupHTTPHandlers(coordinator)

	// Start heartbeat checker in a separate goroutine
	go StartHeartbeatChecker(coordinator)

	log.Info().Msg("Coordinator server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
