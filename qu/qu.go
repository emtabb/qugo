package qu

import "github.com/emtabb/state"

type Quantum interface {
	All() []state.State
	Collect() []state.State
	Init([]state.State) Quantum
	Index(state.State) int32
	Include(state.State) bool
	Of(...state.State) Quantum
	Map(func (state.State) state.State) Quantum
	FlatMap(func ([]state.State) []state.State) Quantum
	Reduce(state.State, func (state.State, state.State) state.State) Quantum
	Filter(func (state.State) bool) Quantum
	Quantized(func (state.State) []state.State) Quantum
	Skip(int32) Quantum
	Limit(int32) Quantum
	Sorted(func(state.State) state.State) Quantum
	Pipe(func(state.State)) Quantum
	Equal(state.State, state.State) bool
}