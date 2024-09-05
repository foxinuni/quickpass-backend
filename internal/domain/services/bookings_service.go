package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type BookingsService interface {
	GetBookingsForUser(user *entities.User) ([]*entities.Occasion, error)
}

var _ BookingsService = (*RepoBookingsService)(nil)

type RepoBookingsService struct {
	occasionRepo repo.OccasionRepository
}

func NewRepoBookingsService(occasionRepo repo.OccasionRepository) BookingsService{
	return &RepoBookingsService{
		occasionRepo: occasionRepo,
	}
}

func (s *RepoBookingsService) GetBookingsForUser(user * entities.User) ([]*entities.Occasion, error){
	//occasion for type false means bookings
	bookings, err := s.occasionRepo.GetAll(repo.OccasionForType(false), repo.OccasionForUser(user))

	if err != nil {
		return nil, err
	}

	//if there was no error then it returns the bookings
	return bookings, nil
}