package controllers

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type MyEventsController struct {
	eventsService services.EventsService
}

// constructor or initializer for the struct
func NewMyEventsController(eventsService services.EventsService) *MyEventsController {
	return &MyEventsController{
		eventsService: eventsService,
	}
}

func (mec *MyEventsController) GetAll(c echo.Context) error {
	//getting the user from the request
	user, ok := c.Get("user").(*entities.User)
	//if there's no user then return error
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	//getting bookings from the service
	bookings, err := mec.eventsService.GetEventsForUser(user)
	if err != nil {
		return err
	}

	type GetAllResponse struct {
		Bookings []*entities.Occasion `json:"occasions"`
	}

	return c.JSON(http.StatusOK, GetAllResponse{Bookings: bookings})

}
