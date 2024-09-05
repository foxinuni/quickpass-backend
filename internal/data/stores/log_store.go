package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrLogNotFound = errors.New("log not found")

type LogFilter struct{}

type LogStore interface {
	GetAll(ctx context.Context, filter LogFilter) ([]models.Log, error)
	GetById(ctx context.Context, id int) (*models.Log, error)
	Create(ctx context.Context, log *models.Log) error
	Update(ctx context.Context, log *models.Log) error
	Delete(ctx context.Context, id int) error
	GetLastFromOcassion(ctx context.Context, id int) (*models.Log, error)
}

// Checks at compile time if PostgresLogStore implements LogStore
var _ LogStore = (*PostgresLogStore)(nil)

type PostgresLogStore struct {
	pool *pgxpool.Pool
}

func NewPostgresLogStore(pool *pgxpool.Pool) *PostgresLogStore {
	return &PostgresLogStore{
		pool: pool,
	}
}

func (s *PostgresLogStore) GetAll(ctx context.Context, filter LogFilter) ([]models.Log, error) {
	var logs []models.Log
	rows, err := s.pool.Query(ctx, `SELECT log_id, occasion_id, time, is_inside FROM logs`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log models.Log
		if err := rows.Scan(&log.LogID, &log.OccasionID, &log.Time, &log.IsInside); err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func (s *PostgresLogStore) GetById(ctx context.Context, id int) (*models.Log, error) {
	var log models.Log
	row := s.pool.QueryRow(ctx, `
		SELECT log_id, occasion_id, time, is_inside
		FROM logs
		WHERE log_id = $1
	`, id)
	if err := row.Scan(&log.LogID, &log.OccasionID, &log.Time, &log.IsInside); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrLogNotFound
		}

		return nil, err
	}

	return &log, nil
}

// added in order to get the last log from a certain occasion, used to know if the user
// is currently inside or outside
func (s *PostgresLogStore) GetLastFromOcassion(ctx context.Context, id int) (*models.Log, error) {
	var log models.Log
	row := s.pool.QueryRow(ctx, `
		SELECT log_id, occasion_id, time, is_inside
		FROM logs
		WHERE occasion_id = $1
		ORDER BY time DESC
		LIMIT 1
	`, id)
	if err := row.Scan(&log.LogID, &log.OccasionID, &log.Time, &log.IsInside); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			//since here it isn't an error if the log wasn't found, means the occasion hasn't any logs associated at the moment
			return nil, nil
		}

		return nil, err
	}

	return &log, nil
}

func (s *PostgresLogStore) Create(ctx context.Context, log *models.Log) error {
	return s.pool.QueryRow(ctx, `INSERT INTO logs (occasion_id, time, is_inside) VALUES ($1, $2, $3) RETURNING log_id`, log.OccasionID, log.Time, log.IsInside).Scan(&log.LogID)
}

func (s *PostgresLogStore) Update(ctx context.Context, log *models.Log) error {
	_, err := s.pool.Exec(ctx, `UPDATE logs SET occasion_id = $1, time = $2, is_inside = $3 WHERE log_id = $4`, log.OccasionID, log.Time, log.IsInside, log.LogID)
	return err
}

func (s *PostgresLogStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM logs WHERE log_id = $1`, id)
	return err
}
