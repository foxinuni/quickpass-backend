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

func DefaultBookingLookupFilter() *BookingLookupFilter {
	return &BookingLookupFilter{}
}

func BookingForUser(user *entities.User) BookingLookupFilterOption {
	return func(f *BookingLookupFilter) {
		f.User = user
	}
}

func BookingLookupToFilter(lookup *BookingLookupFilter) stores.BookingFilter {
	var bookingFilter stores.BookingFilter
	if lookup.User != nil {
		bookingFilter.UserID = &lookup.User.UserID
	}

	return bookingFilter
}

type BookingRepository interface {
	GetAll(filters ...BookingLookupFilterOption) ([]*entities.Booking, error)
	Create(booking *entities.Booking) error
	GetById(id int) (*entities.Booking, error)
}

type StoreBookingRepository struct {
	bookingStore stores.BookingStore
	accomodationStore stores.AccomodationStore
}

func NewStoreBookingRepository(bookingStore stores.BookingStore, accomodationStore stores.AccomodationStore) BookingRepository {
	return &StoreBookingRepository{
		bookingStore: bookingStore,
		accomodationStore: accomodationStore,
	}
}

func (r *StoreBookingRepository) GetAll(filters ...BookingLookupFilterOption) ([]*entities.Booking, error) {
	lookup := DefaultBookingLookupFilter()
	for _, filter := range filters {
		filter(lookup)
	}

	bookings, err := r.bookingStore.GetAll(context.Background(), BookingLookupToFilter(lookup))
	if err != nil {
		return nil, err
	}

	var result []*entities.Booking
	for _, booking := range bookings {
		result = append(result, ModelToBooking(&booking, nil))
	}

	return result, nil
}

func (r *StoreBookingRepository) Create(booking *entities.Booking) error {
	model := BookingToModel(booking)
	if err := r.bookingStore.Create(context.Background(), model); err != nil {
		return err
	}

	*booking = *ModelToBooking(model, booking.Accomodation)
	return nil
}

func (r *StoreBookingRepository) GetById(id int) (*entities.Booking, error) {
	booking, err := r.bookingStore.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	accomodation, err := r.accomodationStore.GetById(context.Background(), booking.AccomodationID)
	if err != nil {
		return nil, err
	}
	return ModelToBooking(
		booking, 
		ModelToAccomodation(accomodation),
	), nil
}