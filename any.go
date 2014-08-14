package un

import (
	"reflect"
)

func init() {
	MakeAny(&Any)
	MakeAny(&AnyInt)
	MakeAny(&AnyString)
}

// Any func(func(A, bool), bool)
// Returns true if all values in the collection (slice or map) pass the predicate truth test

// var Any func(func(value interface{}) bool, interface{}) bool
var Any func(fn, slice_or_map interface{}) bool

// AnyInt
// Returns true if all values in a []int pass the predicate truth test
// Predicate function accepts an int and returns a boolean
var AnyInt func(func(value int) bool, []int) bool

// AnyString
// Returns true if all values in a []string pass the predicate truth test
// Predicate function accepts a string and returns a boolean
var AnyString func(func(value string) bool, []string) bool

// MakeEach implements a typed Each function in the form Each func(func(A, B), []A)
func MakeAny(fn interface{}) {
	Maker(fn, any)
}

func any(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	var ret bool
	if col.Kind() == reflect.Map {
		ret = anyMap(fn, col)
	}

	if col.Kind() == reflect.Slice {
		ret = anySlice(fn, col)
	}

	return Valueize(reflect.ValueOf(ret))
}

func anySlice(fn, s reflect.Value) bool {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		if ok := callPredicate(fn, v); ok {
			return true
		}
	}
	return false
}

func anyMap(fn, m reflect.Value) bool {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		if ok := callPredicate(fn, v); ok {
			return true
		}
	}
	return false
}
