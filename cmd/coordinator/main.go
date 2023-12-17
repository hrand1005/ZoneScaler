package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hrand1005/ZoneScaler/coordinator"
	"github.com/phuslu/log"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal().Msg("Usage: ./coordinator <JSON config>")
	}

	conf, err := coordinator.LoadConfig(os.Args[1])
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	c := coordinator.New()

	http.HandleFunc("/add-node", c.AddNodeHandler)
	http.HandleFunc("/remove-node", c.RemoveNodeHandler)

	go coordinator.StartHeartbeatChecker(c)

	coordinatorPort := fmt.Sprintf(":%d", conf.Port)
	log.Info().Msgf("Coordinator server started on %s", coordinatorPort)
	if err := http.ListenAndServe(coordinatorPort, nil); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
