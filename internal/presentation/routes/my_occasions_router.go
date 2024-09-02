package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type MyOccasionsRouter struct {
	myOccasionsController *controllers.MyOccasionsController
	authStrategy          middlewares.AuthStrategy
}

func NewMyOccasionsRouter(myOccasionsController *controllers.MyOccasionsController, authStrategy middlewares.AuthStrategy) *MyOccasionsRouter {
	return &MyOccasionsRouter{
		myOccasionsController: myOccasionsController,
		authStrategy:          authStrategy,
	}
}

func (mor *MyOccasionsRouter) RegisterRoutes(echo *echo.Echo) {
	// Create a new group for the my occasions routes
	myOccasionsGroup := echo.Group("/my-occasions", middlewares.AuthMiddleware(mor.authStrategy))

	// Register the get all my occasions route
	myOccasionsGroup.GET("", mor.myOccasionsController.GetAll)

	// Register the get my occasion by ID route
	myOccasionsGroup.GET("/:id", mor.myOccasionsController.Get)
}
