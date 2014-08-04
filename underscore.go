package un

import (
	"reflect"
)

func init() {
}

// Maker takes a function pointer (fn) and implements it with the given reflection-based function implementation
// Internally uses reflect.MakeFunc
func Maker(fn interface{}, impl func(args []reflect.Value) (results []reflect.Value)) {
	fnV := reflect.ValueOf(fn).Elem()
	fnI := reflect.MakeFunc(fnV.Type(), impl)
	fnV.Set(fnI)
}

// ToI takes a slice and converts it to type []interface[]
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

// Valueize takes a number of arguments and returns them as []reflect.Value
func Valueize(values ...interface{}) []reflect.Value {
	ret := make([]reflect.Value, len(values))

	for i := 0; i < len(values); i++ {
		v := values[i]
		if t := reflect.TypeOf(v).String(); t == "reflect.Value" {
			ret[i] = v.(reflect.Value)
		} else {
			ret[i] = reflect.ValueOf(v)
		}
	}

	return ret
}

// InterfaceToValue converts a value of interface{} to a value of Interface()
// That is, converts to the underlying type of the reflect.Value
func interfaceToValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		return reflect.ValueOf(v.Interface())
	}
	return v
}

func predicate(fn reflect.Value, args ...reflect.Value) bool {
	in := fn.Type().NumIn()
	res := fn.Call(args[0:in])
	return res[0].Bool()
}
