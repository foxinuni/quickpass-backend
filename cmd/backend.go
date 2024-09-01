package main

import (
	"os"

	"github.com/foxinuni/quickpass-backend/internal/core"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation"
	"github.com/foxinuni/quickpass-backend/internal/presentation/auth"
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/routes"
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
	// Create store factory
	_, err := core.NewPostgresStoreFactory(config)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create store factory")
	}

	// Create services
	authService := services.NewJwtAuthService(config)

	// Create strategies & middlewares
	authStrategy := auth.NewAuthServiceStrategy(authService)

	// Create controllers
	loginController := controllers.NewLoginController(authService)

	// Create routers
	authRouter := routes.NewAuthRouter(loginController, authStrategy)

	// Create the Quickpass API
	api := presentation.NewQuickpassAPI(config, authRouter)
	if err := api.Listen(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}
