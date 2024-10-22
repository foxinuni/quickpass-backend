package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type EventsRouter struct {
	eventsController *controllers.EventsController
}

func NewEventsRouter(eventsController *controllers.EventsController) *EventsRouter{
	return &EventsRouter{eventsController: eventsController}
}

func (er *EventsRouter) RegisterRoutes(echo *echo.Echo){
	eventsGroup := echo.Group("/events")

	//registering routes
	eventsGroup.GET("", er.eventsController.GetAll)
	eventsGroup.GET("/:id", er.eventsController.GetOccasionsFromEvent)
	eventsGroup.POST("/:id/invite", er.eventsController.InviteUsersToEvent)
	eventsGroup.POST("/:id/invite/all", er.eventsController.InviteAllUsersToEvent)
}