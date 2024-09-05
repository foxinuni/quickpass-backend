package services

import (
	"fmt"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	commonServices "github.com/foxinuni/quickpass-backend/internal/shared/common_services"
)

type EventsService interface {
	GetEventsForUser(user *entities.User) ([]*entities.Occasion, error)
	GetOccasionsFromEvent(eventId int)([]*entities.Occasion, error)
	GetAllEvents()([]*entities.Event, error)
	InviteUsers(eventId int, occasionIds []*int) (int, error)
}

var _ EventsService = (*RepoEventsService)(nil)

type RepoEventsService struct {
	occasionRepo repo.OccasionRepository
	eventRepo repo.EventRepository
	stateService StateService
}

func NewRepoEventsService(occasionRepo repo.OccasionRepository, eventRepo repo.EventRepository, stateService StateService ) EventsService {
	return &RepoEventsService{
		occasionRepo: occasionRepo,
		eventRepo:eventRepo,
		stateService: stateService,
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

func (s *RepoEventsService) InviteUsers(eventId int, occasionIds []*int) (int, error){
	event, err := s.eventRepo.GetById(eventId)
	if err != nil {
		return 0, nil
	}
	var counter int
	for _, id := range occasionIds {
		occasion, err := s.occasionRepo.GetById(*id)
		//IF THIS OCCASION WAS NOT FOUND CONTINUE WITH NEXT ONE
		if err != nil{
			continue
		}
		usersEmail := occasion.GetUser().GetEmail()
		eventName := event.GetName()
		startDate := event.GetStartDate()
		endDate := event.GetEndDate()
		err = commonServices.SendEmail(
			usersEmail,
			fmt.Sprintf("Invitacion evento %s", eventName),
			fmt.Sprintf("Estas cordialmente invitado al evento "+ 
			"%s, recuerda que inicia en %s y finaliza el %s, confirma tu " +
			"asistencia por medio de nuesra aplicacion:", eventName,startDate, endDate),
		)
		//if there was an error with this particular email, continue with next occasion
		if err != nil{
			continue
		}
		//updating state of that occasion to invited
		state, err := s.stateService.GetOrCreateState("invited")
		occasion.SetState(state)
		s.occasionRepo.Update(occasion)
		counter++
	}
	return counter, nil

}
