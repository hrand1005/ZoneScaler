package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	coordinator := NewCoordinator()

	// Setup HTTP handlers
	SetupHTTPHandlers(coordinator)

	// Start heartbeat checker in a separate goroutine
	go StartHeartbeatChecker(coordinator)

	fmt.Println("Coordinator server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
