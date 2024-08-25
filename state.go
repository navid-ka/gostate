package state

type State interface {
	SwitchState(key string)
	Get(key string) bool
	InitState(initialStates map[string]bool)
}

type state struct {
	State map[string]bool
}

// NewState returns a new state object
func NewState() *state {
	return &state{State: make(map[string]bool)}
}

// InitStates accepts a map of string keys and boolean values to initialize any state
func (s *state) InitState(initialStates map[string]bool) {
	for key, value := range initialStates {
		s.State[key] = value
	}
}
// Get returns the value of the key in the state map
func (s *state) Get(key string) bool {
	return s.State[key]
}

// SwitchState sets the value of the key to true and all other keys to false
func (s *state) SwitchState(key string) {
	for k := range s.State {
		s.State[k] = false
	}
	s.State[key] = true
}