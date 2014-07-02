package camelcase

import (
  "reflect"
)

func Contains(s interface{}, o interface{}) bool {
	a := ToInterface(s)
	for _, x := range a {
	    if x == o {
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

