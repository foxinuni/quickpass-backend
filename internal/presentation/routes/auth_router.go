package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type AuthRouter struct {
	loginController *controllers.LoginController
	authStrategy    middlewares.AuthStrategy
}

func NewAuthRouter(loginController *controllers.LoginController, authStrategy middlewares.AuthStrategy) *AuthRouter {
	return &AuthRouter{
		loginController: loginController,
		authStrategy:    authStrategy,
	}
}

func (lr *AuthRouter) RegisterRoutes(echo *echo.Echo) {
	// Create a new group for the login routes
	loginGroup := echo.Group("/auth")

	// Register the login route
	loginGroup.POST("/login", lr.loginController.Login)
	loginGroup.POST("/logout", lr.loginController.Logout, middlewares.AuthMiddleware(lr.authStrategy))
}
