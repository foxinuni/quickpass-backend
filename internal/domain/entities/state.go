package entities

type State struct {
	StateID   int    `json:"state_id"`
	StateName string `json:"state_name"`
}

func NewState(stateID int, stateName string) *State {
	return &State{
		StateID:   stateID,
		StateName: stateName,
	}
}

func (s *State) GetStateID() int {
	return s.StateID
}

func (s *State) GetStateName() string {
	return s.StateName
}

func (s *State) SetStateID(stateID int) {
	s.StateID = stateID
}

func (s *State) SetStateName(stateName string) {
	s.StateName = stateName
}
