package un

import "reflect"

func init() {
	MakeMap(&Map)
	MakeMap(&MapString)
	MakeMap(&MapInt)
	MakePMap(&MapP)
	MakePMap(&MapPString)
	// MakeMap(&MapStringToBool)
}

// Map func(func(A) C, []A) []C
// Applies the given iterator function to each element of a collection (slice or map) and returns a new slice of the computed results.
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Iterator functions accept a value, and the index or key is an optional argument.
var Map func(interface{}, interface{}) []interface{}

// Applies the given iterator function to each element of a []string and returns a new []string of the computed results.
var MapString func(func(string) string, []string) []string

// Applies the given iterator function to each element of a []int and returns a new []int of the computed results.
var MapInt func(func(int) int, []int) []int

// Applies the given iterator function to each element of a collection (slice or map) and returns a new slice of the computed results.
var MapP func(interface{}, interface{}, ...int) []interface{}

// Applies the given iterator function to each element of a []string and returns a new []string of the computed results.
// <p>Uses a Worker Pool using either the global worker value (un.SetWorker) or as an optional parameter</p>
// <p>MapPString(fn, col, n)</p>
var MapPString func(func(string) string, []string, ...int) []string

// var MapStringToBool func([]string, func(string) bool) []bool

// MakeMap implements a typed Map function in the form Map func(func(A) C, []A) []C
func MakeMap(fn interface{}) {
	Maker(fn, mapImpl)
}

// MakePMap implements a typed Parallel Map function in the form Map func(func(A) C, []A) []C
func MakePMap(fn interface{}) {
	Maker(fn, mapPImpl)
}

func mapImpl(values []reflect.Value) []reflect.Value {

	fn, col := extractArgs(values)

	ret := makeSlice(fn, col.Len())

	// if list.Kind() == reflect.Map {
	// 	ret = everyMap(fn, list)
	// }

	if col.Kind() == reflect.Slice {
		ret = mapSlice(fn, col)
	}

	return []reflect.Value{ret}
}

func mapSlice(fn, col reflect.Value) reflect.Value {
	ret := makeSlice(fn, col.Len())

	for i := 0; i < col.Len(); i++ {
		e := col.Index(i)
		r := fn.Call([]reflect.Value{e})
		ret.Index(i).Set(r[0])
	}
	return ret
}

func mapWorker(fn reflect.Value, job chan []reflect.Value, res reflect.Value) {
	for {
		v, ok := <-job
		if !ok {
			break
		}
		if len(v) > 0 {
			r := fn.Call(v)
			res.Send(r[0])
		}
	}
}

func mapPImpl(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	workers := 1 //workers
	if len(values) == 3 {
		if l := values[2].Len(); l == 1 {
			workers = int(values[2].Index(0).Int())
		}
	}

	t := fn.Type().Out(0)
	job, res := makeWorkerChans(t)

	ret := makeSlice(fn, col.Len())

	for i := 1; i <= workers; i++ {
		go mapWorker(fn, job, res)
	}

	if col.Kind() == reflect.Slice {
		mapPSlice(job, col)
	}

	if col.Kind() == reflect.Map {
		mapPMap(job, col)
	}

	close(job)

	for i := 0; i < col.Len(); i++ {
		v, ok := res.Recv()
		if !ok {
			break
		}
		ret.Index(i).Set(v)
	}

	return []reflect.Value{ret}
}

func mapPSlice(job chan []reflect.Value, col reflect.Value) {
	for i := 0; i < col.Len(); i++ {
		e := col.Index(i)
		job <- []reflect.Value{e}
	}
}

func mapPMap(job chan []reflect.Value, col reflect.Value) {
	for _, k := range col.MapKeys() {
		v := col.MapIndex(k)
		job <- []reflect.Value{v, k}
	}
}

func refMapMap(m map[string]int, fn func(string, int) string) []string {
	ret := make([]string, 0, len(m))

	for k, v := range m {
		ret = append(ret, fn(k, v))
	}
	return ret
}
