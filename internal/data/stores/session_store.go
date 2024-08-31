package stores

import (
	"context"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
)

type SessionFilter struct{}

type SessionStore interface {
	GetAll(ctx context.Context, filter SessionFilter) ([]models.Session, error)
	GetById(ctx context.Context, id int) (*models.Session, error)
	Create(ctx context.Context, session *models.Session) error
	Update(ctx context.Context, session *models.Session) error
	Delete(ctx context.Context, id int) error
}
