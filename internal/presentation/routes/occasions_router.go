package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type OccasionsRouter struct {
	occasionsController *controllers.OccasionsController
	authStrategy        middlewares.AuthStrategy
}

func NewOccasionsRouter(occasionsController *controllers.OccasionsController, authStrategy middlewares.AuthStrategy) *OccasionsRouter {
	return &OccasionsRouter{
		occasionsController: occasionsController,
		authStrategy:        authStrategy,
	}
}

func (or *OccasionsRouter) RegisterRoutes(echo *echo.Echo) {
	// Register the create occasion route
	echo.POST("/occasions", or.occasionsController.Create, middlewares.AuthMiddleware(or.authStrategy))
}
