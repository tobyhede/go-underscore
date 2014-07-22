package un

import (
	"reflect"
)

func init() {
	MakePartition(&Partition)
	MakePartition(&PartitionInt)
}

/**
	Partition func([]A, func(A) bool) ([]A []A)
**/

var Partition func(interface{}, func(interface{}) bool) ([]interface{}, []interface{})

var PartitionInt func([]int, func(int) bool) ([]int, []int)

var PartitionString func([]string, func(string) bool) ([]string, []string)

func MakePartition(fn interface{}) {
	Maker(fn, _partition)
}

func _partition(values []reflect.Value) []reflect.Value {
	slice := iToValue(values[0])
	fn := values[1]

	var t, f reflect.Value

	if values[0].Kind() == reflect.Interface {
		t = reflect.ValueOf(make([]interface{}, 0))
		f = reflect.ValueOf(make([]interface{}, 0))
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
