package presentation

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/foxinuni/quickpass-backend/internal/presentation/routes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
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
	eventsRouter 		*routes.EventsRouter
}

func NewQuickpassAPI(
	options QuickpassAPIOptions,
	authRouter *routes.AuthRouter,
	myOccasionsRouter *routes.MyOccasionsRouter,
	myBookingsRouter *routes.MyBookingsRouter,
	myEventsRouter *routes.MyEventsRouter,
	actionsRouter *routes.ActionsRouter,
	eventsRouter *routes.EventsRouter,
) *QuickpassAPI {
	return &QuickpassAPI{
		options:           options,
		authRouter:        authRouter,
		myOccasionsRouter: myOccasionsRouter,
		myBookingsRouter:  myBookingsRouter,
		myEventsRouter:    myEventsRouter,
		actionsRouter:     actionsRouter,
		eventsRouter: eventsRouter,
	}
}

func (api *QuickpassAPI) Listen() error {
	// Create a new Echo instance
	app := echo.New()

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
