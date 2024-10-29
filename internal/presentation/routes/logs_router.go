package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type LogsRouter struct {
	logsController *controllers.LogsController
}

func NewLogsRouter(logsController *controllers.LogsController) *LogsRouter {
	return &LogsRouter{
		logsController: logsController,
	}
}

func (r *LogsRouter) RegisterRoutes(echo *echo.Echo) {
	echo.GET("/logs/events/:id", r.logsController.GetAllEventLogs)
	echo.GET("/logs/bookings/:id", r.logsController.GetAllBookingLogs)
}
