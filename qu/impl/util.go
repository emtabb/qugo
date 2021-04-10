package impl

import (
	"github.com/emtabb/state"
)

func ToInterface(_state state.State) interface{} {
	data := _state
	return data
}

func ToInterfaces(states []state.State) []interface{} {
	data := make([]interface{}, len(states))
	for i, _state := range states {
		data[i] = ToInterface(_state)
	}
	return data
}
