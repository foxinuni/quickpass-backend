package entities

import "time"

type Event struct {
	EventID   int       `json:"event_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func NewEvent(eventID int, name string, address string, startDate time.Time, endDate time.Time) *Event {
	return &Event{
		EventID:   eventID,
		Name:      name,
		Address:   address,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func (e *Event) GetEventID() int {
	return e.EventID
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetAddress() string {
	return e.Address
}

func (e *Event) GetStartDate() time.Time {
	return e.StartDate
}

func (e *Event) GetEndDate() time.Time {
	return e.EndDate
}

func (e *Event) SetEventID(eventID int) {
	e.EventID = eventID
}

func (e *Event) SetName(name string) {
	e.Name = name
}

func (e *Event) SetAddress(address string) {
	e.Address = address
}

func (e *Event) SetStartDate(startDate time.Time) {
	e.StartDate = startDate
}

func (e *Event) SetEndDate(endDate time.Time) {
	e.EndDate = endDate
}
