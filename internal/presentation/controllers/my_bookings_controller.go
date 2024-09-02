package controllers

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type MyBookingsController struct {
	bookingService services.BookingsService
}

//constructor or initializer for the struct
func NewMyBookingsController(bookingService services.BookingsService) *MyBookingsController{
	return &MyBookingsController{
		bookingService: bookingService,
	}
}

func(mbc *MyBookingsController) GetAll(c echo.Context) error {
	//getting the user from the request
	user, ok := c.Get("user").(*entities.User)
	//if there's no user then return error
	if !ok{
		return echo.NewHTTPError(http.StatusInternalServerError, "user required but not found")
	}

	//getting bookings from the service
	bookings, err := mbc.bookingService.GetBookingsForUser(user)
	if err != nil{
		return err
	}

	type GetAllResponse struct {
		Bookings []*entities.Booking `json:"bookings"`
	}

	return c.JSON(http.StatusOK, GetAllResponse{Bookings: bookings})

}