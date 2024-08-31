package core

import "github.com/foxinuni/quickpass-backend/internal/data/stores"

type StoreFactory interface {
	NewAccomodationStore() (*stores.AccomodationStore, error)
	NewBookingStore() (*stores.BookingStore, error)
	NewEventStore() (*stores.EventStore, error)
	NewLogStore() (*stores.LogStore, error)
	NewOccasionStore() (*stores.OccasionStore, error)
	NewSessionStore() (*stores.SessionStore, error)
	NewStateStore() (*stores.StateStore, error)
	NewUserStore() (*stores.UserStore, error)
}
