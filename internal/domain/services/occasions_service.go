package services

import (
	"errors"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

var ErrNotVisibleForOccasion = errors.New("you are forbiden to see this occasion")

type OccassionService interface {
	GetOccasionsForUser(user *entities.User, active bool) ([]*entities.Occasion, error)
	GetOccasionForUsersWithId(user *entities.User, occasionID int) (*entities.Occasion, error)
	ConfirmOccasionForUser(user *entities.User, occasionID int, confirm bool) (*entities.Occasion, error)
}

type RepoOccassionsService struct {
	occasionRepo repo.OccasionRepository
	stateService StateService
}

func NewRepoOccassionsService(occasionRepo repo.OccasionRepository, stateService StateService) OccassionService {
	return &RepoOccassionsService{
		occasionRepo: occasionRepo,
		stateService: stateService,
	}
}

func (s *RepoOccassionsService) GetOccasionsForUser(user *entities.User, active bool) ([]*entities.Occasion, error) {
	// Get the occasions from the repository
	occasions, err := s.occasionRepo.GetAll(repo.OccasionForUser(user))
	if err != nil {
		return nil, err
	}

	// Return the unfiltered occasions
	if !active {
		return occasions, nil
	}

	var filteredOccasions []*entities.Occasion
	for _, occasion := range occasions {
		var afterStart bool
		var beforeEnd bool
		if occasion.Event != nil {
			afterStart = time.Now().After(occasion.Event.GetStartDate())
			beforeEnd = occasion.Event.GetStartDate().Before(time.Now())
		}
		if occasion.Booking != nil {
			afterStart = time.Now().After(occasion.Booking.GetEntryDate())
			beforeEnd = occasion.Booking.GetExitDate().Before(time.Now())
		}
		if afterStart && beforeEnd {
			filteredOccasions = append(filteredOccasions, occasion)
		}
	}

	return filteredOccasions, nil
}

func (s *RepoOccassionsService) GetOccasionForUsersWithId(user *entities.User, occasionID int) (*entities.Occasion, error) {
	// Get the occasion from the repository
	occasion, err := s.occasionRepo.GetById(occasionID)
	if err != nil {
		return nil, err
	}

	// Check if the occasion belongs to the user
	if occasion.GetUser().GetUserID() != user.GetUserID() {
		return nil, ErrNotVisibleForOccasion
	}

	return occasion, nil
}

func (s *RepoOccassionsService) ConfirmOccasionForUser(user *entities.User, occasionID int, confirm bool) (*entities.Occasion, error) {
	// Get the occasion from the repository
	occasion, err := s.occasionRepo.GetById(occasionID)
	if err != nil {
		return nil, err
	}

	// Check if the occasion belongs to the user
	if occasion.GetUser().GetUserID() != user.GetUserID() {
		return nil, ErrNotVisibleForOccasion
	}

	// Get the state for the occasion
	var state *entities.State
	if confirm {
		created, err := s.stateService.GetOrCreateState("confirmed")
		if err != nil {
			return nil, err
		}

		state = created
	} else {
		created, err := s.stateService.GetOrCreateState("declined")
		if err != nil {
			return nil, err
		}

		state = created
	}

	// Confirm the occasion
	occasion.SetState(state)

	// Update the occasion in the repository
	if err := s.occasionRepo.Update(occasion); err != nil {
		return nil, err
	}

	return occasion, nil
}
