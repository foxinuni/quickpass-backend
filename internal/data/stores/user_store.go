package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type UserFilters struct{}

type UserStore interface {
	GetAll(ctx context.Context, filter UserFilters) ([]models.User, error)
	GetById(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int) error
}
