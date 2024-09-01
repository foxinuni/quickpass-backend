package middlewares

import (
	"context"
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/labstack/echo/v4"
)

var _ context.Context = context.Background()

type AuthStrategy interface {
	Authenticate(token string) (*entities.Session, error)
}

func AuthMiddleware(authStrategy AuthStrategy) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get the token from the Authorization header
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "API token is required")
			}

			// Authenticate the token
			session, err := authStrategy.Authenticate(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			// Set the session and user in the context
			c.Set("session", session)
			c.Set("user", session.GetUser())

			// Call the next handler
			return next(c)
		}
	}
}
