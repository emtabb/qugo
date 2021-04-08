package qu

type Quantum interface {
	All() []interface{}
	Collect() []interface{}
	Init([]interface{}) Quantum
	Index(interface{}) int32
	Include(interface{}) bool
	Of(...interface{}) Quantum
	Map(func (interface{}) interface{}) Quantum
	FlatMap(func ([]interface{}) []interface{}) Quantum
	Reduce(interface{}, func (interface{}, interface{}) interface{}) Quantum
	Filter(func (interface{}) bool) Quantum
	Quantized(func (interface{}) []interface{}) Quantum
	Skip(int32) Quantum
	Limit(int32) Quantum
	Sorted(func(interface{}) interface{}) Quantum
	Pipe(func(interface{})) Quantum
	Equal(interface{}, interface{}) bool
}