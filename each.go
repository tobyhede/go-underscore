package un

import (
	"reflect"
)

func init() {
	MakeEach(&Each)
	MakeEach(&EachInt)
}

/**
	Each func([]A, func(A))
**/

var Each func(interface{}, func(interface{}))

var EachInt func([]int, func(int))

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

/**
	Reference Each Implementations
**/
func RefEach(slice []string, fn func(string)) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i])
	}
}
