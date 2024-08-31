package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type OccasionFilter struct{}

type OccasionStore interface {
	GetAll(ctx context.Context, filter OccasionFilter) ([]models.Occasion, error)
	GetById(ctx context.Context, id int) (*models.Occasion, error)
	Create(ctx context.Context, occasion *models.Occasion) error
	Update(ctx context.Context, occasion *models.Occasion) error
	Delete(ctx context.Context, id int) error
}
