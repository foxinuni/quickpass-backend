package presentation

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/foxinuni/quickpass-backend/internal/presentation/routes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

type QuickpassAPIOptions interface {
	GetListenAddress() string
}

type QuickpassAPI struct {
	options           QuickpassAPIOptions
	authRouter        *routes.AuthRouter
	myOccasionsRouter *routes.MyOccasionsRouter
	myBookingsRouter  *routes.MyBookingsRouter
	myEventsRouter    *routes.MyEventsRouter
	actionsRouter     *routes.ActionsRouter
	eventsRouter      *routes.EventsRouter
	sessionsRouter    *routes.SessionsRouter
	occasionsRouter   *routes.OccasionsRouter
	bookingsRouter    *routes.BookingsRouter
	logsRouter			*routes.LogsRouter
	webSockerRouter 	*routes.WebSocketRouter
}

func NewQuickpassAPI(
	options QuickpassAPIOptions,
	authRouter *routes.AuthRouter,
	myOccasionsRouter *routes.MyOccasionsRouter,
	myBookingsRouter *routes.MyBookingsRouter,
	myEventsRouter *routes.MyEventsRouter,
	actionsRouter *routes.ActionsRouter,
	eventsRouter *routes.EventsRouter,
	sessionsRouter *routes.SessionsRouter,
	occasionsRouter *routes.OccasionsRouter,
	bookingsRouter *routes.BookingsRouter,
	logsRouter			*routes.LogsRouter,
	webSockerRouter 	*routes.WebSocketRouter,
) *QuickpassAPI {
	return &QuickpassAPI{
		options:           options,
		authRouter:        authRouter,
		myOccasionsRouter: myOccasionsRouter,
		myBookingsRouter:  myBookingsRouter,
		myEventsRouter:    myEventsRouter,
		actionsRouter:     actionsRouter,
		eventsRouter:      eventsRouter,
		sessionsRouter:    sessionsRouter,
		occasionsRouter:   occasionsRouter,
		bookingsRouter:    bookingsRouter,
		logsRouter: 	logsRouter,
		webSockerRouter: 	webSockerRouter,
	}
}

func (api *QuickpassAPI) Listen() error {
	// Create a new Echo instance
	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost", "http://127.0.0.1"}, // Allow specific base origins
		AllowOriginFunc: func(origin string) (bool, error) {
			// Allow any localhost or 127.0.0.1 with any port
			return origin == "http://localhost" || origin == "http://127.0.0.1" || isLocalhostWithPort(origin), nil
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// Hide the banner and port
	app.HideBanner = true
	app.HidePort = true

	// Use Go-Playground's validator for DTOs
	app.Validator = &CustomValidator{validator: validator.New()}

	// Register the middlewares
	app.Use(middlewares.RequestLogMiddleware)
	app.Use(middlewares.ErrorHandlerMiddleware)

	// Health check route
	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// Register the routes
	api.authRouter.RegisterRoutes(app)
	api.myOccasionsRouter.RegisterRoutes(app)
	api.myBookingsRouter.RegisterRoutes(app)
	api.myEventsRouter.RegisterRoutes(app)
	api.actionsRouter.RegisterRoutes(app)
	api.eventsRouter.RegisterRoutes(app)
	api.sessionsRouter.RegisterRoutes(app)
	api.occasionsRouter.RegisterRoutes(app)
	api.bookingsRouter.RegisterRoutes(app)
	api.logsRouter.RegisterRoutes(app)
	api.webSockerRouter.RegisterRoutes(app)

	log.Info().Msgf("HTTP server is now listening on %s", api.options.GetListenAddress())
	if err := app.Start(api.options.GetListenAddress()); err != nil {
		return err
	}

	return nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(dto interface{}) error {
	if err := cv.validator.Struct(dto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func isLocalhostWithPort(origin string) bool {
	return len(origin) >= 17 && origin[:17] == "http://localhost:" || len(origin) >= 18 && origin[:18] == "http://127.0.0.1:"
}
