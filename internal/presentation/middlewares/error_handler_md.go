package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		// Get the error code
		code := 500
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		// Log the error
		log.Error().Err(err).Msgf("Caught %q (%d) when handling request!", http.StatusText(code), code)

		// Return the error
		return c.JSON(code, map[string]interface{}{
			"status":  code,
			"message": http.StatusText(code),
			"error":   err.Error(),
		})
	}
}
