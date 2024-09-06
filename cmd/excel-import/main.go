//go:build wireinject
// +build wireinject

package main

import (
	"flag"
	"os"

	"github.com/foxinuni/quickpass-backend/internal/core"
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	file   string
	config *core.ApplicationConfig
)
var importSet = wire.NewSet(
	services.NewExcelImportService,

	// Services
	services.NewRepoStateService,

	// Repositories
	repo.NewStoreOccasionRepository,
	repo.NewStoreUserRepository,
	repo.NewStoreSessionRepository,
	repo.NewStoreStateRepository,
	repo.NewStoreEventRepository,
	repo.NewStoreAccomodationRepository,
	repo.NewStoreBookingRepository,

	// Stores
	core.BuildUserStore,
	core.BuildEventStore,
	core.BuildStateStore,
	core.BuildOccasionStore,
	core.BuildAccomoStore,
	core.BuildBookingStore,
	core.BuildLogStore,

	// Store factory
	core.NewPostgresStoreFactory, // This is the factory that creates all stores
)

func init() {
	// pretty logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	flag.StringVar(&file, "file", "test.xlsx", "Excel file to import")
	flag.Parse()

	// validate file
	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic("Excel file does not exist")
	}

	// load .env file
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No \".env\" file found! Using environment variables.")
	}

	// load configuration
	if c, err := core.LoadConfig(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	} else {
		config = c
	}
}

func main() {
	// open file
	reader, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	// Create service and test
	importService, err := buildImportService(config)
	if err != nil {
		panic(err)
	}

	log.Info().Msgf("Importing from file %q...", file)
	if err := importService.ImportFromFile(reader); err != nil {
		panic(err)
	}
}

func buildImportService(_ *core.ApplicationConfig) (*services.ExcelImportService, error) {
	wire.Build(
		importSet,
		wire.Bind(new(core.PgStoreFactoryOptions), new(*core.ApplicationConfig)),
	)
	return nil, nil
}
