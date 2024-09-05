package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type EventsService interface {
	GetEventsForUser(user *entities.User) ([]*entities.Occasion, error)
	GetOccasionsFromEvent(eventId int)([]*entities.Occasion, error)
	GetAllEvents()([]*entities.Event, error)
}

var _ EventsService = (*RepoEventsService)(nil)

type RepoEventsService struct {
	occasionRepo repo.OccasionRepository
	eventRepo repo.EventRepository
}

func NewRepoEventsService(occasionRepo repo.OccasionRepository) EventsService {
	return &RepoEventsService{
		occasionRepo: occasionRepo,
	}
}

func (s *RepoEventsService) GetEventsForUser(user *entities.User) ([]*entities.Occasion, error) {
	//occasion for type true means events
	events, err := s.occasionRepo.GetAll(repo.OccasionForType(true), repo.OccasionForUser(user))

	if err != nil {
		return nil, err
	}

	//if there was no error then it returns the bookings
	return events, nil
}

func (s *RepoEventsService) GetOccasionsFromEvent(eventId int)([]*entities.Occasion, error){
	event, err:= s.eventRepo.GetById(eventId)
	if err != nil {
		return nil, err
	}
	occasions, err := s.occasionRepo.GetAll(repo.OccasionForEvent(event))
	if err != nil {
		return nil, err
	}
	return occasions, nil
}

func (s *RepoEventsService) GetAllEvents()([]*entities.Event, error){
	events, err := s.eventRepo.GetAll()
	if err != nil{
		return nil, err
	}
	return events, nil
}
