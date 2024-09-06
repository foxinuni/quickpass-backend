package repo

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
)

type AccomodationRepository interface {
	GetByAddress(address string) (*entities.Accomodation, error)
	Create(accomodation *entities.Accomodation) error
}

type StoreAccomodationRepository struct {
	accomodationStore stores.AccomodationStore
}

func NewStoreAccomodationRepository(accomodationStore stores.AccomodationStore) AccomodationRepository {
	return &StoreAccomodationRepository{
		accomodationStore: accomodationStore,
	}
}

func (r *StoreAccomodationRepository) GetByAddress(address string) (*entities.Accomodation, error) {
	accomodation, err := r.accomodationStore.GetByAddress(context.Background(), address)
	if err != nil {
		return nil, err
	}

	return ModelToAccomodation(accomodation), nil
}

func (r *StoreAccomodationRepository) Create(accomodation *entities.Accomodation) error {
	model := AccomodationToModel(accomodation)
	if err := r.accomodationStore.Create(context.Background(), model); err != nil {
		return err
	}

	*accomodation = *ModelToAccomodation(model)
	return nil
}
