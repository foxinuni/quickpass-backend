package models

type Accomodation struct {
	AccomodationID int
	IsHouse        bool
	Address        string
}

func NewAccomodation(accomodationID int, isHouse bool, address string) *Accomodation {
	return &Accomodation{
		AccomodationID: accomodationID,
		IsHouse:        isHouse,
		Address:        address,
	}
}
