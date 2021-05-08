package qugo

import "github.com/emtabb/qugo/qu"
import "github.com/emtabb/qugo/qu/impl"

func Operator() qu.Quantum {
	return new(impl.QuantumImpl)
}

var INFINITY = impl.INFINITY
var UN_LIMITED = impl.UN_LIMITED