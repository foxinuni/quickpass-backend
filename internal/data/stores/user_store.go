package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserFilters struct{}

type UserStore interface {
	GetAll(ctx context.Context, filter UserFilters) ([]models.User, error)
	GetById(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresUserStore implements UserStore
var _ UserStore = (*PostgresUserStore)(nil)

type PostgresUserStore struct {
	pool *pgxpool.Pool
}

func NewPostgresUserStore(pool *pgxpool.Pool) *PostgresUserStore {
	return &PostgresUserStore{
		pool: pool,
	}
}

func (s *PostgresUserStore) GetAll(ctx context.Context, filter UserFilters) ([]models.User, error) {
	panic("not implemented")
}

func (s *PostgresUserStore) GetById(ctx context.Context, id int) (*models.User, error) {
	panic("not implemented")
}

func (s *PostgresUserStore) Create(ctx context.Context, user *models.User) error {
	panic("not implemented")
}

func (s *PostgresUserStore) Update(ctx context.Context, user *models.User) error {
	panic("not implemented")
}

func (s *PostgresUserStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
