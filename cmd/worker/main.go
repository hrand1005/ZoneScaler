package main

import (
	"os"

	"github.com/phuslu/log"

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
	w := worker.New(conf)
	w.Start()
}
