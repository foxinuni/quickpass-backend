package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

var (
	ErrOccasionNotInCourse  = errors.New("The occasion isn't in course right now, not able to do actions")
	ErrOccasionNotConfirmed = errors.New("The occasion isn't confirmed, cannot perform actions")
)

type ActionsService interface {
	NewAction(user *entities.User, occasionID int) (bool, error)
}

type RepoActionsService struct {
	logRepo      repo.LogRepository
	occasionRepo repo.OccasionRepository
}

func NewRepoActionsService(logRepo repo.LogRepository, occasionRepo repo.OccasionRepository) ActionsService {
	return &RepoActionsService{
		logRepo:      logRepo,
		occasionRepo: occasionRepo,
	}
}

func (as *RepoActionsService) NewAction(user *entities.User, occasionID int) (bool, error) {
	occasion, err := as.occasionRepo.GetById(occasionID)
	if err != nil {
		return false, err
	}
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
	if !(afterStart && beforeEnd) {
		return false, ErrOccasionNotInCourse
	}
	if occasion.GetState().GetStateName() != "confirmed" {
		fmt.Print(occasion.GetState().GetStateName())
		return false, ErrOccasionNotConfirmed
	}

	return as.logRepo.NewAction(occasionID)
}
