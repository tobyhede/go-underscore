package __

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestEach(t *testing.T) {

// 	s := []int{1}

// 	Each(s, func(e interface{}, i int) {
// 		c := e.(int)
// 		c++
// 		fmt.Println(c)
// 		fmt.Println(e)
// 	})

// 	fmt.Println(s)
// }

// func TestContains(t *testing.T) {
// 	in := "d"
// 	out := "z"
// 	s := []string{"a", "b", "c", "d"}

// 	if b := Contains(s, in); b != true {
// 		t.Error("Slice should contain obj")
// 	}

// 	if b := Contains(s, out); b == true {
// 		t.Error("Slice should not contain obj")
// 	}
// }

// func TestToInterfaceString(t *testing.T) {
// 	s := []string{"a", "b", "c"}
// 	i := ToInterface(s)

// 	if len(s) != len(i) {
// 		t.Error("[]interface{} should be the same length as slice")
// 	}
// }


// func TestToInterfacePanic(t *testing.T) {
// 	defer func() {
// 		if r := recover(); r == nil {
// 			t.Error("ToInterface should panic with non-slice args")
// 		}
// 	}()

// 	ToInterface(99)
// }
//
// func TestToInterfaceInt(t *testing.T) {
// 	s := []int{1, 1, 3, 5, 8, 13}
// 	i := ToInterface(s)

// 	if len(s) != len(i) {
// 		t.Error("[]interface{} should be the same length as slice")
// 	}
// }


func TestReflect(t *testing.T) {
	slice := []int{1, 1, 3, 5, 8, 13}
	s := reflect.ValueOf(&slice)
	// p := reflect.Pointer(slice)
	fmt.Println(slice)
	fmt.Println(s)
	// v := s.Index(0).Interface()
	// fmt.Println(&v)



}

