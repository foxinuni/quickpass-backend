package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type BookingsService interface {
	GetBookingsForUser(user *entities.User) ([]*entities.Booking, error)
}

var _ BookingsService = (*RepoBookingsService)(nil)

type RepoBookingsService struct {
	bookingRepo repo.BookingRepository
}

func NewRepoBookingsService(bookingRepo repo.BookingRepository) BookingsService{
	return &RepoBookingsService{
		bookingRepo: bookingRepo,
	}
}

func (s *RepoBookingsService) GetBookingsForUser(user * entities.User) ([]*entities.Booking, error){
	bookings, err := s.bookingRepo.GetAll(repo.BookingForUser(user))

	if err != nil {
		return nil, err
	}

	//if there was no error then it returns the bookings
	return bookings, nil
}