package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StateFilter struct{}

type StateStore interface {
	GetAll(ctx context.Context, filter StateFilter) ([]models.State, error)
	GetById(ctx context.Context, id int) (*models.State, error)
	Create(ctx context.Context, state *models.State) error
	Update(ctx context.Context, state *models.State) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresStateStore implements StateStore
var _ StateStore = (*PostgresStateStore)(nil)

type PostgresStateStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStateStore(pool *pgxpool.Pool) *PostgresStateStore {
	return &PostgresStateStore{
		pool: pool,
	}
}

func (s *PostgresStateStore) GetAll(ctx context.Context, filter StateFilter) ([]models.State, error) {
	panic("not implemented")
}

func (s *PostgresStateStore) GetById(ctx context.Context, id int) (*models.State, error) {
	panic("not implemented")
}

func (s *PostgresStateStore) Create(ctx context.Context, state *models.State) error {
	panic("not implemented")
}

func (s *PostgresStateStore) Update(ctx context.Context, state *models.State) error {
	panic("not implemented")
}

func (s *PostgresStateStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
