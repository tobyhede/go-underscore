package __

import (
	"reflect"
)

func Contains(slice interface{}, o interface{}) bool {
	s := ToInterface(slice)
	for _, i := range s {
		if i == o {
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
