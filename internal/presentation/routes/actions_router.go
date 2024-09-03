package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/foxinuni/quickpass-backend/internal/presentation/middlewares"
	"github.com/labstack/echo/v4"
)

type ActionsRouter struct {
	actionsController *controllers.ActionsController
	authStrategy      middlewares.AuthStrategy
}

func NewActionsRouter(actionsController *controllers.ActionsController, authStrategy middlewares.AuthStrategy) *ActionsRouter {
	return &ActionsRouter{
		actionsController: actionsController,
		authStrategy:      authStrategy,
	}
}

func (ar *ActionsRouter) RegisterRoutes(echo *echo.Echo) {
	//creating group
	actionsGroup := echo.Group("/actions", middlewares.AuthMiddleware(ar.authStrategy))

	//suscribing to endpoint
	actionsGroup.POST("", ar.actionsController.NewAction)
}
