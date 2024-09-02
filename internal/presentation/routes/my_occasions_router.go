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
	// Register the get all my occasions route
	echo.GET("/my_ocassions", mor.myOccasionsController.GetAll, middlewares.AuthMiddleware(mor.authStrategy))

	// Register the get my occasion by ID route
	echo.GET("/my_ocassions/:id", mor.myOccasionsController.Get, middlewares.AuthMiddleware(mor.authStrategy))

	// Register the patch my occasion route
	echo.PATCH("/my_ocassions/:id", mor.myOccasionsController.Update, middlewares.AuthMiddleware(mor.authStrategy))
}
