package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type EventFilter struct{}

type EventStore interface {
	GetAll(ctx context.Context, filter EventFilter) ([]models.Event, error)
	GetById(ctx context.Context, id int) (*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
	Update(ctx context.Context, event *models.Event) error
	Delete(ctx context.Context, id int) error
}
