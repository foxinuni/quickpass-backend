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
	actionsGroup := echo.Group("/events")

	//registering routes
	actionsGroup.GET("", er.eventsController.GetAll)
	actionsGroup.GET("/:id", er.eventsController.GetOccasionsFromEvent)
	actionsGroup.POST("/:id/invite", er.eventsController.InviteUsersToEvent)
}