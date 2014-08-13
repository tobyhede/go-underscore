package un

import (
	"fmt"
	"reflect"
	"time"
)

func init() {
	MakeMap(&Map)
	MakeMap(&MapString)
	MakeMap(&MapInt)
	MakePMap(&MapPString)
	// MakeMap(&MapStringToBool)
}

// Map func(func(A) C, []A) []C
// Applies the given iterator function to each element of a collection (slice or map) and returns a new slice.
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Iterator functions accept a value, and the index or key is an optional argument.
var Map func(interface{}, interface{}) []interface{}

// Applies the given iterator function to each element of a []string and returns a new []string of the computed results.
var MapString func(func(string) string, []string) []string

// Applies the given iterator function to each element of a []int and returns a new []int of the computed results.
var MapInt func(func(int) int, []int) []int

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

func mapWorker(fn, jobs, results reflect.Value) {
	for {
		v, ok := jobs.Recv()
		if !ok {
			break
		}
		r := fn.Call([]reflect.Value{v})
		results.Send(r[0])
	}
}

func mapPImpl(values []reflect.Value) []reflect.Value {
	fn, col := extractArgs(values)

	workers := workers
	if len(values) == 3 {
		if l := values[2].Len(); l == 1 {
			workers = int(values[2].Index(0).Int())
		}
	}

	t := col.Type().Elem()
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

	job.Close()

	for i := 0; i < col.Len(); i++ {
		v, ok := res.Recv()
		if !ok {
			break
		}
		ret.Index(i).Set(v)
	}

	return []reflect.Value{ret}
}

func mapPSlice(job, col reflect.Value) {
	for i := 0; i < col.Len(); i++ {
		e := col.Index(i)
		job.Send(e)
	}
}

func mapPMap(job, col reflect.Value) {
	for _, k := range col.MapKeys() {
		v := col.MapIndex(k)
		job.Send(v)
	}
}

func refWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func refMapMap(m map[string]int, fn func(string, int) string) []string {
	ret := make([]string, 0, len(m))

	for k, v := range m {
		ret = append(ret, fn(k, v))
	}
	return ret
}
