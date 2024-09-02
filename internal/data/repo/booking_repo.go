package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)


type BookingLookupFilterOption func(*BookingLookupFilter)

type BookingLookupFilter struct {
	User *entities.User
}

func DefaultBookingLookupFilter() *BookingLookupFilter{
	return &BookingLookupFilter{}
}

func BookingForUser(user *entities.User) BookingLookupFilterOption{
	return func(f *BookingLookupFilter){
		f.User = user
	}
}

func LookupToFilterBookings(lookup *BookingLookupFilter) stores.BookingFilter{
	var bookingFilter stores.BookingFilter
	if lookup.User != nil {
		bookingFilter.UserID =&lookup.User.UserID
	}
	return bookingFilter
}



type BookingRepository interface {
	GetAll(filters ...BookingLookupFilterOption)([]*entities.Booking, error)
}

type StoreBookingRepository struct {
	bookingStore stores.BookingStore
}

func NewStoreBookingRepository(bookingStore stores.BookingStore) BookingRepository {
	return &StoreBookingRepository{
		bookingStore: bookingStore,
	}
}

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
