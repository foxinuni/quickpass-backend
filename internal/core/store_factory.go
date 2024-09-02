package core

import (
	"context"
	"errors"

	"github.com/foxinuni/quickpass-backend/internal/data/stores"
	"github.com/golang-migrate/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type StoreFactory interface {
	NewAccomodationStore() stores.AccomodationStore
	NewBookingStore() stores.BookingStore
	NewEventStore() stores.EventStore
	NewLogStore() stores.LogStore
	NewOccasionStore() stores.OccasionStore
	NewSessionStore() stores.SessionStore
	NewStateStore() stores.StateStore
	NewUserStore() stores.UserStore
}

type PgStoreFactoryOptions interface {
	GetDatabaseURL() string
	GetMigrationsSource() string
}

type PostgresStoreFactory struct {
	pool    *pgxpool.Pool
	options PgStoreFactoryOptions
}

func NewPostgresStoreFactory(options PgStoreFactoryOptions) (StoreFactory, error) {
	// Create the pgx connection pool
	log.Info().Msgf("Using postgres database (uri: %q, migrations: %q)", options.GetDatabaseURL(), options.GetMigrationsSource())
	pool, err := pgxpool.New(context.Background(), options.GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	// Run the migrations with go-migrate
	m, err := migrate.New(options.GetMigrationsSource(), options.GetDatabaseURL())
	if err != nil {
		return nil, err
	}

	log.Info().Msgf("Running migrations- This could take some time...")
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}

		log.Info().Msgf("Database is up to date! No new migrations to run.")

	}

	return &PostgresStoreFactory{
		options: options,
		pool:    pool,
	}, nil
}

func (f *PostgresStoreFactory) NewAccomodationStore() stores.AccomodationStore {
	return stores.NewPostgresAccomodationStore(f.pool)
}

func (f *PostgresStoreFactory) NewBookingStore() stores.BookingStore {
	return stores.NewPostgresBookingStore(f.pool)
}

func (f *PostgresStoreFactory) NewEventStore() stores.EventStore {
	return stores.NewPostgresEventStore(f.pool)
}

func (f *PostgresStoreFactory) NewLogStore() stores.LogStore {
	return stores.NewPostgresLogStore(f.pool)
}

func (f *PostgresStoreFactory) NewOccasionStore() stores.OccasionStore {
	return stores.NewPostgresOccasionStore(f.pool)
}

func (f *PostgresStoreFactory) NewSessionStore() stores.SessionStore {
	return stores.NewPostgresSessionStore(f.pool)
}

func (f *PostgresStoreFactory) NewStateStore() stores.StateStore {
	return stores.NewPostgresStateStore(f.pool)
}

func (f *PostgresStoreFactory) NewUserStore() stores.UserStore {
	return stores.NewPostgresUserStore(f.pool)
}

// ----------------- Store Builders -----------------
func BuildAccomoStore(factory StoreFactory) stores.AccomodationStore {
	return factory.NewAccomodationStore()
}

func BuildBookingStore(factory StoreFactory) stores.BookingStore {
	return factory.NewBookingStore()
}

func BuildEventStore(factory StoreFactory) stores.EventStore {
	return factory.NewEventStore()
}

func BuildLogStore(factory StoreFactory) stores.LogStore {
	return factory.NewLogStore()
}

func BuildOccasionStore(factory StoreFactory) stores.OccasionStore {
	return factory.NewOccasionStore()
}

func BuildSessionStore(factory StoreFactory) stores.SessionStore {
	return factory.NewSessionStore()
}

func BuildStateStore(factory StoreFactory) stores.StateStore {
	return factory.NewStateStore()
}

func BuildUserStore(factory StoreFactory) stores.UserStore {
	return factory.NewUserStore()
}
