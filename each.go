package un

import (
	"reflect"
	"sync"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
}

// Each func(func(A B), []A)
// Applies the given function to each item of a slice or map
// Note: unlike map, each does not return a collection
var Each func(func(interface{}, interface{}), interface{})

// EachInt on a slice of ints
var EachInt func(func(value, i int), []int)

// MakeEach implements a typed Each function in the form Each func(func(A), []A)
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
		fn.Call(Valueize(k, v))
	}
}

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
func RefEach(slice []string, fn func(string)) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i])
	}
}

// Reference Parallel Each Implementation
func RefPEach(slice []string, fn func(string)) {
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
