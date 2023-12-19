package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/phuslu/log"

	"github.com/hrand1005/ZoneScaler/common"
	"github.com/hrand1005/ZoneScaler/worker"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal().Msg("Usage: ./coordinator <JSON config>")
	}

	conf, err := worker.LoadConfig(os.Args[1])
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to load config")
	}

	node := common.GameNode{
		ID:            uuid.NewString(),
		Address:       fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Load:          0,
		Regions:       []string{"region1", "region2"},
		IsActive:      true,
		LastHeartbeat: time.Now(),
	}

	coordinatorAddr := fmt.Sprintf("%v:%v", conf.CoordinatorHost, conf.CoordinatorPort)
	err = worker.RegisterWithCoordinator(coordinatorAddr, node)
	if err != nil {
		log.Fatal().Err(err)
		return
	}

	http.HandleFunc("/player", worker.PlayerDataHandler)

	serverAddr := fmt.Sprintf(":%v", conf.Port)
	log.Info().Msgf("Worker HTTP server listening on %s", serverAddr)
	if err = http.ListenAndServe(serverAddr, nil); err != nil {
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start HTTP server")
		}
	}
}
