package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrEventNotFound = errors.New("event not found")

type EventFilter struct{}

type EventStore interface {
	GetAll(ctx context.Context, filter EventFilter) ([]models.Event, error)
	GetById(ctx context.Context, id int) (*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
	Update(ctx context.Context, event *models.Event) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresEventStore implements EventStore
var _ EventStore = (*PostgresEventStore)(nil)

type PostgresEventStore struct {
	pool *pgxpool.Pool
}

func NewPostgresEventStore(pool *pgxpool.Pool) *PostgresEventStore {
	return &PostgresEventStore{
		pool: pool,
	}
}

func (s *PostgresEventStore) GetAll(ctx context.Context, filter EventFilter) ([]models.Event, error) {
	var events []models.Event
	rows, err := s.pool.Query(ctx, `SELECT event_id, start_date, end_date, address, name FROM events`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		if err := rows.Scan(&event.EventID, &event.StartDate, &event.EndDate, &event.Address, &event.Name); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (s *PostgresEventStore) GetById(ctx context.Context, id int) (*models.Event, error) {
	var event models.Event
	row := s.pool.QueryRow(ctx, `SELECT event_id, start_date, end_date, address, name FROM events WHERE event_id = $1`, id)
	if err := row.Scan(&event.EventID, &event.StartDate, &event.EndDate, &event.Address, &event.Name); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEventNotFound
		}

		return nil, err
	}

	return &event, nil
}

func (s *PostgresEventStore) Create(ctx context.Context, event *models.Event) error {
	return s.pool.QueryRow(ctx, `
		INSERT INTO events (start_date, end_date, address, name)
		VALUES ($1, $2, $3, $4)
		RETURNING event_id
	`, event.StartDate, event.EndDate, event.Address, event.Name).Scan(&event.EventID)
}

func (s *PostgresEventStore) Update(ctx context.Context, event *models.Event) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE events
		SET start_date = $1, end_date = $2, address = $3, name = $4
		WHERE event_id = $5
	`, event.StartDate, event.EndDate, event.Address, event.Name, event.EventID)
	return err
}

func (s *PostgresEventStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM events WHERE event_id = $1`, id)
	return err
}
