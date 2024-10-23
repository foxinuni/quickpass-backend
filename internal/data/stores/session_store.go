package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrSessionNotFound = errors.New("session not found")

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
	var sessions []models.Session
	rows, err := s.pool.Query(ctx, `SELECT session_id, user_id, enabled, token, phone_model FROM sessions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var session models.Session
		if err := rows.Scan(&session.SessionID, &session.UserID, &session.Enabled, &session.Token, &session.PhoneModel); err != nil {
			return nil, err
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (s *PostgresSessionStore) GetById(ctx context.Context, id int) (*models.Session, error) {
	var session models.Session
	row := s.pool.QueryRow(ctx, `SELECT session_id, user_id, enabled, token, phone_model FROM sessions WHERE session_id = $1`, id)

	// Scan the row into the session model
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Enabled, &session.Token, &session.PhoneModel); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
	}

	// Return the session
	return &session, nil
}

func (s *PostgresSessionStore) GetByToken(ctx context.Context, token string) (*models.Session, error) {
	var session models.Session
	row := s.pool.QueryRow(ctx, `SELECT session_id, user_id, enabled, token, phone_model FROM sessions WHERE token = $1`, token)

	// Scan the row into the session model
	if err := row.Scan(&session.SessionID, &session.UserID, &session.Enabled, &session.Token, &session.PhoneModel); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrSessionNotFound
		}

		return nil, err
	}

	// Return the session
	return &session, nil
}

func (s *PostgresSessionStore) Create(ctx context.Context, session *models.Session) error {
	return s.pool.QueryRow(ctx, `
		INSERT INTO sessions (user_id, enabled, token, phone_model) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING session_id
	`, session.UserID, session.Enabled, session.Token, session.PhoneModel).Scan(&session.SessionID)
}

func (s *PostgresSessionStore) Update(ctx context.Context, session *models.Session) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE sessions 
		SET user_id = $1, enabled = $2, token = $3, phone_model = $4
		WHERE session_id = $5
	`, session.UserID, session.Enabled, session.Token, session.PhoneModel, session.SessionID)
	return err
}

func (s *PostgresSessionStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM sessions WHERE session_id = $1`, id)
	return err
}
