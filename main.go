package main

import (
	"github.com/rs/zerolog/log"

	"github.com/JenswBE/go-commerce/config"
	"github.com/JenswBE/go-commerce/usecases"
)

func main() {
	// Parse config
	svcConfig, err := config.ParseConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config")
	}

	// Start service
	usecases.StartService(svcConfig)
}
