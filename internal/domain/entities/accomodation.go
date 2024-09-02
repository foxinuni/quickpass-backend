package entities

type Accomodation struct {
	AccomodationID int    `json:"accomodation_id"`
	IsHouse        bool   `json:"is_house"`
	Address        string `json:"address"`
}

func NewAccomodation(accomodationID int, isHouse bool, address string) *Accomodation {
	return &Accomodation{
		AccomodationID: accomodationID,
		IsHouse:        isHouse,
		Address:        address,
	}
}

func (a *Accomodation) GetAccomodationID() int {
	return a.AccomodationID
}

func (a *Accomodation) GetIsHouse() bool {
	return a.IsHouse
}

func (a *Accomodation) GetAddress() string {
	return a.Address
}

func (a *Accomodation) SetAccomodationID(accomodationID int) {
	a.AccomodationID = accomodationID
}

func (a *Accomodation) SetIsHouse(isHouse bool) {
	a.IsHouse = isHouse
}

func (a *Accomodation) SetAddress(address string) {
	a.Address = address
}
