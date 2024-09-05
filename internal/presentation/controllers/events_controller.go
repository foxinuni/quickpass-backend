package controllers

import (
	"net/http"
	"strconv"

	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation/dtos"
	"github.com/labstack/echo/v4"
)

type EventsController struct {
	eventsService services.EventsService
}

func NewEventsController(eventsService services.EventsService) *EventsController{
	return &EventsController{
		eventsService: eventsService,
	}
}


func (ec *EventsController) GetAll(c echo.Context) error {
	events, err := ec.eventsService.GetAllEvents()
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, events )
}

func(ec *EventsController) GetOccasionsFromEvent(c echo.Context) error{

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event ID")
	}
	occasions, err := ec.eventsService.GetOccasionsFromEvent(eventId)
	if err != nil {
		return  echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, occasions)
}

//TO BE IMPLEMENTED
func (ec * EventsController) InviteUsersToEvent(c echo.Context) error {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event ID")
	}
	//Validate DTO of list of occasionID's
	var occasionsID dtos.UserXEvent
	if err := c.Bind(&occasionsID); err != nil {
		return err
	}

	//validate DTO 
	if err := c.Validate(&occasionsID); err != nil {
		return err
	}
	
	number, err := ec.eventsService.InviteUsers(eventId, occasionsID.OccasionsID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"number" : number,
	})
}