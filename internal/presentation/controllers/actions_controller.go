package controllers

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation/dtos"
	"github.com/labstack/echo/v4"
)

type ActionsController struct {
	actionsService services.ActionsService
}

func NewActionsController(actionsService services.ActionsService) *ActionsController {
	return &ActionsController{
		actionsService: actionsService,
	}
}

func (ac *ActionsController) NewAction(c echo.Context) error {
	//getting the user from the request
	user, ok := c.Get("user").(*entities.User)
	//if there's no user then return error
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	// validation of DTO
	var action dtos.ActionDTO
	if err := c.Bind(&action); err != nil {
		return err
	}
	if err := c.Validate(&action); err != nil {
		return err
	}

	// call the actions service
	isInside, err := ac.actionsService.NewAction(user, action.OccasionID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"inside": isInside,
	})
}
