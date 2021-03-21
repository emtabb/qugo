package qu

import . "github.com/emtabb/state"

type Quantum interface {
	ByInterfaces([]interface{}) Quantum
	Collect() []State
	CollectInterface() []interface{}
	Init([]State) Quantum
	Index(State) int32
	Include(State) bool 
	Of(...State) Quantum
	Map(func (State) State) Quantum
	FlatMap(func ([]State) []State) Quantum
	Reduce(State, func (State, State) State) Quantum
	Filter(func (State) bool) Quantum
	Quantized(func (State) []State) Quantum
	Skip(int32) Quantum
	Limit(int32) Quantum
	Sorted(func(State) State) Quantum
	Pipe(func(State)) Quantum
	Equal(State, State) bool
}