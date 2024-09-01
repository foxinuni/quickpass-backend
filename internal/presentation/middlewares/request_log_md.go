package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func RequestLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log the request
		log.Debug().Msgf("%-6s %s from [%s]", c.Request().Method, c.Request().URL.Path, c.RealIP())

		return next(c)
	}
}
