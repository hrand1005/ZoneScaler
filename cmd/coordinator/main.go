package main

import (
	"os"

	"github.com/hrand1005/ZoneScaler/common"
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

	// pass in empty game data for now
	c := coordinator.New(conf, &common.GameData{})
	c.Start()
}
