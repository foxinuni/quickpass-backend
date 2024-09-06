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
	routes.NewMyEventsRouter,
	routes.NewActionsRouter,
	routes.NewEventsRouter,
	routes.NewSessionsRouter,

	// Controllers
	controllers.NewLoginController,
	controllers.NewMyOccasionsController,
	controllers.NewMyBookingsController,
	controllers.NewMyEventsController,
	controllers.NewActionsController,
	controllers.NewEventsController,
	controllers.NewSessionController,

	// Strategies
	auth.NewAuthServiceStrategy,

	// Services
	services.NewExcelImportService,
	services.NewJwtAuthService,
	services.NewSendgridEmailService,
	services.NewRepoSessionService,
	services.NewRepoOccassionsService,
	services.NewRepoStateService,
	services.NewRepoBookingsService,
	services.NewRepoActionsService,
	services.NewRepoEventsService,

	// Repositories
	repo.NewStoreOccasionRepository,
	repo.NewStoreUserRepository,
	repo.NewStoreSessionRepository,
	repo.NewStoreStateRepository,
	repo.NewStoreLogRepository,
	repo.NewStoreEventRepository,
	repo.NewStoreAccomodationRepository,
	repo.NewStoreBookingRepository,

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

func BootstrapServer(options *core.ApplicationConfig) (*presentation.QuickpassAPI, error) {
	// Annotate options with wire.Build
	wire.Build(
		backendSet,
		wire.Bind(new(presentation.QuickpassAPIOptions), new(*core.ApplicationConfig)),
		wire.Bind(new(services.SendgridEmailServiceOptions), new(*core.ApplicationConfig)),
		wire.Bind(new(services.JwtAuthServiceOptions), new(*core.ApplicationConfig)),
		wire.Bind(new(core.PgStoreFactoryOptions), new(*core.ApplicationConfig)),
	)

	return nil, nil
}
