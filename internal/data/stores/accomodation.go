package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type AccomodationFilter struct{}

type AccomodationStore interface {
	GetAll(ctx context.Context, filter AccomodationFilter) ([]models.Accomodation, error)
	GetById(ctx context.Context, id int) (*models.Accomodation, error)
	Create(ctx context.Context, accomodation *models.Accomodation) error
	Update(ctx context.Context, accomodation *models.Accomodation) error
	Delete(ctx context.Context, id int) error
}
