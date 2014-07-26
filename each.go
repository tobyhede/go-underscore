package un

import (
	"fmt"
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
var Each func(interface{}, func(interface{}))

// EachInt on a slice of ints
var EachInt func(func(int), []int)

// MakeEach implements a typed Each function in the form Each func(func(A), []A)
func MakeEach(fn interface{}) {
	Maker(fn, _each)
}

func _each(values []reflect.Value) []reflect.Value {
	v := iToValue(values[0])
	fn := values[1]

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

	l := len(slice)

	// sem := make(semaphore, l)
	// close(sem)

	for i := 0; i < l; i++ {
		s := slice[i]
		done.Add(1)
		go func(s string) {
			fn(s)
			done.Done()
		}(s)
	}

	done.Wait()
	fmt.Println("end")
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
