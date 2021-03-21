package qugo

import . "github.com/emtabb/qugo/qu"
import . "github.com/emtabb/qugo/qu/impl"

func Operator() Quantum {
	return new(QuantumImpl)
}