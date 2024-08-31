package presentation

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type QuickpassAPIOptions interface {
	GetListenAddress() string
}

type QuickpassAPI struct {
	options QuickpassAPIOptions
}

func NewQuickpassAPI(options QuickpassAPIOptions) *QuickpassAPI {
	return &QuickpassAPI{options: options}
}

func (api *QuickpassAPI) Listen() error {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	log.Info().Msgf("HTTP server is now listening on %s", api.options.GetListenAddress())
	if err := app.Start(api.options.GetListenAddress()); err != nil {
		return err
	}

	return nil
}
