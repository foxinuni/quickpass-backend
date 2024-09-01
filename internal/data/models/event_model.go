package models

import "time"

type Event struct {
	EventID   string
	Name      string
	Address   string
	StartDate time.Time
	EndDate   time.Time
}

func NewEvent(eventID string, name string, address string, startDate time.Time, endDate time.Time) *Event {
	return &Event{
		EventID:   eventID,
		Name:      name,
		Address:   address,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
