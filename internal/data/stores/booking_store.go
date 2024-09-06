package stores

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrBookingNotFound = errors.New("booking not found")

type BookingFilter struct {
	UserID *int
}

type BookingStore interface {
	GetAll(ctx context.Context, filter BookingFilter) ([]models.Booking, error)
	GetById(ctx context.Context, id int) (*models.Booking, error)
	Create(ctx context.Context, booking *models.Booking) error
	Update(ctx context.Context, booking *models.Booking) error
	Delete(ctx context.Context, id int) error
}

// Checks at compile time if PostgresBookingStore implements BookingStore
var _ BookingStore = (*PostgresBookingStore)(nil)

type PostgresBookingStore struct {
	pool *pgxpool.Pool
}

func NewPostgresBookingStore(pool *pgxpool.Pool) *PostgresBookingStore {
	return &PostgresBookingStore{
		pool: pool,
	}
}

func (s *PostgresBookingStore) GetAll(ctx context.Context, filter BookingFilter) ([]models.Booking, error) {
	var bookings []models.Booking
	rows, err := s.pool.Query(ctx, `SELECT booking_id, entry_date, leaving_date, acc_id FROM bookings`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(&booking.BookingID, &booking.EntryDate, &booking.ExitDate, &booking.AccomodationID); err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (s *PostgresBookingStore) GetById(ctx context.Context, id int) (*models.Booking, error) {
	var booking models.Booking
	row := s.pool.QueryRow(ctx, `SELECT booking_id, entry_date, leaving_date, acc_id FROM bookings WHERE booking_id = $1`, id)
	if err := row.Scan(&booking.BookingID, &booking.EntryDate, &booking.ExitDate, &booking.AccomodationID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrBookingNotFound
		}

		return nil, err
	}

	return &booking, nil
}

func (s *PostgresBookingStore) Create(ctx context.Context, booking *models.Booking) error {
	return s.pool.QueryRow(ctx, `
		INSERT INTO bookings (entry_date, leaving_date, acc_id)
		VALUES ($1, $2, $3)
		RETURNING booking_id
	`, booking.EntryDate, booking.ExitDate, booking.AccomodationID).Scan(&booking.BookingID)
}

func (s *PostgresBookingStore) Update(ctx context.Context, booking *models.Booking) error {
	_, err := s.pool.Exec(ctx, `
		UPDATE bookings
		SET entry_date = $1, leaving_date = $2, acc_id = $3
		WHERE booking_id = $4
	`, booking.EntryDate, booking.ExitDate, booking.AccomodationID, booking.BookingID)
	return err
}

func (s *PostgresBookingStore) Delete(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, `DELETE FROM bookings WHERE booking_id = $1`, id)
	return err
}
