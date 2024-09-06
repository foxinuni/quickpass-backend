package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type OccasionsRouter struct {
	occasionsController *controllers.OccasionsController
}

func NewOccasionsRouter(occasionsController *controllers.OccasionsController) *OccasionsRouter {
	return &OccasionsRouter{
		occasionsController: occasionsController,
	}
}

func (or *OccasionsRouter) RegisterRoutes(echo *echo.Echo) {
	// Register the create occasion route
	echo.POST("/occasions", or.occasionsController.Create)
}
