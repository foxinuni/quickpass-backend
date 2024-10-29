package controllers

import (
	"net/http"
	"strconv"

	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type LogsController struct {
	actionsService services.ActionsService
}

// constructor or initializer for the struct
func NewLogsController(actionsService services.ActionsService) *LogsController {
	return &LogsController{
		actionsService: actionsService,
	}
}

func (mbc *LogsController) GetAllEventLogs(c echo.Context) error {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event ID")
	}
	logs, err := mbc.actionsService.GetLogs(&eventId, nil)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, logs)
}

func (mbc *LogsController) GetAllBookingLogs(c echo.Context) error {
	bookingId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid booking ID")
	}
	logs, err := mbc.actionsService.GetLogs( nil, &bookingId)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, logs)
}

