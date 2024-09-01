package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LogFilter struct{}

type LogStore interface {
	GetAll(ctx context.Context, filter LogFilter) ([]models.Log, error)
	GetById(ctx context.Context, id int) (*models.Log, error)
	Create(ctx context.Context, log *models.Log) error
	Update(ctx context.Context, log *models.Log) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresLogStore implements LogStore
var _ LogStore = (*PostgresLogStore)(nil)

type PostgresLogStore struct {
	pool *pgxpool.Pool
}

func NewPostgresLogStore(pool *pgxpool.Pool) *PostgresLogStore {
	return &PostgresLogStore{
		pool: pool,
	}
}

func (s *PostgresLogStore) GetAll(ctx context.Context, filter LogFilter) ([]models.Log, error) {
	panic("not implemented")
}

func (s *PostgresLogStore) GetById(ctx context.Context, id int) (*models.Log, error) {
	panic("not implemented")
}

func (s *PostgresLogStore) Create(ctx context.Context, log *models.Log) error {
	panic("not implemented")
}

func (s *PostgresLogStore) Update(ctx context.Context, log *models.Log) error {
	panic("not implemented")
}

func (s *PostgresLogStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
