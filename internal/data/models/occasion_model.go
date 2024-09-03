package models

type Occasion struct {
	OccasionID int
	UserID     int
	EventID    *int
	BookingID  *int
	StateID    int
}

func NewOccasion(occasionID int, userID int, eventID *int, bookingID *int, stateID int) *Occasion {
	return &Occasion{
		OccasionID: occasionID,
		UserID:     userID,
		EventID:    eventID,
		BookingID:  bookingID,
		StateID:    stateID,
	}
}
