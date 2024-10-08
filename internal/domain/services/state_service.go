package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

var (
	StateRegistered = "registered"
	StateInvited    = "invited"
	StateConfirmed  = "confirmed"
	StateDeclined   = "declined"
)

type StateService interface {
	GetAllStates() ([]*entities.State, error)
	GetStateByID(id int) (*entities.State, error)
	GetOrCreateState(name string) (*entities.State, error)
}

type RepoStateService struct {
	stateRepo repo.StateRepository
}

func NewRepoStateService(stateRepo repo.StateRepository) StateService {
	return &RepoStateService{
		stateRepo: stateRepo,
	}
}

func (s *RepoStateService) GetAllStates() ([]*entities.State, error) {
	return s.stateRepo.GetAll()
}

func (s *RepoStateService) GetStateByID(id int) (*entities.State, error) {
	return s.stateRepo.GetById(id)
}

func (s *RepoStateService) GetOrCreateState(name string) (*entities.State, error) {
	state, err := s.stateRepo.GetByName(name)
	if err != nil {
		if err == stores.ErrStateNotFound {
			// Create the state if it does not exist
			entities := entities.NewState(0, name)
			if err := s.stateRepo.Create(entities); err != nil {
				return nil, err
			}

			return entities, nil
		}

		return nil, err
	}

	return state, nil
}
