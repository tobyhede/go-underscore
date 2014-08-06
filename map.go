package un

import (
	"reflect"
)

func init() {
	MakeMap(&Map)
	MakeMap(&MapString)
	MakeMap(&MapInt)
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

var MapInt func(func(int) int, []int) []int

// var MapStringToBool func([]string, func(string) bool) []bool

func MakeMap(fn interface{}) {
	Maker(fn, mapImpl)
}

func mapImpl(values []reflect.Value) []reflect.Value {

	fn := interfaceToValue(values[0])
	col := interfaceToValue(values[1])

	var ret reflect.Value

	retType := reflect.SliceOf(fn.Type().Out(0))
	ret = reflect.MakeSlice(retType, col.Len(), col.Len())

	// if list.Kind() == reflect.Map {
	// 	ret = everyMap(fn, list)
	// }

	if col.Kind() == reflect.Slice {
		ret = mapSlice(fn, col)
	}

	return []reflect.Value{ret}
}

func mapSlice(fn, col reflect.Value) reflect.Value {

	retType := reflect.SliceOf(fn.Type().Out(0))
	ret := reflect.MakeSlice(retType, col.Len(), col.Len())

	for i := 0; i < col.Len(); i++ {
		e := col.Index(i)
		r := fn.Call([]reflect.Value{e})
		ret.Index(i).Set(r[0])
	}
	return ret
}

/**
	Reference Map impementations
**/
func RefSliceMap(slice []string, fn func(string) string) []string {
	ret := make([]string, len(slice), len(slice))

	for i := 0; i < len(slice); i++ {
		ret[i] = fn(slice[i])
	}

	return ret
}

func RefMapMap(m map[string]int, fn func(string, int) string) []string {
	ret := make([]string, 0, len(m))

	for k, v := range m {
		ret = append(ret, fn(k, v))
	}
	return ret
}
