package un

import (
	"reflect"
)

func init() {
	MakeNone(&None)
	MakeNone(&NoneInt)
	MakeNone(&NoneString)
}

// None func(func(A, bool), bool)
// Returns true if all values in the collection (slice or map) pass the predicate truth test

// var None func(func(value interface{}) bool, interface{}) bool
var None func(fn, slice_or_map interface{}) bool

// NoneInt
// Returns true if all values in a []int pass the predicate truth test
// Predicate function accepts an int and returns a boolean
var NoneInt func(func(value int) bool, []int) bool

// NoneString
// Returns true if all values in a []string pass the predicate truth test
// Predicate function accepts a string and returns a boolean
var NoneString func(func(value string) bool, []string) bool

// MakeEach implements a typed Each function in the form Each func(func(A, B), []A)
func MakeNone(fn interface{}) {
	Maker(fn, none)
}

func none(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	var ret bool
	if col.Kind() == reflect.Map {
		ret = noneMap(fn, col)
	}

	if col.Kind() == reflect.Slice {
		ret = noneSlice(fn, col)
	}

	return Valueize(reflect.ValueOf(ret))
}

func noneSlice(fn, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		if ok := callPredicate(fn, v); ok {
			return false
		}
	}
	return true
}

func noneMap(fn, m reflect.Value) bool {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		if ok := callPredicate(fn, v); ok {
			return false
		}
	}
	return true
}
