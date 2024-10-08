package entities

type Occasion struct {
	OccasionID int      `json:"occasion_id"`
	User       *User    `json:"user"`
	Event      *Event   `json:"event"`
	Booking    *Booking `json:"booking"`
	State      *State   `json:"state"`
	Inside     bool     `json:"inside"`
}

func NewOccasion(occasionID int, user *User, event *Event, booking *Booking, state *State, isInside bool) *Occasion {
	return &Occasion{
		OccasionID: occasionID,
		User:       user,
		Event:      event,
		Booking:    booking,
		State:      state,
		Inside:     isInside,
	}
}

func (o *Occasion) GetOccasionID() int {
	return o.OccasionID
}

func (o *Occasion) GetUser() *User {
	return o.User
}

func (o *Occasion) GetEvent() *Event {
	return o.Event
}

func (o *Occasion) GetBooking() *Booking {
	return o.Booking
}

func (o *Occasion) GetState() *State {
	return o.State
}

func (o *Occasion) SetOccasionID(occasionID int) {
	o.OccasionID = occasionID
}

func (o *Occasion) SetUser(user *User) {
	o.User = user
}

func (o *Occasion) SetEvent(event *Event) {
	o.Event = event
}

func (o *Occasion) SetBooking(booking *Booking) {
	o.Booking = booking
}

func (o *Occasion) SetState(state *State) {
	o.State = state
}
