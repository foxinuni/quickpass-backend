package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type StateRepository interface {
	GetAll() ([]*entities.State, error)
	GetById(id int) (*entities.State, error)
	GetByName(name string) (*entities.State, error)
	Create(state *entities.State) error
	Update(state *entities.State) error
	Delete(stateId int) error
}

type StoreStateRepository struct {
	stateStore stores.StateStore
}

func NewStoreStateRepository(stateStore stores.StateStore) StateRepository {
	return &StoreStateRepository{
		stateStore: stateStore,
	}
}

func (r *StoreStateRepository) GetAll() ([]*entities.State, error) {
	states, err := r.stateStore.GetAll(context.Background(), stores.StateFilter{})
	if err != nil {
		return nil, err
	}

	var result []*entities.State
	for _, state := range states {
		result = append(result, ModelToState(&state))
	}

	return result, nil
}

func (r *StoreStateRepository) GetById(id int) (*entities.State, error) {
	state, err := r.stateStore.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return ModelToState(state), nil
}

func (r *StoreStateRepository) GetByName(name string) (*entities.State, error) {
	state, err := r.stateStore.GetByName(context.Background(), name)
	if err != nil {
		return nil, err
	}

	return ModelToState(state), nil
}

func (r *StoreStateRepository) Create(state *entities.State) error {
	model := StateToModel(state)

	if err := r.stateStore.Create(context.Background(), model); err != nil {
		return err
	}

	*state = *ModelToState(model)
	return nil
}

func (r *StoreStateRepository) Update(state *entities.State) error {
	model := StateToModel(state)

	if err := r.stateStore.Update(context.Background(), model); err != nil {
		return err
	}

	*state = *ModelToState(model)
	return nil
}

func (r *StoreStateRepository) Delete(stateId int) error {
	return r.stateStore.Delete(context.Background(), stateId)
}
