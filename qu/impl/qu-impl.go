package impl;

import . "github.com/emtabb/qugo/qu"

type QuantumImpl struct {
	Quantum
	states []interface{}
	skip int32
	limit int32
	tensor [][]interface{}
}

func (quantum *QuantumImpl) Init(states []interface{}) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum *QuantumImpl) Of(states ...interface{}) Quantum {
	quantum.states = states[:]
	return quantum
}

func (quantum *QuantumImpl) Map(f func (interface{}) interface{}) Quantum {
	for i, state := range quantum.states {
		quantum.states[i] = f(state)
	}
	return quantum
}

func (quantum *QuantumImpl) FlatMap(f func ([]interface{}) []interface{}) Quantum {
	return quantum
}

func (quantum *QuantumImpl) Reduce(reduceState interface{}, f func (state1 interface{}, state2 interface{}) interface{}) Quantum {
	for i, state := range quantum.states { 
		quantum.states[i] = f(state, reduceState)
	}
	return quantum
}

func (quantum *QuantumImpl) Quantized(f func (interface{}) []interface{}) Quantum {

	return quantum
}

func (quantum *QuantumImpl) Index(obsState interface{}) int32 {
	for i, state := range quantum.states {
		if quantum.Equal(state, obsState) {
			return int32(i)
		}
	}
	return -1
}

func (quantum *QuantumImpl) Include(obsState interface{}) bool {
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

func (quantum *QuantumImpl) Sorted(f func(interface{}) interface{}) Quantum {

	return quantum
}

func (quantum *QuantumImpl) Equal(state interface{}, obsState interface{}) bool {
	if state == obsState {
		return true
	}
	return false
}

func (quantum *QuantumImpl) Pipe(f func(interface{})) Quantum {
	for _, state := range quantum.states {
		f(state)
	}
	return quantum
}

// This Method will update with Tree Search
func (quantum *QuantumImpl) Filter(f func(interface{}) bool) Quantum {
	filterStates := make([]interface{}, 0)
	for _, state := range quantum.states {
		if f(state) {
			filterStates = append(filterStates, state)
		}
	}
	quantum.states = filterStates[:]
	return quantum
}

func (quantum *QuantumImpl) Collect() []interface{} {
	return quantum.pageable(quantum.states)
}

func (quantum *QuantumImpl) pageable(states []interface{}) []interface{} {
	if quantum.limit == 0 {
		quantum.limit = 500
	}
	scopeLimit := int32(0)
	lenState := int32(len(quantum.states))
	if lenState < quantum.limit {
		scopeLimit = lenState
	} else {
		scopeLimit = quantum.limit
	}

	quantum.states = quantum.states[quantum.skip:scopeLimit]
	return states
}

func (quantum *QuantumImpl) All() []interface{} {
	return quantum.states
}