package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrStateNotFound = errors.New("state not found")

type StateFilter struct{}

type StateStore interface {
	GetAll(ctx context.Context, filter StateFilter) ([]models.State, error)
	GetById(ctx context.Context, id int) (*models.State, error)
	GetByName(ctx context.Context, name string) (*models.State, error)
	Create(ctx context.Context, state *models.State) error
	Update(ctx context.Context, state *models.State) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresStateStore implements StateStore
var _ StateStore = (*PostgresStateStore)(nil)

type PostgresStateStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStateStore(pool *pgxpool.Pool) *PostgresStateStore {
	return &PostgresStateStore{
		pool: pool,
	}
}

func (s *PostgresStateStore) GetAll(ctx context.Context, filter StateFilter) ([]models.State, error) {
	var states []models.State
	rows, err := s.pool.Query(ctx, `SELECT state_id, name FROM states`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var state models.State
		if err := rows.Scan(&state.StateID, &state.StateName); err != nil {
			return nil, err
		}

		states = append(states, state)
	}

	return states, nil
}

func (s *PostgresStateStore) GetById(ctx context.Context, id int) (*models.State, error) {
	var state models.State
	row := s.pool.QueryRow(ctx, `SELECT state_id, name FROM states WHERE state_id = $1`, id)
	if err := row.Scan(&state.StateID, &state.StateName); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStateNotFound
		}

		return nil, err
	}

	return &state, nil
}

func (s *PostgresStateStore) GetByName(ctx context.Context, name string) (*models.State, error) {
	var state models.State
	row := s.pool.QueryRow(ctx, `SELECT state_id, name FROM states WHERE name = $1`, name)
	if err := row.Scan(&state.StateID, &state.StateName); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrStateNotFound
		}

		return nil, err
	}

	return &state, nil
}

func (s *PostgresStateStore) Create(ctx context.Context, state *models.State) error {
	return s.pool.QueryRow(ctx, `INSERT INTO states (name) VALUES ($1) RETURNING state_id`, state.StateName).Scan(&state.StateID)
}

func (s *PostgresStateStore) Update(ctx context.Context, state *models.State) error {
	_, err := s.pool.Exec(ctx, `UPDATE states SET name = $1 WHERE state_id = $2`, state.StateName, state.StateID)
	return err
}

func (s *PostgresStateStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM states WHERE state_id = $1`, id)
	return err
}
