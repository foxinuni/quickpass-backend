package controllers

import (
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
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

func (c *SessionController) GetAll(echo echo.Context) error {
	sessions, err := c.sessionService.GetAllSessions()
	if err != nil {
		return err
	}

	return echo.JSON(200, sessions)
}
