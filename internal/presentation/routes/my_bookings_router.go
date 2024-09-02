package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type MyBookingsRouter struct {
	myBookingsController *controllers.MyBookingsController
	authStrategy         middlewares.AuthStrategy
}

func NewMyBookingsRouter(myBookingsController *controllers.MyBookingsController, authStrategy middlewares.AuthStrategy) *MyBookingsRouter {
	return &MyBookingsRouter{
		myBookingsController: myBookingsController,
		authStrategy:         authStrategy,
	}
}

func (mbr *MyBookingsRouter) RegisterRoutes(echo *echo.Echo) {
	//creating group
	myBookingsGroup := echo.Group("/my_bookings", middlewares.AuthMiddleware(mbr.authStrategy))

	//suscribing controller
	myBookingsGroup.GET("", mbr.myBookingsController.GetAll)
}
