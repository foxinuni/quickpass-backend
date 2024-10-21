package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type WebSocketRouter struct {
	controller *controllers.WebSocketsController
}

func NewWebSocketRouter(controller *controllers.WebSocketsController) *WebSocketRouter {
	return &WebSocketRouter{
		controller: controller,
	}
}

func (r *WebSocketRouter) RegisterRoutes(echo *echo.Echo) {
	echo.GET("/ws/events/:id", r.controller.EventsWebSocketHanlder)

	go controllers.EventBroadcaster()
	//echo.GET("/ws/bookings", handleWebSocket)
}
