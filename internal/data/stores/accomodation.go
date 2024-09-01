package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccomodationFilter struct{}

type AccomodationStore interface {
	GetAll(ctx context.Context, filter AccomodationFilter) ([]models.Accomodation, error)
	GetById(ctx context.Context, id int) (*models.Accomodation, error)
	Create(ctx context.Context, accomodation *models.Accomodation) error
	Update(ctx context.Context, accomodation *models.Accomodation) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresAccomodationStore implements AccomodationStore
var _ AccomodationStore = (*PostgresAccomodationStore)(nil)

type PostgresAccomodationStore struct {
	pool *pgxpool.Pool
}

func NewPostgresAccomodationStore(pool *pgxpool.Pool) *PostgresAccomodationStore {
	return &PostgresAccomodationStore{
		pool: pool,
	}
}

func (s *PostgresAccomodationStore) GetAll(ctx context.Context, filter AccomodationFilter) ([]models.Accomodation, error) {
	panic("not implemented")
}

func (s *PostgresAccomodationStore) GetById(ctx context.Context, id int) (*models.Accomodation, error) {
	panic("not implemented")
}

func (s *PostgresAccomodationStore) Create(ctx context.Context, accomodation *models.Accomodation) error {
	panic("not implemented")
}

func (s *PostgresAccomodationStore) Update(ctx context.Context, accomodation *models.Accomodation) error {
	panic("not implemented")
}

func (s *PostgresAccomodationStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
