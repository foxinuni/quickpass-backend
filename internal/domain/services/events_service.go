package services

import (
	"fmt"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/rs/zerolog/log"
)

type EventsService interface {
	GetAllEvents() ([]*entities.Event, error)
	GetEventsForUser(user *entities.User) ([]*entities.Occasion, error)
	GetOccasionsFromEvent(eventId int) ([]*entities.Occasion, error)
	InviteUsers(eventId int, occasionIds []int) (int, error)
}

var _ EventsService = (*RepoEventsService)(nil)

type RepoEventsService struct {
	occasionRepo repo.OccasionRepository
	eventRepo    repo.EventRepository
	stateService StateService
	emailService EmailService
}

func NewRepoEventsService(
	occasionRepo repo.OccasionRepository,
	eventRepo repo.EventRepository,
	stateService StateService,
	emailService EmailService,
) EventsService {
	return &RepoEventsService{
		occasionRepo: occasionRepo,
		eventRepo:    eventRepo,
		stateService: stateService,
		emailService: emailService,
	}
}

func (s *RepoEventsService) GetEventsForUser(user *entities.User) ([]*entities.Occasion, error) {
	// occasion for type true means events
	events, err := s.occasionRepo.GetAll(repo.OccasionForType(true), repo.OccasionForUser(user))

	if err != nil {
		return nil, err
	}

	// if there was no error then it returns the bookings
	return events, nil
}

func (s *RepoEventsService) GetOccasionsFromEvent(eventId int) ([]*entities.Occasion, error) {
	// get the event by id
	event, err := s.eventRepo.GetById(eventId)
	if err != nil {
		return nil, err
	}

	// get all the occasions for that event
	occasions, err := s.occasionRepo.GetAll(repo.OccasionForEvent(event))
	if err != nil {
		return nil, err
	}

	return occasions, nil
}

func (s *RepoEventsService) GetAllEvents() ([]*entities.Event, error) {
	// get all the events
	events, err := s.eventRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (s *RepoEventsService) InviteUsers(eventId int, occasionIds []int) (int, error) {
	// get the event by id
	event, err := s.eventRepo.GetById(eventId)
	if err != nil {
		return 0, nil
	}

	// updating state of that occasion to invited
	state, err := s.stateService.GetOrCreateState(StateInvited)
	if err != nil {
		return 0, err
	}

	// send email to all the users
	var counter int
	for _, id := range occasionIds {
		// if this occasion does not exist, continue with the next one
		occasion, err := s.occasionRepo.GetById(id)
		if err != nil {
			log.Warn().Err(err).Msgf("Error getting occasion with id %d", id)
			continue
		}

		// get the user email
		userEmail := occasion.GetUser().GetEmail()
		eventAddress := event.GetAddress()
		eventName := event.GetName()
		startDate := event.GetStartDate()
		endDate := event.GetEndDate()

		// send the email
		err = s.emailService.SendEmail(
			userEmail,
			fmt.Sprintf("Haz sido invitado al evento %s", eventName),
			fmt.Sprintf(`
				Estas cordialmente invitado al evento %q. Este evento se dara en: %q.
				Recuerda que inicia en %s y finaliza el %s, confirma tu asistencia por medio de nuesra aplicacion!
			`, eventName, eventAddress, startDate, endDate),
		)
		if err != nil {
			log.Warn().Err(err).Msgf("Error sending email to %q", userEmail)
			continue
		}

		occasion.SetState(state)
		if err := s.occasionRepo.Update(occasion); err != nil {
			log.Warn().Err(err).Msgf("Error updating occasion with id %d", id)
			continue
		}

		counter++
	}

	return counter, nil
}
