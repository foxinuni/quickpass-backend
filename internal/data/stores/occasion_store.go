package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrOccasionNotFound = errors.New("occasion not found")

type OccasionFilter struct {
	UserID    *int
	EventID   *int
	BookingID *int
	StateID   *int
	TypeOccasion *bool	//true if its event, false if its booking
}

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

	rows, err := s.pool.Query(ctx, `
		SELECT occasion_id, user_id, event_id, booking_id, state_id 
		FROM occasions
		WHERE
				(CASE WHEN $1::int IS NULL THEN TRUE ELSE user_id = $1::int END)
			AND (CASE WHEN $2::int IS NULL THEN TRUE ELSE event_id = $2::int END)
			AND (CASE WHEN $3::int IS NULL THEN TRUE ELSE booking_id = $3::int END)
			AND (CASE WHEN $4::int IS NULL THEN TRUE ELSE state_id = $4::int END)
			AND (CASE WHEN $5::bool IS NULL THEN TRUE WHEN $5::bool = TRUE THEN event_id IS NOT NULL ELSE booking_id IS NOT NULL END)
	`, filter.UserID, filter.EventID, filter.BookingID, filter.StateID, filter.TypeOccasion)
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
