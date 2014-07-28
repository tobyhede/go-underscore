package un

import (
	"reflect"
	"sync"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
	MakeEach(&EachStringInt)
}

// Each func(func(A, B), []A)
// Applies the given iterator function to each element of a collection (slice or map).
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Note: each does not return a value, you may want un.Map
var Each func(func(value interface{}, i interface{}), interface{})

// EachInt
// Applies the given iterator function to each element of []int
// Iterator function arguments are *value, index*
var EachInt func(func(value, i int), []int)

// EachStringInt
// Applies the given iterator function to each element of map[string]int
// Iterator function arguments are *value, key*
var EachStringInt func(func(value int, key string), map[string]int)

// MakeEach implements a typed Each function in the form Each func(func(A, B), []A)
func MakeEach(fn interface{}) {
	Maker(fn, each)
}

func each(values []reflect.Value) []reflect.Value {
	fn := values[0]
	list := interfaceToValue(values[1])

	if list.Kind() == reflect.Map {
		eachMap(fn, list)
	}

	if list.Kind() == reflect.Slice {
		eachSlice(fn, list)
	}

	return nil
}

func eachSlice(fn, s reflect.Value) {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		fn.Call(Valueize(v, reflect.ValueOf(i)))
	}
}

func eachMap(fn, m reflect.Value) {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		fn.Call(Valueize(v, k))
	}
}

// WIP
func _pEach(values []reflect.Value) []reflect.Value {
	// var done sync.WaitGroup

	fn := values[0]

	v := interfaceToValue(values[1])

	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		fn.Call([]reflect.Value{e})
	}

	return nil
}

// Reference Each Implementation
func refEach(slice []string, fn func(string)) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i])
	}
}

// Reference Parallel Each Implementation
func refPEach(slice []string, fn func(string)) {
	var done sync.WaitGroup
	for _, s := range slice {
		s := s
		done.Add(1)
		go func() {
			fn(s)
			done.Done()
		}()
	}

	done.Wait()
}
