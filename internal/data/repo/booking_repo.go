package repo

import (
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
)


type BookingRepository interface {
	//GetAll(filters ...BookingLookupFilterOption)([]*entities.Booking, error)
}

type StoreBookingRepository struct {
	bookingStore stores.BookingStore
}

func NewStoreBookingRepository(bookingStore stores.BookingStore) BookingRepository {
	return &StoreBookingRepository{
		bookingStore: bookingStore,
	}
}
/*
func (r * StoreBookingRepository) GetAll(filters ...BookingLookupFilterOption) ([]*entities.Booking, error){
	lookup := DefaultBookingLookupFilter()
	for _, f := range filters {
		f(lookup)
	}

	filter :=  LookupToFilterBookings(lookup)
	bookings, err := r.bookingStore.GetAll(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var result []*entities.Booking
	for _, booking := range bookings{
		populated, err := r.
	}
}
*/