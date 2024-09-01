package models

type State struct {
	StateID   int
	StateName string
}

func NewState(stateID int, stateName string) *State {
	return &State{
		StateID:   stateID,
		StateName: stateName,
	}
}
