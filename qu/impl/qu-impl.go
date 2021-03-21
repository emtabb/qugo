package impl;

import "github.com/emtabb/state"
import . "github.com/emtabb/qugo/qu"

type QuantumImpl struct {
	Quantum
	states []state.State
	skip int32
	limit int32
	tensor [][]state.State
}

func (quantum *QuantumImpl) ByInterfaces(interfaces []interface{}) Quantum {
	toStates := make([]state.State, len(interfaces))
	for i, _state := range interfaces {
		toStates[i] = _state.(state.State)
	}
	quantum.states = toStates[:]
	return quantum
}

func (quantum *QuantumImpl) Init(states []state.State) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum *QuantumImpl) Of(states ...state.State) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum *QuantumImpl) Map(f func (state.State) state.State) Quantum {
	for i, state := range quantum.states {
		quantum.states[i] = f(state)
	}
	return quantum
}

func (quantum *QuantumImpl) FlatMap(f func ([]state.State) []state.State) Quantum {
	return quantum
}

func (quantum *QuantumImpl) Reduce(reductState state.State, f func (state1 state.State, state2 state.State) state.State) Quantum {
	for i, state := range quantum.states { 
		quantum.states[i] = f(state, reductState)
	}
	return quantum
}

func (quantum *QuantumImpl) Quantized(f func (state.State) []state.State) Quantum {

	return quantum
}

func (quantum *QuantumImpl) Index(obsState state.State) int32 {
	for i, state := range quantum.states {
		if quantum.Equal(state, obsState) {
			return int32(i)
		}
	}
	return -1
}

func (quantum *QuantumImpl) Include(obsState state.State) bool {
	return quantum.Index(obsState) >= 0
}

func (quantum *QuantumImpl) Skip(skip int32) Quantum {
	quantum.skip = skip
	return quantum
}

func (quantum *QuantumImpl) Limit(limit int32) Quantum {
	quantum.limit = limit
	return quantum
}

func (quantum *QuantumImpl) Sorted(f func(state.State) state.State) Quantum {

	return quantum
}

func (quantum *QuantumImpl) Equal(state state.State, obsState state.State) bool {
	if state.ToString() == obsState.ToString() {
		return true
	}
	return false
}

func (quantum *QuantumImpl) Pipe(f func(state.State)) Quantum {
	for _, state := range quantum.states {
		f(state)
	}
	return quantum
}

// This Method will update with Tree Search
func (quantum *QuantumImpl) Filter(f func(state.State) bool) Quantum {
	filterStates := make([]state.State, 0)
	for _, state := range quantum.states {
		if f(state) {
			filterStates = append(filterStates, state)
		}
	}
	quantum.states = filterStates[:]
	return quantum
}

func (quantum *QuantumImpl) Collect() []state.State {
	
	return quantum.pageable(quantum.states)
}

func (quantum *QuantumImpl) pageable(states []state.State) []state.State {
	// if quantum.limit == 0 {
	// 	quantum.limit = 500
	// }

	// quantum.states = quantum.states[quantum.skip:quantum.limit]
	// scopeLimit := int32(len(quantum.states)) > quantum.limit ? int32(len(quantum.states)) : quantum.limit
	return states
}

func (quantum *QuantumImpl) pageableInterface(states []state.State) []interface{} {
	interfaces := make([]interface{}, len(quantum.states))
	for i, state := range states {
		interfaces[i] = state
	}
	return interfaces
}

func (quantum *QuantumImpl) CollectInterface() []interface{} {
	if quantum.limit == 0 {
		quantum.limit = 500
	}
	quantum.states = quantum.states[quantum.skip:quantum.limit]
	return quantum.pageableInterface(quantum.states)
}