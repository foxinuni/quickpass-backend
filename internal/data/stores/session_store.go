package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionFilter struct{}

type SessionStore interface {
	GetAll(ctx context.Context, filter SessionFilter) ([]models.Session, error)
	GetById(ctx context.Context, id int) (*models.Session, error)
	GetByToken(ctx context.Context, token string) (*models.Session, error)
	Create(ctx context.Context, session *models.Session) error
	Update(ctx context.Context, session *models.Session) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresSessionStore implements SessionStore
var _ SessionStore = (*PostgresSessionStore)(nil)

type PostgresSessionStore struct {
	pool *pgxpool.Pool
}

func NewPostgresSessionStore(pool *pgxpool.Pool) *PostgresSessionStore {
	return &PostgresSessionStore{
		pool: pool,
	}
}

func (s *PostgresSessionStore) GetAll(ctx context.Context, filter SessionFilter) ([]models.Session, error) {
	panic("not implemented")
}

func (s *PostgresSessionStore) GetById(ctx context.Context, id int) (*models.Session, error) {
	panic("not implemented")
}

func (s *PostgresSessionStore) GetByToken(ctx context.Context, token string) (*models.Session, error) {
	panic("not implemented")
}

func (s *PostgresSessionStore) Create(ctx context.Context, session *models.Session) error {
	panic("not implemented")
}

func (s *PostgresSessionStore) Update(ctx context.Context, session *models.Session) error {
	panic("not implemented")
}

func (s *PostgresSessionStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
