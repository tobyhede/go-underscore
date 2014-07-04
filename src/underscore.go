package __

import (
	"fmt"
	"reflect"
)

// type iFn func(interface{}) interface{}


func Each(slice interface{}, fn func(el interface{}, i int)) {
	s := ToInterface(slice)
	for i, e := range s {
		fmt.Println(e)
		fmt.Println(e)
		fn(e, i)
	}
}

func Contains(slice interface{}, o interface{}) bool {
	s := ToInterface(slice)
	for _, e := range s {
		if e == o {
			return true
		}
	}
	return false
}


func ToInterface(slice interface{}) []interface{} {
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
