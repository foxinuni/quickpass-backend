package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type LogFilter struct{}

type LogStore interface {
	GetAll(ctx context.Context, filter LogFilter) ([]models.Log, error)
	GetById(ctx context.Context, id int) (*models.Log, error)
	Create(ctx context.Context, log *models.Log) error
	Update(ctx context.Context, log *models.Log) error
	Delete(ctx context.Context, id int) error
}
