package services

import (
	"fmt"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/rs/zerolog/log"
)

type BookingsService interface {
	GetAllBookings() ([]*entities.Booking, error)
	GetBookingsForUser(user *entities.User) ([]*entities.Occasion, error)
	sendInvitation(occasionIds []int) (int, error)
	InviteOccasion(occasionId int) (int, error)
	InviteAllBookings() (int, error)
}

var _ BookingsService = (*RepoBookingsService)(nil)

type RepoBookingsService struct {
	bookingRepo  repo.BookingRepository
	occasionRepo repo.OccasionRepository
	stateService StateService
	emailService EmailService
}

func NewRepoBookingsService(
	occasionRepo repo.OccasionRepository,
	bookingRepo repo.BookingRepository,
	stateService StateService,
	emailService EmailService,
) BookingsService {
	return &RepoBookingsService{
		occasionRepo: occasionRepo,
		bookingRepo:  bookingRepo,
		stateService: stateService,
		emailService: emailService,
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

func (s *RepoBookingsService) sendInvitation(occasionIds []int) (int, error){
	// updating state of that occasion to invited
	state, err := s.stateService.GetOrCreateState(StateInvited)
	if err != nil {
		return 0, err
	}

	// send email to all the users
	var counter int
	for _, id := range occasionIds{
		// if this occasion does not exist, continue with the next one
		occasion, err := s.occasionRepo.GetById(id)
		if err != nil {
			log.Warn().Err(err).Msgf("Error getting occasion with id %d", id)
			continue
		}
		booking, err := s.bookingRepo.GetById(occasion.Booking.BookingID)
		if err != nil {
			continue
		}
		// get the user email
		userEmail := occasion.GetUser().GetEmail()
		accomodationAddress := booking.GetAccomodation().Address
		isHouse := ""
		if booking.GetAccomodation().IsHouse{
			isHouse = "una casa"
		}else{
			isHouse = "un apartamento"
		}
		startDate := booking.GetEntryDate()
		endDate := booking.GetExitDate()

		// send the email
		err = s.emailService.SendEmail(
			userEmail,
			"Haz sido invitado a una reserva",
			fmt.Sprintf(`
				Estas cordialmente invitado a tu reserva en %q. La direccion es: %q.
				Recuerda que la fecha de entrada es  %s y la salida el %s, confirma tu asistencia por medio de nuesra aplicacion!
			`, isHouse, accomodationAddress, startDate, endDate),
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


func (s *RepoBookingsService) InviteAllBookings() (int, error){
	occasions, err := s.occasionRepo.GetAll(repo.OccasionForType(false))
	if err != nil{
		return 0, err
	}
	var occasionIds []int
	for _,occasion := range occasions{
		if occasion.State.GetStateName() == StateRegistered{
			occasionIds = append(occasionIds, occasion.OccasionID)
		}
	}
	return s.sendInvitation(occasionIds)
}

func (s *RepoBookingsService) InviteOccasion(occasionId int) (int, error){
	var occasionIds []int
	occasionIds = append(occasionIds, occasionId)
	return s.sendInvitation(occasionIds)
}
