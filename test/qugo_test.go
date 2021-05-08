package test

import (
	"encoding/json"
	"github.com/emtabb/qugo"
	"github.com/emtabb/qugo/qu"
	"github.com/emtabb/qugo/qu/impl"
	"github.com/emtabb/state"
	"log"
	"testing"
)

type DataImplementState struct {
	StrField string `json:"str_field"`
	IntField int `json:"int_field"`
}

func seedingDataInterface() []interface{} {
	return impl.ToInterfaces(seedingDataStates().ToArray())
}

func seedingDataInterfaceByInteger() []interface{} {
	interfaces := make([]interface{}, 0)
	integers := []int { 1, 2, 3, 4, 5, 6, 7 ,8 ,9 ,10}
	for i := range integers {
		interfaces = append(interfaces, integers[i])
	}
	return interfaces
}

func seedingDataStates() state.States {
	states := []state.State {
		DataImplementState{StrField: "field1", IntField: 1},
		DataImplementState{StrField: "field2", IntField: 2},
		DataImplementState{StrField: "field3", IntField: 3},
		DataImplementState{StrField: "field4", IntField: 4},
		DataImplementState{StrField: "field5", IntField: 5},
		DataImplementState{StrField: "field6", IntField: 6},
		DataImplementState{StrField: "field7", IntField: 7},
		DataImplementState{StrField: "field8", IntField: 8},
		DataImplementState{StrField: "field9", IntField: 9},
		DataImplementState{StrField: "field10", IntField: 10},
	}
	return new(state.List).Generate().ByStates(states)
}

func functionForTestMapLimitOffsetData(_state state.State) state.State {
	changeData := _state.(DataImplementState)
	changeData.IntField++
	changeData.StrField += " new field"
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
		log.Println(qu2.ToArray()[i])
	}
}

func functionForFlatMap(_state state.State) qu.Quantum {
	changeData := _state.(DataImplementState)
	changeData.IntField++
	changeData.StrField += " new field"
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
			log.Println(space.StrField, space.IntField)
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
			return changeData.IntField > 5
		}).Collect()
	log.Println(qu)
}

func TestIndex(t *testing.T) {
	qu := qugo.Operator().ByInterfaces(seedingDataInterface()).
		Index(DataImplementState{StrField: "field5", IntField: 5})
	log.Println(qu)
}

func TestIntegerInterfaces(t *testing.T) {
	qugo.Operator().ByInterfaces(seedingDataInterfaceByInteger()).
		ForEach(func(state_ state.State) {
			log.Println(state_)
	})
}

func TestMapWithJsonParseForEach(t *testing.T)  {
	qugo.Operator().ByInterfaces(seedingDataInterface()).
		Map(func(s state.State) state.State {
			changeData := s.(DataImplementState)
			data, _ := json.Marshal(changeData)
			log.Println("data ", data)
			var revertData DataImplementState
			json.Unmarshal(data, &revertData)
			log.Println("revert data ", revertData)
			return s
		}).
		Limit(6).Skip(1).
		ForEach(func(s state.State) {
			log.Println(s)
		})
}