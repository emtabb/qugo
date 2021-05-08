package qu

import "github.com/emtabb/state"

type Quantum interface {
	state.State
	All() state.States
	ByInterfaces([]interface{}) Quantum
	Collect() state.States
	CollectInterface() []interface{}
	InitStates(state.States) Quantum
	Index(state.State) int32
	Include(state.State) bool
	ForEach(func (state.State))
	Filter(func (state.State) bool) Quantum
	FlatMap(func (state.State) Quantum) Quantum
	Of(...state.State) Quantum
	Map(func (state.State) state.State) Quantum
	Reduce(func (state.State, state.State) state.State) Quantum
	Quantized(func (state.State) state.States) Quantum
	Skip(int32) Quantum
	Limit(int32) Quantum
	Sorted(func(state.State, state.State) bool) Quantum
	Pipe(func(state.State)) Quantum
}