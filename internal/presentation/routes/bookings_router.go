package routes

import (
	"github.com/foxinuni/quickpass-backend/internal/presentation/controllers"
	"github.com/labstack/echo/v4"
)

type BookingsRouter struct {
	bookingsController *controllers.BookingsController
}

func NewBookingsRouter(bookingsController *controllers.BookingsController) *BookingsRouter {
	return &BookingsRouter{
		bookingsController: bookingsController,
	}
}

func (br *BookingsRouter) RegisterRoutes(echo *echo.Echo) {
	//creating group
	bookingsGroup := echo.Group("/bookings")

	//suscribing controller
	bookingsGroup.GET("", br.bookingsController.GetAll)
}
