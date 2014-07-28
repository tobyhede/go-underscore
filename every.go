package un

import (
	"reflect"
)

func init() {
	MakeEvery(&Every)
}

// Every func(func(A, bool), bool)
// Returns true if all values in the collection (slice or map) pass the predicate truth test
// Note: each does not return a value, you may want un.Map
var Every func(func(value interface{}) bool, interface{}) bool

// MakeEach implements a typed Each function in the form Each func(func(A, B), []A)
func MakeEvery(fn interface{}) {
	Maker(fn, every)
}

func every(values []reflect.Value) []reflect.Value {
	fn := values[0]
	list := interfaceToValue(values[1])

	var ret bool
	if list.Kind() == reflect.Map {
		ret = everyMap(fn, list)
	}

	if list.Kind() == reflect.Slice {
		ret = everySlice(fn, list)
	}

	return Valueize(reflect.ValueOf(ret))
}

func everySlice(fn, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		res := fn.Call(Valueize(v))
		if ok := res[0].Interface().(bool); !ok {
			return false
		}
	}
	return true
}

func everyMap(fn, m reflect.Value) bool {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		res := fn.Call(Valueize(v))
		if ok := res[0].Interface().(bool); !ok {
			return false
		}
	}
	return true
}
