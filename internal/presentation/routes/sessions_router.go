package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type SessionsRouter struct {
	sessionController *controllers.SessionController
}

func NewSessionsRouter(sessionController *controllers.SessionController) *SessionsRouter {
	return &SessionsRouter{
		sessionController: sessionController,
	}
}

func (r *SessionsRouter) RegisterRoutes(echo *echo.Echo) {
	echo.GET("/sessions", r.sessionController.GetAll)
}
