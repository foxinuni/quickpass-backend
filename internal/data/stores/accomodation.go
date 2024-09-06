package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAccomodationNotFound = errors.New("accomodation not found")

type AccomodationFilter struct{}

type AccomodationStore interface {
	GetAll(ctx context.Context, filter AccomodationFilter) ([]models.Accomodation, error)
	GetById(ctx context.Context, id int) (*models.Accomodation, error)
	GetByAddress(ctx context.Context, address string) (*models.Accomodation, error)
	Create(ctx context.Context, accomodation *models.Accomodation) error
	Update(ctx context.Context, accomodation *models.Accomodation) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresAccomodationStore implements AccomodationStore
var _ AccomodationStore = (*PostgresAccomodationStore)(nil)

type PostgresAccomodationStore struct {
	pool *pgxpool.Pool
}

func NewPostgresAccomodationStore(pool *pgxpool.Pool) *PostgresAccomodationStore {
	return &PostgresAccomodationStore{
		pool: pool,
	}
}

func (s *PostgresAccomodationStore) GetAll(ctx context.Context, filter AccomodationFilter) ([]models.Accomodation, error) {
	var accomodations []models.Accomodation
	rows, err := s.pool.Query(ctx, `SELECT acc_id, is_house, address FROM accomodations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accomodation models.Accomodation
		if err := rows.Scan(&accomodation.AccomodationID, &accomodation.IsHouse, &accomodation.Address); err != nil {
			return nil, err
		}

		accomodations = append(accomodations, accomodation)
	}

	return accomodations, nil

}

func (s *PostgresAccomodationStore) GetById(ctx context.Context, id int) (*models.Accomodation, error) {
	var accomodation models.Accomodation
	row := s.pool.QueryRow(ctx, `SELECT acc_id, is_house, address FROM accomodations WHERE acc_id = $1`, id)
	if err := row.Scan(&accomodation.AccomodationID, &accomodation.IsHouse, &accomodation.Address); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAccomodationNotFound
		}

		return nil, err
	}

	return &accomodation, nil
}

func (s *PostgresAccomodationStore) GetByAddress(ctx context.Context, address string) (*models.Accomodation, error) {
	var accomodation models.Accomodation
	row := s.pool.QueryRow(ctx, `SELECT acc_id, is_house, address FROM accomodations WHERE address = $1`, address)
	if err := row.Scan(&accomodation.AccomodationID, &accomodation.IsHouse, &accomodation.Address); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAccomodationNotFound
		}

		return nil, err
	}

	return &accomodation, nil
}

func (s *PostgresAccomodationStore) Create(ctx context.Context, accomodation *models.Accomodation) error {
	return s.pool.QueryRow(ctx, `
		INSERT INTO accomodations (is_house, address)
		VALUES ($1, $2)
		RETURNING acc_id
	`, accomodation.IsHouse, accomodation.Address).Scan(&accomodation.AccomodationID)
}

func (s *PostgresAccomodationStore) Update(ctx context.Context, accomodation *models.Accomodation) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE accomodations
		SET is_house = $1, address = $2
		WHERE acc_id = $3
		`, accomodation.IsHouse, accomodation.Address, accomodation.AccomodationID)
	return err
}

func (s *PostgresAccomodationStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM accomodations WHERE acc_id = $1`, id)
	return err
}
