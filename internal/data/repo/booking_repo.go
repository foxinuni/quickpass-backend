package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type BookingRepository interface {
	Create(booking *entities.Booking) error
}

type StoreBookingRepository struct {
	bookingStore stores.BookingStore
}

func NewStoreBookingRepository(bookingStore stores.BookingStore) BookingRepository {
	return &StoreBookingRepository{
		bookingStore: bookingStore,
	}
}

func (r *StoreBookingRepository) Create(booking *entities.Booking) error {
	model := BookingToModel(booking)
	if err := r.bookingStore.Create(context.Background(), model); err != nil {
		return err
	}

	*booking = *ModelToBooking(model, booking.Accomodation)
	return nil
}
