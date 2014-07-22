package un

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println()
	MakeMap(&Map)
	MakeMap(&MapString)
	MakeMap(&MapInt)
	MakeMap(&MapStringToBool)
}

/**
	Map func([]A, func(A) B) []B
**/

var Map func(interface{}, func(interface{}) interface{}) []interface{}

var MapString func([]string, func(string) string) []string

var MapInt func([]int, func(int) int) []int

var MapStringToBool func([]string, func(string) bool) []bool


func MakeMap(fn interface{}) {
	Maker(fn, _map)
}

func _map(values []reflect.Value) []reflect.Value {

	v := iToValue(values[0])
	fn := values[1]

	var ret reflect.Value

	outT := reflect.SliceOf(fn.Type().Out(0))
	ret = reflect.MakeSlice(outT, v.Len(), v.Len())

	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		r := fn.Call([]reflect.Value{e})
		ret.Index(i).Set(r[0])
	}

	return []reflect.Value{ret}
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

