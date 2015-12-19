package engine

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

type State uint

type Stated interface {
	States() map[State]bool
}

var States StateManager = StateManager{}

type StateManager struct {
	Current State
}

func (s *StateManager) Switch(newState State) {
	stateExit()
	s.Current = newState
	stateEntry()
}

func (s *StateManager) SystemInCurrentState(system interface{}) bool {
	if statedSystem, ok := system.(Stated); ok {
		if active, exists := statedSystem.States()[s.Current]; !(exists && active) {
			// this system's not active in this state
			return false
		}
	}
	// this system's active in this state (or all states)
	return true
}
