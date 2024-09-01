package models

import "time"

type Booking struct {
	BookingID      int
	AccomodationID int
	EntryDate      time.Time
	ExitDate       time.Time
}
