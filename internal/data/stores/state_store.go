package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type StateFilter struct{}

type StateStore interface {
	GetAll(ctx context.Context, filter StateFilter) ([]models.State, error)
	GetById(ctx context.Context, id int) (*models.State, error)
	Create(ctx context.Context, state *models.State) error
	Update(ctx context.Context, state *models.State) error
	Delete(ctx context.Context, id int) error
}
