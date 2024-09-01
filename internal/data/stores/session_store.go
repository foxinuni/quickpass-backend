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
	var session models.Session
	row := s.pool.QueryRow(ctx, `SELECT id, user_id, enabled, token, phone_model, imei FROM sessions WHERE id = $1`, id)

	// Scan the row into the session model
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Enabled, &session.Token, &session.PhoneModel, &session.IMEI); err != nil {
		return nil, err
	}

	// Return the session
	return &session, nil
}

func (s *PostgresSessionStore) GetByToken(ctx context.Context, token string) (*models.Session, error) {
	var session models.Session
	row := s.pool.QueryRow(ctx, `SELECT id, user_id, enabled, token, phone_model, imei FROM sessions WHERE token = $1`, token)

	// Scan the row into the session model
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Enabled, &session.Token, &session.PhoneModel, &session.IMEI); err != nil {
		return nil, err
	}

	// Return the session
	return &session, nil
}

func (s *PostgresSessionStore) Create(ctx context.Context, session *models.Session) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO sessions (user_id, enabled, token, phone_model, imei) 
		VALUES ($1, $2, $3, $4, $5)
	`, session.UserID, session.Enabled, session.Token, session.PhoneModel, session.IMEI)

	return err
}

func (s *PostgresSessionStore) Update(ctx context.Context, session *models.Session) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE sessions 
		SET user_id = $1, enabled = $2, token = $3, phone_model = $4, imei = $5
		WHERE id = $6
	`, session.UserID, session.Enabled, session.Token, session.PhoneModel, session.IMEI, session.SessionID)

	return err
}

func (s *PostgresSessionStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM sessions WHERE id = $1`, id)
	return err
}
