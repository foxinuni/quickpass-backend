package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EventFilter struct{}

type EventStore interface {
	GetAll(ctx context.Context, filter EventFilter) ([]models.Event, error)
	GetById(ctx context.Context, id int) (*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
	Update(ctx context.Context, event *models.Event) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresEventStore implements EventStore
var _ EventStore = (*PostgresEventStore)(nil)

type PostgresEventStore struct {
	pool *pgxpool.Pool
}

func NewPostgresEventStore(pool *pgxpool.Pool) *PostgresEventStore {
	return &PostgresEventStore{
		pool: pool,
	}
}

func (s *PostgresEventStore) GetAll(ctx context.Context, filter EventFilter) ([]models.Event, error) {
	panic("not implemented")
}

func (s *PostgresEventStore) GetById(ctx context.Context, id int) (*models.Event, error) {
	panic("not implemented")
}

func (s *PostgresEventStore) Create(ctx context.Context, event *models.Event) error {
	panic("not implemented")
}

func (s *PostgresEventStore) Update(ctx context.Context, event *models.Event) error {
	panic("not implemented")
}

func (s *PostgresEventStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
