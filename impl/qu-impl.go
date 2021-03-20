package impl;

import "github.com/emtabb/state"

type Quantum struct {
	states []state.State

	tensor [][]state.State
}

func (quantum Quantum) Init(states []state.State) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum Quantum) Of(states ...state.State) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum Quantum) Map(f func (state.State) state.State) Quantum {
	for i, state := range quantum.states {
		quantum.states[i] = f(state)
	}
	return quantum
}

func (quantum Quantum) FlatMap(f func (state.State) state.State) Quantum {

	return quantum
}

func (quantum Quantum) Quantized(f func ([]state.State) []state.State) Quantum {

	return quantum
}

func (quantum Quantum) Index(obsState state.State) int {
	for i, state := range quantum.states {
		if quantum.Equal(state, obsState) {
			return i
		}
	}
	return -1
}

func (quantum Quantum) Include(obsState state.State) bool {
	return quantum.Index(obsState) >= 0
}

func (quantum Quantum) Skip(skip int) Quantum {

	return quantum
}

func (quantum Quantum) Limit(limit int) Quantum {

	return quantum
}

func (quantum Quantum) Sorted(f func(state.State) state.State) Quantum {

	return quantum
}

func (quantum Quantum) Equal(state state.State, obsState state.State) bool {
	if state.ToString() == obsState.ToString() {
		return true
	}
	return false
}

func (quantum Quantum) Pipe(f func(state.State)) {
	for _, state := range quantum.states {
		f(state)
	}
}

// This Method will update with Tree Search
func (quantum Quantum) Filter(f func(state.State) bool) Quantum {
	filterStates := make([]state.State, 0)
	for _, state := range quantum.states {
		if f(state) {
			filterStates = append(filterStates, state)
		}
	}
	quantum.states = filterStates[:]
	return quantum
}

func (quantum Quantum) Collect() []state.State {
	return quantum.states
}

func (quantum Quantum) CollectInterface() [] interface{} {

	return nil
}