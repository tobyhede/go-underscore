package un

import (
	"reflect"
	"sync"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
	// MakeEach(&EachString)
	MakeEach(&EachStringInt)
	MakeEachP(&EachP)
}

// Each func(func(A, B), []A)
// Applies the given iterator function to each element of a collection (slice or map).
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Iterator functions accept a value, and the index or key is an optional argument.
// Note: each does not return a value, you may want un.Map
// var Each func(func(value, i interface{}), interface{})
var Each func(fn interface{}, slice_or_map interface{})

// EachP Parallel Each
// *Concurrently* applies the given iterator function to each element of a collection (slice or map).
var EachP func(fn interface{}, slice_or_map interface{})

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

// MakeEachP implements a typed Parallel-Each function in the form EachP func(func(A, B), []A)
func MakeEachP(fn interface{}) {
	Maker(fn, eachP)
}

func each(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	if col.Kind() == reflect.Map {
		eachMap(fn, col)
	}

	if col.Kind() == reflect.Slice {
		eachSlice(fn, col)
	}

	return nil
}

func eachSlice(fn, s reflect.Value) {
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		eachCall(fn, v, reflect.ValueOf(i))
	}
}

func eachMap(fn, m reflect.Value) {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		eachCall(fn, v, k)
	}
}

func eachCall(fn, v, i reflect.Value) {
	args := []reflect.Value{v}
	if in := fn.Type().NumIn(); in == 2 {
		args = append(args, i)
	}
	fn.Call(args)
}

func eachP(values []reflect.Value) []reflect.Value {

	fn, col := extractArgs(values)

	if col.Kind() == reflect.Map {
		eachMapP(fn, col)
	}

	if col.Kind() == reflect.Slice {
		eachSliceP(fn, col)
	}

	return nil
}

func eachSliceP(fn, s reflect.Value) {
	var done sync.WaitGroup
	for i := 0; i < s.Len(); i++ {
		v := s.Index(i)
		done.Add(1)
		go func() {
			eachCall(fn, v, reflect.ValueOf(i))
			done.Done()
		}()
	}
	done.Wait()
}

func eachMapP(fn, m reflect.Value) {
	var done sync.WaitGroup
	keys := m.MapKeys()
	done.Add(len(keys))

	for _, k := range keys {
		v := m.MapIndex(k)
		go func(fn, v, k reflect.Value) {
			eachCall(fn, v, k)
			done.Done()
		}(fn, v, k)
	}
	done.Wait()
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
