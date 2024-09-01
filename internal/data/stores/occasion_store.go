package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OccasionFilter struct{}

type OccasionStore interface {
	GetAll(ctx context.Context, filter OccasionFilter) ([]models.Occasion, error)
	GetById(ctx context.Context, id int) (*models.Occasion, error)
	Create(ctx context.Context, occasion *models.Occasion) error
	Update(ctx context.Context, occasion *models.Occasion) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresOccasionStore implements OccasionStore
var _ OccasionStore = (*PostgresOccasionStore)(nil)

type PostgresOccasionStore struct {
	pool *pgxpool.Pool
}

func NewPostgresOccasionStore(pool *pgxpool.Pool) *PostgresOccasionStore {
	return &PostgresOccasionStore{
		pool: pool,
	}
}

func (s *PostgresOccasionStore) GetAll(ctx context.Context, filter OccasionFilter) ([]models.Occasion, error) {
	panic("not implemented")
}

func (s *PostgresOccasionStore) GetById(ctx context.Context, id int) (*models.Occasion, error) {
	panic("not implemented")
}

func (s *PostgresOccasionStore) Create(ctx context.Context, occasion *models.Occasion) error {
	panic("not implemented")
}

func (s *PostgresOccasionStore) Update(ctx context.Context, occasion *models.Occasion) error {
	panic("not implemented")
}

func (s *PostgresOccasionStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
