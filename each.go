package un

import (
	"reflect"
	"sync"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
}

// Each func(func(A), []A)
// Applies the given function to each item of a slice or map
// Note: unlike map, each does not return a collection
var Each func(func(interface{}), interface{})

// EachInt on a slice of ints
var EachInt func(func(int), []int)

// MakeEach implements a typed Each function in the form Each func(func(A), []A)
func MakeEach(fn interface{}) {
	Maker(fn, _each)
}

func _each(values []reflect.Value) []reflect.Value {
	fn := values[0]
	list := interfaceToValue(values[1])

	if list.Kind() == reflect.Map {
		for _, v := range list.MapKeys() {
			fn.Call([]reflect.Value{v})
		}
	}

	if list.Kind() == reflect.Slice {
		for i := 0; i < list.Len(); i++ {
			e := list.Index(i)
			fn.Call([]reflect.Value{e})
		}
	}

	return nil
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

// type empty {}
// ...
// data := make([]float, N);
// res := make([]float, N);
// sem := make(chan empty, N);  // semaphore pattern
// ...
// for i,xi := range data {
//     go func (i int, xi float) {
//         res[i] = doSomething(i,xi);
//         sem <- empty{};
//     } (i, xi);
// }
// // wait for goroutines to finish
// for i := 0; i < N; ++i { <-sem }
