//go:build wireinject
// +build wireinject

package main

import (
	"github.com/foxinuni/quickpass-backend/internal/core"
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation"
	"github.com/foxinuni/quickpass-backend/internal/presentation/auth"
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/routes"
	"github.com/google/wire"
)

/*
	MENSAJE PARA CUALQUIERA MIRANDO: ESTE ES UN ARCHIVO DE GOOGLE WIRE, PARA INYECTAR DEPENDENCIAS EN EL PROYECTO.
	SI NO TIENES WIRE INSTALADO, INSTÁLALO CON EL SIGUIENTE COMANDO:
	> go get github.com/google/wire/cmd/wire

	DESPUÉS DE INSTALAR WIRE, EJECUTA EL SIGUIENTE COMANDO PARA GENERAR EL CÓDIGO DEPENDIENDO DE ESTE ARCHIVO:
	> wire ./...

	ESTO GENERARÁ UN ARCHIVO WIRE_GEN.GO QUE DEBERÁS INCLUIR EN TU PROYECTO.
*/

// First we create the store factory to create all stores
var backendSet = wire.NewSet(
	// API server
	presentation.NewQuickpassAPI,

	// Routers
	routes.NewAuthRouter,
	routes.NewMyOccasionsRouter,
	routes.NewMyBookingsRouter,

	// Controllers
	controllers.NewLoginController,
	controllers.NewMyOccasionsController,
	controllers.NewMyBookingsController,

	// Strategies
	auth.NewAuthServiceStrategy,

	// Services
	services.NewRepoOccassionsService,
	services.NewJwtAuthService,
	services.NewRepoStateService,
	services.NewRepoBookingsService,

	// Repositories
	repo.NewStoreOccasionRepository,
	repo.NewStoreUserRepository,
	repo.NewStoreSessionRepository,
	repo.NewStoreStateRepository,

	// Stores
	core.BuildUserStore,
	core.BuildSessionStore,
	core.BuildEventStore,
	core.BuildBookingStore,
	core.BuildStateStore,
	core.BuildOccasionStore,
	core.BuildAccomoStore,
	core.BuildLogStore,

	// Store factory
	core.NewPostgresStoreFactory, // This is the factory that creates all stores
)

func BootstrapServer(options *ApplicationConfig) (*presentation.QuickpassAPI, error) {
	// Annotate options with wire.Build
	wire.Build(
		backendSet,
		wire.Bind(new(presentation.QuickpassAPIOptions), new(*ApplicationConfig)),
		wire.Bind(new(services.JwtAuthServiceOptions), new(*ApplicationConfig)),
		wire.Bind(new(core.PgStoreFactoryOptions), new(*ApplicationConfig)),
	)

	return nil, nil
}
