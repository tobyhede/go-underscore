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

// Each func(func(A, B), []A)
// Applies the given iterator function to each element of a collection (slice or map).
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Iterator functions accept a value, and the index or key is an optional argument.
// Note: each does not return a value, you may want un.Map
// var Each func(func(value, i interface{}), interface{})
// var Map func(interface{}, func(interface{}) interface{}) []interface{}
var Map func(interface{}, interface{}) []interface{}

var MapString func(func(string) string, []string) []string

var MapPString func(func(string) string, []string, ...int) []string

var MapInt func(func(int) int, []int) []int

// var MapStringToBool func([]string, func(string) bool) []bool

// Each func(func(A, B), []A)
// Applies the given iterator function to each element of a collection (slice or map).
func MakeMap(fn interface{}) {
	Maker(fn, mapImpl)
}

// Each func(func(A, B), []A)
// Applies the given iterator function to each element of a collection (slice or map).
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

	for j := 0; j < col.Len(); j++ {
		e := col.Index(j)
		job.Send(e)
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

func refWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

/**
	Reference Map impementations
**/
func refPSliceMap(fn func(string) string, slice []string) []string {
	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go refWorker(w, jobs, results)
	}

	// Here we send 9 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	var is []int

	// Finally we collect all the results of the work.
	for a := 1; a <= 9; a++ {
		i := <-results
		is = append(is, i)
	}

	display(is)
	return nil
}

func refMapMap(m map[string]int, fn func(string, int) string) []string {
	ret := make([]string, 0, len(m))

	for k, v := range m {
		ret = append(ret, fn(k, v))
	}
	return ret
}

// func refPEach(slice []string, fn func(string)) {
// 	var done sync.WaitGroup

// 	for _, s := range slice {
// 		s := s
// 		done.Add(1)
// 		go func() {
// 			fn(s)
// 			done.Done()
// 		}()
// 	}

// 	done.Wait()
// }
