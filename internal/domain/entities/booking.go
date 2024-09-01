package entities

import "time"

type Booking struct {
	BookingID    int
	Accomodation *Accomodation
	EntryDate    time.Time
	ExitDate     time.Time
}

func NewBooking(bookingID int, accomodation *Accomodation, entryDate time.Time, exitDate time.Time) *Booking {
	return &Booking{
		BookingID:    bookingID,
		Accomodation: accomodation,
		EntryDate:    entryDate,
		ExitDate:     exitDate,
	}
}

func (b *Booking) GetBookingID() int {
	return b.BookingID
}

func (b *Booking) GetAccomodation() *Accomodation {
	return b.Accomodation
}

func (b *Booking) GetEntryDate() time.Time {
	return b.EntryDate
}

func (b *Booking) GetExitDate() time.Time {
	return b.ExitDate
}

func (b *Booking) SetBookingID(bookingID int) {
	b.BookingID = bookingID
}

func (b *Booking) SetAccomodation(accomodation *Accomodation) {
	b.Accomodation = accomodation
}

func (b *Booking) SetEntryDate(entryDate time.Time) {
	b.EntryDate = entryDate
}

func (b *Booking) SetExitDate(exitDate time.Time) {
	b.ExitDate = exitDate
}
