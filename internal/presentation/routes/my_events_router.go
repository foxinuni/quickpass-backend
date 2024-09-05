package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type MyEventsRouter struct {
	myEventsController *controllers.MyEventsController
	authStrategy       middlewares.AuthStrategy
}

func NewMyEventsRouter(myEventsController *controllers.MyEventsController, authStrategy middlewares.AuthStrategy) *MyEventsRouter {
	return &MyEventsRouter{
		myEventsController: myEventsController,
		authStrategy:       authStrategy,
	}
}

func (mbr *MyEventsRouter) RegisterRoutes(echo *echo.Echo) {
	//creating group
	myBookingsGroup := echo.Group("/my_events", middlewares.AuthMiddleware(mbr.authStrategy))

	//suscribing controller
	myBookingsGroup.GET("", mbr.myEventsController.GetAll)
}
