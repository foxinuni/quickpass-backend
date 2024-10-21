package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

var (
	ErrOccasionNotInCourse  = errors.New("the occasion isn't in course right now, not able to do actions")
	ErrOccasionNotConfirmed = errors.New("the occasion isn't confirmed, cannot perform actions")
)

type ActionsService interface {
	NewAction(user *entities.User, occasionID int) (bool, error)
	GetLogs(eventId *int, bookingId *int) ([]*entities.LogHistory, error)
	GetLastLog(occasionId int) (*entities.LogHistory, *int, *int, error)
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


func (as *RepoActionsService) GetLogs(eventId *int, bookingId *int) ([]*entities.LogHistory, error){
	return as.logRepo.GetLogs(eventId, bookingId)
}

func (as *RepoActionsService) GetLastLog(occasionId int) (*entities.LogHistory, *int, *int, error){
	occasion, err := as.occasionRepo.GetById(occasionId)
	if err != nil{
		return nil,nil, nil, err
	}
	log, err := as.logRepo.GetLastLogFrom(occasionId)
	if err != nil{
		return nil,nil, nil, err
	}
	var bookingId *int = nil
	var eventId *int = nil
	if occasion.Booking != nil{
		bookingId = &occasion.Booking.BookingID
	}
	if occasion.Event != nil{
		bookingId = &occasion.Event.EventID
	}
	return log, eventId, bookingId, nil
}