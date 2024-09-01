package models

import "time"

type Booking struct {
	BookingID      int
	AccomodationID int
	EntryDate      time.Time
	ExitDate       time.Time
}

func NewBooking(bookingID int, accomodationID int, entryDate time.Time, exitDate time.Time) *Booking {
	return &Booking{
		BookingID:      bookingID,
		AccomodationID: accomodationID,
		EntryDate:      entryDate,
		ExitDate:       exitDate,
	}
}
