package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type EventsService interface {
	GetEventsForUser(user *entities.User) ([]*entities.Occasion, error)
}

var _ EventsService = (*RepoEventsService)(nil)

type RepoEventsService struct {
	//bookingRepo repo.BookingRepository
	occasionRepo repo.OccasionRepository
}

func NewRepoEventsService(occasionRepo repo.OccasionRepository) EventsService {
	return &RepoEventsService{
		occasionRepo: occasionRepo,
	}
}

func (s *RepoEventsService) GetEventsForUser(user *entities.User) ([]*entities.Occasion, error) {
	//occasion for type true means events
	bookings, err := s.occasionRepo.GetAll(repo.OccasionForType(true))

	if err != nil {
		return nil, err
	}

	//if there was no error then it returns the bookings
	return bookings, nil
}
