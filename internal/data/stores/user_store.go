package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserFilters struct{}

type UserStore interface {
	GetAll(ctx context.Context, filter UserFilters) ([]models.User, error)
	GetById(ctx context.Context, id int) (*models.User, error)
	GetByPhone(ctx context.Context, number string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
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
	var user models.User
	row := s.pool.QueryRow(ctx, `SELECT user_id, email, number FROM users WHERE user_id = $1`, id)

	if err := row.Scan(&user.UserID, &user.Email, &user.Number); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStore) GetByPhone(ctx context.Context, number string) (*models.User, error) {
	var user models.User
	row := s.pool.QueryRow(ctx, `SELECT user_id, email, number FROM users WHERE number = $1`, number)

	if err := row.Scan(&user.UserID, &user.Email, &user.Number); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}


func (s *PostgresUserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	row := s.pool.QueryRow(ctx, `SELECT user_id, email, number FROM users WHERE email = $1`, email)

	if err := row.Scan(&user.UserID, &user.Email, &user.Number); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStore) Create(ctx context.Context, user *models.User) error {
	return s.pool.QueryRow(ctx, `INSERT INTO users (email, number) VALUES ($1, $2) RETURNING user_id`, user.Email, user.Number).Scan(&user.UserID)
}

func (s *PostgresUserStore) Update(ctx context.Context, user *models.User) error {
	_, err := s.pool.Exec(ctx, `UPDATE users SET email = $1, number = $2 WHERE user_id = $3`, user.Email, user.Number, user.UserID)
	return err
}

func (s *PostgresUserStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, "DELETE FROM users WHERE user_id = $1", id)
	return err
}
