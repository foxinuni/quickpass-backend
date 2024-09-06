package controllers

import (
	"net/http"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/labstack/echo/v4"
)

type BookingsController struct {
	bookingService services.BookingsService
}

// constructor or initializer for the struct
func NewBookingsController(bookingService services.BookingsService) *BookingsController {
	return &BookingsController{
		bookingService: bookingService,
	}
}

// GetAll returns all bookings
func (bc *BookingsController) GetAll(c echo.Context) error {
	// getting bookings from the service
	bookings, err := bc.bookingService.GetAllBookings()
	if err != nil {
		return err
	}

	type GetAllResponse struct {
		Bookings []*entities.Booking `json:"bookings"`
	}

	return c.JSON(http.StatusOK, GetAllResponse{Bookings: bookings})
}
