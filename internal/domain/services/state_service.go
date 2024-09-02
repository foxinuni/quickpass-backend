package services

import (
	"github.com/foxinuni/quickpass-backend/internal/data/repo"
	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type StateService interface {
	GetAllStates() ([]*entities.State, error)
	GetStateByID(id int) (*entities.State, error)
	GetOrCreateState(name string) (*entities.State, error)
}

type StoreStateService struct {
	stateRepo repo.StateRepository
}

func NewStoreStateService(stateRepo repo.StateRepository) StateService {
	return &StoreStateService{
		stateRepo: stateRepo,
	}
}

func (s *StoreStateService) GetAllStates() ([]*entities.State, error) {
	return s.stateRepo.GetAll()
}

func (s *StoreStateService) GetStateByID(id int) (*entities.State, error) {
	return s.stateRepo.GetById(id)
}

func (s *StoreStateService) GetOrCreateState(name string) (*entities.State, error) {
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
