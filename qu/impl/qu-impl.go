package impl

import (
	. "github.com/emtabb/qugo/qu"
	"github.com/emtabb/state"
)

type QuantumImpl struct {
	Quantum
	states []state.State

	skip int32
	limit int32
	tensor []state.States
}

func (quantum *QuantumImpl) All() state.States {
	return new(state.List).ByStates(quantum.states)
}

func (quantum *QuantumImpl) ByInterfaces(interfaces []interface{}) Quantum {
	quantum.states = make([]state.State, len(interfaces))
	for i := range interfaces {
		quantum.states[i] = interfaces[i].(state.State)
	}

	return quantum
}

func (quantum *QuantumImpl) InitStates(states state.States) Quantum {
	quantum.states = make([]state.State, states.Size())
	quantum.states = states.ToArray()[:]
	return quantum
}

func (quantum *QuantumImpl) Of(states ...state.State) Quantum {
	quantum.states = make([]state.State, len(states))
	quantum.states = states[:]
	return quantum
}

func (quantum *QuantumImpl) Map(f func (state.State) state.State) Quantum {
	for i, _state := range quantum.states {
		quantum.states[i] = f(_state)
	}
	return quantum
}

func (quantum *QuantumImpl) FlatMap(f func (state.State) Quantum) Quantum {
	for i := range quantum.states {
		quantum.tensor = append(quantum.tensor, f(quantum.states[i]).Collect())
	}
	flatState := make([]state.State, 0)
	for _, tensor := range quantum.tensor {
		flatState = append(flatState, tensor.ToArray()...)
	}
	quantum.states = flatState[:]
	return quantum
}

func (quantum *QuantumImpl) ForEach(f func(state.State)) {
	for _, _state := range quantum.pageable(new(state.List).ByStates(quantum.states)).ToArray() {
		f(_state)
	}
}

func (quantum *QuantumImpl) Reduce(reduceState state.State, f func (state1 state.State, state2 state.State) state.State) Quantum {
	for i, _state := range quantum.states {
		quantum.states[i] = f(reduceState, _state)
	}
	return quantum
}

func (quantum *QuantumImpl) Quantized(f func (state.State) state.States) Quantum {
	for _, _state := range quantum.states {
		quantum.tensor = append(quantum.tensor, f(_state))
	}
	return quantum
}

func (quantum *QuantumImpl) Index(obsState state.State) int32 {
	for i, _state := range quantum.states {
		if _state == obsState {
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

func (quantum *QuantumImpl) Sorted(f func(state.State, state.State) bool) Quantum {
	return quantum
}

func (quantum *QuantumImpl) Pipe(f func(state.State)) Quantum {
	for _, _state := range quantum.states {
		f(_state)
	}
	return quantum
}

// This Method will update with Tree Search
func (quantum *QuantumImpl) Filter(f func(state.State) bool) Quantum {
	filterStates := make([]state.State, 0)
	for _, _state := range quantum.states {
		if f(_state) {
			filterStates = append(filterStates, _state)
		}
	}
	quantum.states = filterStates[:]
	return quantum
}

func (quantum *QuantumImpl) Collect() state.States {
	return quantum.pageable(new(state.List).ByStates(quantum.states))
}

func (quantum *QuantumImpl) CollectInterface() []interface{} {
	return quantum.pageableInterfaces(quantum.states)
}

//============================================================================

func (quantum *QuantumImpl) pageable(states state.States) state.States {
	quantum.handleLimitOffset(states.ToArray())
	if quantum.skip == 0 {
		quantum.states = states.ToArray()[: quantum.limit]
	} else {
		quantum.states = states.ToArray()[quantum.skip: quantum.limit]
	}
	outState := new(state.List)
	outState.ByStates(quantum.states)
	return outState
}

func (quantum *QuantumImpl) pageableInterfaces(states []state.State) []interface{} {
	quantum.handleLimitOffset(states)
	if quantum.skip == 0 {
		quantum.states = states[:quantum.limit]
	} else {
		quantum.states = states[quantum.skip - 1:quantum.limit]
	}
	return ToInterfaces(quantum.states)
}

func (quantum *QuantumImpl) handleLimitOffset(states []state.State) {
	if quantum.limit == 0 {
		quantum.limit = 500
	}

	scopeLimit := int32(len(states))
	if scopeLimit < quantum.limit {
		quantum.limit = scopeLimit
	}
}
