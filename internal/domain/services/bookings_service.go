package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type BookingsService interface {
	GetAllBookings() ([]*entities.Booking, error)
	GetBookingsForUser(user *entities.User) ([]*entities.Occasion, error)
}

var _ BookingsService = (*RepoBookingsService)(nil)

type RepoBookingsService struct {
	bookingRepo  repo.BookingRepository
	occasionRepo repo.OccasionRepository
}

func NewRepoBookingsService(
	occasionRepo repo.OccasionRepository,
	bookingRepo repo.BookingRepository,
) BookingsService {
	return &RepoBookingsService{
		occasionRepo: occasionRepo,
		bookingRepo:  bookingRepo,
	}
}

func (s *RepoBookingsService) GetAllBookings() ([]*entities.Booking, error) {
	// occasion for type false means bookings
	bookings, err := s.bookingRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// if there was no error then it returns the bookings
	return bookings, nil
}

func (s *RepoBookingsService) GetBookingsForUser(user *entities.User) ([]*entities.Occasion, error) {
	// occasion for type false means occasions
	occasions, err := s.occasionRepo.GetAll(repo.OccasionForUser(user), repo.OccasionForType(false))
	if err != nil {
		return nil, err
	}

	// if there was no error then it returns the bookings
	return occasions, nil
}
