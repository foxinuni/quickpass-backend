package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var config *ApplicationConfig

func init() {
	// Pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No \".env\" file found! Using environment variables.")
	}

	// load configuration
	if c, err := LoadConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	} else {
		config = c
	}
}

func main() {
	// Create the Quickpass API
	server, err := BootstrapServer(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to bootstrap API")
	}

	if err := server.Listen(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
