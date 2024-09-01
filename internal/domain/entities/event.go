package entities

import "time"

type Event struct {
	EventID   string
	Name      string
	Address   string
	StartDate time.Time
	EndDate   time.Time
}
