package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type BookingFilter struct{}

type BookingStore interface {
	GetAll(ctx context.Context, filter BookingFilter) ([]models.Booking, error)
	GetById(ctx context.Context, id int) (*models.Booking, error)
	Create(ctx context.Context, booking *models.Booking) error
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, id int) error
}
