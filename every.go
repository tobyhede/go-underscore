package un

import (
	"reflect"
)

func init() {
	MakeEvery(&Every)
	MakeEvery(&EveryInt)
	MakeEvery(&EveryString)
}

// Every func(func(A, bool), bool)
// Returns true if all values in the collection (slice or map) pass the predicate truth test

// var Every func(func(value interface{}) bool, interface{}) bool
var Every func(fn, slice_or_map interface{}) bool

// EveryInt
// Returns true if all values in a []int pass the predicate truth test
// Predicate function accepts an int and returns a boolean
var EveryInt func(func(value int) bool, []int) bool

// EveryString
// Returns true if all values in a []string pass the predicate truth test
// Predicate function accepts a string and returns a boolean
var EveryString func(func(value string) bool, []string) bool

// MakeEach implements a typed Each function in the form Each func(func(A, B), []A)
func MakeEvery(fn interface{}) {
	Maker(fn, every)
}

func every(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	var ret bool
	if col.Kind() == reflect.Map {
		ret = everyMap(fn, col)
	}

	if col.Kind() == reflect.Slice {
		ret = everySlice(fn, col)
	}

	return Valueize(reflect.ValueOf(ret))
}

func everySlice(fn, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		if ok := callPredicate(fn, v, reflect.ValueOf(i)); !ok {
			return false
		}
	}
	return true
}

func everyMap(fn, m reflect.Value) bool {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		if ok := callPredicate(fn, v); !ok {
			return false
		}
	}
	return true
}
