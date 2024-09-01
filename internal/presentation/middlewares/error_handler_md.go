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

		// Setup the default error response
		message := err.Error()
		code := 500

		// Check if it's an HTTP error
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code

			if msg, ok := he.Message.(string); ok {
				message = msg
			} else {
				message = http.StatusText(code)
			}
		}

		// Log the error
		log.Error().Err(err).Msgf("Caught %q (%d) when handling request!", http.StatusText(code), code)

		// Return the error
		return c.JSON(code, map[string]interface{}{
			"status":  code,
			"error":   http.StatusText(code),
			"message": message,
		})
	}
}
