package __

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println()
	MakeContains(&Contains)
	MakeContains(&StringContains)
	// MakeMap(&Map)
	MakeMap(&StringMap)
	MakeMap(&StringToBoolMap)

	MakePartition(&Partition)
	MakePartition(&PartitionInt)

	MakeReduce(&ReduceInt)

	MakeReduceR(&ReduceRInt)
}

var Contains func(interface{}, interface{}) bool

var StringContains func([]string, string) bool

var Map func(interface{}, func(interface{}) interface{}) interface{}

var StringMap func([]string, func(string) string) []string

var StringToBoolMap func([]string, func(string) bool) []bool

var Partition func(interface{}, func(interface{}) bool) ([]interface{}, []interface{})

var PartitionInt func([]int, func(int) bool) ([]int, []int)

var PartitionString func([]string, func(string) bool) ([]string, []string)

var ReduceInt func([]int, func(int, int) int, int) int

var ReduceRInt func([]int, func(int, int) int, int) int

func Maker(wrapper interface{}, fn func(args []reflect.Value) (results []reflect.Value)) {
	wrapperFn := reflect.ValueOf(wrapper).Elem()
	v := reflect.MakeFunc(wrapperFn.Type(), fn)
	wrapperFn.Set(v)
}

func MakeContains(fn interface{}) {
	Maker(fn, _contains)
}

func MakeMap(fn interface{}) {
	Maker(fn, _map)
}

func MakePartition(fn interface{}) {
	Maker(fn, _partition)
}

func MakeReduce(fn interface{}) {
	Maker(fn, _reduce)
}

func MakeReduceR(fn interface{}) {
	Maker(fn, _reduceR)
}

func _contains(values []reflect.Value) []reflect.Value {

	v := interfaceToValue(values[0])
	o := values[1].Interface()

	for i := 0; i < v.Len(); i++ {
		e := v.Index(i).Interface()
		if e == o {
			return wrap(reflect.ValueOf(true))
		}
	}
	return wrap(reflect.ValueOf(false))
}

func _map(values []reflect.Value) []reflect.Value {

	v := interfaceToValue(values[0])
	fn := values[1]

	outType := reflect.SliceOf(fn.Type().Out(0))
	ret := reflect.MakeSlice(outType, v.Len(), v.Len())

	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		r := fn.Call([]reflect.Value{e})
		ret.Index(i).Set(r[0])
	}
	return wrap(ret)
}

func _partition(values []reflect.Value) []reflect.Value {
	slice := interfaceToValue(values[0])
	fn := values[1]

	var t, f reflect.Value

	if values[0].Kind() == reflect.Interface {
		t = reflect.ValueOf(make([]interface{},0))
		f = reflect.ValueOf(make([]interface{},0))
	} else {
		t = reflect.MakeSlice(slice.Type(), 0, 0)
		f = reflect.MakeSlice(slice.Type(), 0, 0)
	}

	for i := 0; i < slice.Len(); i++ {
		e := slice.Index(i)
		r := fn.Call([]reflect.Value{e})
		if r[0].Bool() {
			t = reflect.Append(t, e)
		} else {
			f = reflect.Append(f, e)
		}
	}
	return []reflect.Value{t, f}
}


func _reduce(values []reflect.Value) []reflect.Value {
	slice := values[0]
	fn := values[1]
	ret := values[2]

	for i := 0; i < slice.Len(); i++ {
		e := slice.Index(i)
		r := fn.Call([]reflect.Value{ret, e})
		ret = r[0]
	}

	return wrap(ret)
}

func _reduceR(values []reflect.Value) []reflect.Value {
	slice := values[0]
	fn := values[1]
	ret := values[2]

	for i := slice.Len()-1; i >= 0; i-- {
		e := slice.Index(i)
		r := fn.Call([]reflect.Value{ret, e})
		ret = r[0]
	}

	return wrap(ret)
}


func wrap(v reflect.Value) []reflect.Value {
	return []reflect.Value{v}
}

func interfaceToValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		return reflect.ValueOf(v.Interface())
	}
	return v
}

func reduce(slice []int, fn func(int, int) int, initial int) int {
	ret := initial

	for i := 0; i < len(slice); i++ {
		e := slice[i]
		ret = fn(ret, e)
	}

	return ret
}

func partition (slice []int, fn func(int) bool) ([]int, []int) {
	a := []int{}
	b := []int{}

	for i := 0; i < len(slice); i++ {
		e := slice[i]
		if fn(e) {
			a = append(a, e)
		} else {
			b = append(b, e)
		}
	}

	return a, b
}

