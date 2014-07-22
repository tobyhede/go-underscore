package un

import (
	"reflect"
)

func init() {
}

func Maker(wrapper interface{}, fn func(args []reflect.Value) (results []reflect.Value)) {
	wrapperFn := reflect.ValueOf(wrapper).Elem()
	v := reflect.MakeFunc(wrapperFn.Type(), fn)
	wrapperFn.Set(v)
}

func ToI(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("ToInterface expects a slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func wrap(v reflect.Value) []reflect.Value {
	return []reflect.Value{v}
}

func iToValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		return reflect.ValueOf(v.Interface())
	}
	return v
}
