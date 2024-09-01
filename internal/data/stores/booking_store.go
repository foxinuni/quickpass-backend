package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookingFilter struct{}

type BookingStore interface {
	GetAll(ctx context.Context, filter BookingFilter) ([]models.Booking, error)
	GetById(ctx context.Context, id int) (*models.Booking, error)
	Create(ctx context.Context, booking *models.Booking) error
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresBookingStore implements BookingStore
var _ BookingStore = (*PostgresBookingStore)(nil)

type PostgresBookingStore struct {
	pool *pgxpool.Pool
}

func NewPostgresBookingStore(pool *pgxpool.Pool) *PostgresBookingStore {
	return &PostgresBookingStore{
		pool: pool,
	}
}

func (s *PostgresBookingStore) GetAll(ctx context.Context, filter BookingFilter) ([]models.Booking, error) {
	panic("not implemented")
}

func (s *PostgresBookingStore) GetById(ctx context.Context, id int) (*models.Booking, error) {
	panic("not implemented")
}

func (s *PostgresBookingStore) Create(ctx context.Context, booking *models.Booking) error {
	panic("not implemented")
}

func (s *PostgresBookingStore) Update(ctx context.Context, booking *models.Booking) error {
	panic("not implemented")
}

func (s *PostgresBookingStore) Delete(ctx context.Context, id int) error {
	panic("not implemented")
}
