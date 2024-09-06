package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type EventRepository interface {
	GetAll() ([]*entities.Event, error)
	GetById(id int) (*entities.Event, error)
	GetByName(name string) (*entities.Event, error)
	Create(event *entities.Event) error
}

type StoreEventRepository struct {
	occasionStore stores.OccasionStore
	eventStore    stores.EventStore
}

func NewStoreEventRepository(
	occasionStore stores.OccasionStore,
	eventStore stores.EventStore,
) EventRepository {
	return &StoreEventRepository{
		occasionStore: occasionStore,
		eventStore:    eventStore,
	}
}

func (r *StoreEventRepository) GetAll() ([]*entities.Event, error) {
	events, err := r.eventStore.GetAll(context.Background())
	if err != nil {
		return nil, err
	}

	var result []*entities.Event
	for _, event := range events {
		eventEntity := ModelToEvent(&event)
		result = append(result, eventEntity)
	}
	return result, nil
}

func (r *StoreEventRepository) GetById(id int) (*entities.Event, error) {
	event, err := r.eventStore.GetById(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return ModelToEvent(event), nil
}

func (r *StoreEventRepository) GetByName(name string) (*entities.Event, error) {
	event, err := r.eventStore.GetByName(context.Background(), name)
	if err != nil {
		return nil, err
	}

	return ModelToEvent(event), nil
}

func (r *StoreEventRepository) Create(event *entities.Event) error {
	model := EventToModel(event)
	if err := r.eventStore.Create(context.Background(), model); err != nil {
		return err
	}

	*event = *ModelToEvent(model)
	return nil
}
