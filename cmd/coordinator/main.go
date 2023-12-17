package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hrand1005/ZoneScaler/coordinator"
)

func main() {
	c := coordinator.New()

	// Setup HTTP handlers
	coordinator.SetupHTTPHandlers(c)

	// Start heartbeat checker in a separate goroutine
	go coordinator.StartHeartbeatChecker(c)

	fmt.Println("Coordinator server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
