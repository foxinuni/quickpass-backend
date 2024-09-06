package controllers

import (
	"net/http"
	"strconv"

	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/foxinuni/quickpass-backend/internal/presentation/dtos"
	"github.com/labstack/echo/v4"
)

type SessionController struct {
	sessionService services.SessionService
}

func NewSessionController(sessionService services.SessionService) *SessionController {
	return &SessionController{
		sessionService: sessionService,
	}
}

func (sc *SessionController) GetAll(c echo.Context) error {
	sessions, err := sc.sessionService.GetAllSessions()
	if err != nil {
		return err
	}

	return c.JSON(200, sessions)
}

func (sc *SessionController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(400, "invalid session ID")
	}

	// Bind data to DTO
	var sessionDTO dtos.SessionPatchDTO
	if err := c.Bind(&sessionDTO); err != nil {
		return err
	}

	// Validate DTO
	if err := c.Validate(&sessionDTO); err != nil {
		return err
	}

	// Call the service
	if err := sc.sessionService.EnableSession(id, sessionDTO.Enabled); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
