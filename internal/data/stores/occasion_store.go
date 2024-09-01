package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrOccasionNotFound = errors.New("occasion not found")

type OccasionFilter struct{}

type OccasionStore interface {
	GetAll(ctx context.Context, filter OccasionFilter) ([]models.Occasion, error)
	GetById(ctx context.Context, id int) (*models.Occasion, error)
	Create(ctx context.Context, occasion *models.Occasion) error
	Update(ctx context.Context, occasion *models.Occasion) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresOccasionStore implements OccasionStore
var _ OccasionStore = (*PostgresOccasionStore)(nil)

type PostgresOccasionStore struct {
	pool *pgxpool.Pool
}

func NewPostgresOccasionStore(pool *pgxpool.Pool) *PostgresOccasionStore {
	return &PostgresOccasionStore{
		pool: pool,
	}
}

func (s *PostgresOccasionStore) GetAll(ctx context.Context, filter OccasionFilter) ([]models.Occasion, error) {
	var occasions []models.Occasion
	rows, err := s.pool.Query(ctx, `SELECT occasion_id, user_id, event_id, booking_id, state_id FROM occasions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var occasion models.Occasion
		if err := rows.Scan(&occasion.OccasionID, &occasion.UserID, &occasion.EventID, &occasion.BookingID, &occasion.StateID); err != nil {
			return nil, err
		}

		occasions = append(occasions, occasion)
	}

	return occasions, nil
}

func (s *PostgresOccasionStore) GetById(ctx context.Context, id int) (*models.Occasion, error) {
	var occasion models.Occasion
	row := s.pool.QueryRow(ctx, `
		SELECT occasion_id, user_id, event_id, booking_id, state_id 
		FROM occasions
		WHERE occasion_id = $1
	`, id)
	if err := row.Scan(&occasion.OccasionID, &occasion.UserID, &occasion.EventID, &occasion.BookingID, &occasion.StateID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrOccasionNotFound
		}

		return nil, err
	}

	return &occasion, nil
}

func (s *PostgresOccasionStore) Create(ctx context.Context, occasion *models.Occasion) error {
	return s.pool.QueryRow(ctx, `
		INSERT INTO occasions (user_id, event_id, booking_id, state_id) 
		VALUES ($1, $2, $3, $4)
		RETURNING occasion_id
	`, occasion.UserID, occasion.EventID, occasion.BookingID, occasion.StateID).Scan(&occasion.OccasionID)
}

func (s *PostgresOccasionStore) Update(ctx context.Context, occasion *models.Occasion) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE occasions
		SET user_id = $1, event_id = $2, booking_id = $3, state_id = $4
		WHERE occasion_id = $5
	`, occasion.UserID, occasion.EventID, occasion.BookingID, occasion.StateID, occasion.OccasionID)
	return err
}

func (s *PostgresOccasionStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, "DELETE FROM occasions WHERE occasion_id = $1", id)
	return err
}
