package test

import (
	"github.com/emtabb/qugo"
	"github.com/emtabb/qugo/qu"
	"github.com/emtabb/qugo/qu/impl"
	"github.com/emtabb/state"
	"log"
	"testing"
)

type DataImplementState struct {
	strField string
	intField int
}

func seedingDataInterface() []interface{} {
	return impl.ToInterfaces(seedingDataStates().ToArray())
}

func seedingDataStates() state.States {
	states := []state.State {
		DataImplementState{strField: "field1", intField: 1},
		DataImplementState{strField: "field2", intField: 2},
		DataImplementState{strField: "field3", intField: 3},
		DataImplementState{strField: "field4", intField: 4},
		DataImplementState{strField: "field5", intField: 5},
		DataImplementState{strField: "field6", intField: 6},
		DataImplementState{strField: "field7", intField: 7},
		DataImplementState{strField: "field8", intField: 8},
		DataImplementState{strField: "field9", intField: 9},
		DataImplementState{strField: "field10", intField: 10},
	}
	return new(state.List).Generate().ByStates(states)
}

func functionForTestMapLimitOffsetData(_state state.State) state.State {
	changeData := _state.(DataImplementState)
	changeData.intField++
	changeData.strField += " new field"
	newArrayStates := []state.State {
		changeData,
	}
	return qugo.Operator().InitStates(new(state.List).ByStates(newArrayStates)).Collect()
}

func functionForTestMapPipe(_state state.State) {
	log.Println(_state)
}

func TestInitStates(t *testing.T) {
	qu := qugo.Operator().InitStates(seedingDataStates()).Collect()
	log.Println(qu)
}

func TestLimitOffsetData(t *testing.T) {
	qu := qugo.Operator().InitStates(seedingDataStates()).Map(functionForTestMapLimitOffsetData).
		Limit(9).Collect()

	for _, element := range qu.ToArray() {
		log.Println("Show: ", element)
	}
}

func TestMapPipe(t *testing.T) {
	qu := qugo.Operator().InitStates(seedingDataStates()).Pipe(functionForTestMapPipe).
		Limit(1).Collect()

	for _, element := range qu.ToArray() {
		log.Println(element)
	}
}

func TestInitInterfaces(t *testing.T) {
	qu := qugo.Operator().ByInterfaces(seedingDataInterface()).
		Map(functionForTestMapLimitOffsetData).
		Limit(6).
		CollectInterface()
	qu2 := qugo.Operator().ByInterfaces(qu).Collect()
	for i := range qu {
		log.Println(qu[i])
	}

	log.Println("=========================")
	for i := range qu2.ToArray() {
		log.Println(qu[i])
	}
}

func functionForFlatMap(_state state.State) qu.Quantum {
	changeData := _state.(DataImplementState)
	changeData.intField++
	changeData.strField += " new field"
	newArrayStates := []state.State {
		changeData,
	}
	return qugo.Operator().InitStates(new(state.List).ByStates(newArrayStates))
}


func TestFlatMap(t *testing.T) {
	qu := qugo.Operator().ByInterfaces(seedingDataInterface()).
		FlatMap(functionForFlatMap).
		Limit(6).Skip(1).
		CollectInterface()
	for i := range qu {
		log.Println(qu[i])
	}
}

func TestFlatMapForEach(t *testing.T)  {
	qugo.Operator().ByInterfaces(seedingDataInterface()).
		FlatMap(functionForFlatMap).
		Limit(6).Skip(1).
		ForEach(func(s state.State) {
			space := s.(DataImplementState)
			log.Println(space.strField, space.intField)
		})
}

func TestMapForEach(t *testing.T)  {
	qugo.Operator().ByInterfaces(seedingDataInterface()).
		Map(functionForTestMapLimitOffsetData).
		Limit(6).Skip(1).
		ForEach(func(s state.State) {
			log.Println(s)
		})
}

func TestFilter(t *testing.T)  {
	qu := qugo.Operator().ByInterfaces(seedingDataInterface()).
		Filter(func(s state.State) bool {
			changeData := s.(DataImplementState)
			return changeData.intField > 5
		}).Collect()
	log.Println(qu)
}

func TestIndex(t *testing.T) {
	qu := qugo.Operator().ByInterfaces(seedingDataInterface()).
		Index(DataImplementState{strField: "field5", intField: 5})
	log.Println(qu)
}