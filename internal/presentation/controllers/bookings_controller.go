package controllers

import (
	"net/http"
	"strconv"

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

func (bc *BookingsController) InviteOccasion(c echo.Context) error {
	occasionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid event ID")
	}
	number, err := bc.bookingService.InviteOccasion(occasionId)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"number": number,
	})
}


func (bc *BookingsController) InviteAllBookings(c echo.Context) error {
	number, err := bc.bookingService.InviteAllBookings()
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"number": number,
	})
}