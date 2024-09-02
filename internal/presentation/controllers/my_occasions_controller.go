package controllers

import (
	"net/http"
	"strconv"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation/dtos"
	"github.com/labstack/echo/v4"
)

type MyOccasionsController struct {
	occasionService services.OccassionService
}

func NewMyOccasionsController(occasionService services.OccassionService) *MyOccasionsController {
	return &MyOccasionsController{
		occasionService: occasionService,
	}
}

func (moc *MyOccasionsController) GetAll(c echo.Context) error {
	var active bool
	if c.QueryParam("active") != "" {
		if a, err := strconv.ParseBool(c.QueryParam("active")); err == nil {
			active = a
		} else {
			return echo.NewHTTPError(400, "invalid active query parameter")
		}
	}

	user, ok := c.Get("user").(*entities.User)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	// Get the user from the context
	occasions, err := moc.occasionService.GetOccasionsForUser(user, active)
	if err != nil {
		return err
	}

	type GetAllResponse struct {
		Occasions []*entities.Occasion `json:"occasions"`
	}

	return c.JSON(http.StatusOK, GetAllResponse{Occasions: occasions})
}

func (moc *MyOccasionsController) Get(c echo.Context) error {
	// Get the user from the context
	user, ok := c.Get("user").(*entities.User)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	// Get the user's occasion by ID
	occasionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid occasion ID")
	}

	// Return the user's occasion
	occasion, err := moc.occasionService.GetOccasionForUsersWithId(user, occasionID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, occasion)
}

func (moc *MyOccasionsController) Update(c echo.Context) error {
	// Get the user from the context
	user, ok := c.Get("user").(*entities.User)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	// Get the user's occasion by ID
	occasionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid occasion ID")
	}

	// Parse the request body
	var patchDTO dtos.PatchMyOccasionDTO
	if err := c.Bind(&patchDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(patchDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Update the user's occasion
	occasion, err := moc.occasionService.ConfirmOccasionForUser(user, occasionID, patchDTO.Confirming)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, occasion)
}
