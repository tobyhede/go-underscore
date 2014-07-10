package __

import (
	"fmt"
	"testing"
	"time"
)

var SLICE = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// func TestEach(t *testing.T) {
// 	s := []int{1,2,3,4,5}

// 	x := 0
// 	Each(s, func(e interface{}, i int) {
// 		c := e.(int)
// 		x = x + c
// 		fmt.Println(c)
// 	})

// 	fmt.Println(x)
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



func contains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}


var COUNT = 10000

var in = "z"


func with_types() {

	start := time.Now()

	for i := 0; i < COUNT; i++ {
		contains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("Typed Contains: ", elapsed)
}

func with_interface() {
	start := time.Now()

	for i := 0; i < COUNT; i++ {
		Contains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("Interface Contains: ", elapsed)
}

func with_magic() {

	start := time.Now()

	for i := 0; i < COUNT; i++ {
		StringContains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("StringContains: ", elapsed)
}


func TestBench(t *testing.T) {
	// with_types()
	// with_interface()
	// with_magic()
}

// func TestMap(t *testing.T) {
// 	mapped := Map(SLICE, func(e interface{}) interface{} {
// 		s := e.(string)
// 		return s + "!"
// 	})
// 	fmt.Println(mapped)
// 	// if mapped[0] != "a!" {
// 	// 	t.Error("First element should be a!")
// 	// }
// 	// if mapped[len(mapped)-1] != "z!" {
// 	// 	t.Error("Last element should be z!")
// 	// }
// }


func TestContains(t *testing.T) {
	in := "a"
	out := "!"

	if b := Contains(SLICE, in); b != true {
		t.Error("Slice should contain ", in)
	}

	if b := Contains(SLICE, out); b != false {
		t.Error("Slice should not contain ", out)
	}
}

func TestStringMap(t *testing.T) {
	fn := func(s string) string {
		return s + "!"
	}

	m := StringMap(SLICE, fn)

	if m[0] != "a!" {
		t.Error("First element should == a!")
	}
}


func TestStringToBoolMap(t *testing.T) {
	b := StringToBoolMap(SLICE, func(s string) bool {
		return s == "z"
	})

	if b[0]  {
		t.Error("First element should == false")
	}

	if !b[len(b)-1]  {
		t.Error("Last element should == true")
	}
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

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i int) bool {
		return (i % 2) == 1
	}

	// odd, even := partition(slice, fn)
	odd, even := PartitionInt(slice, fn)

	fmt.Println(odd)
	fmt.Println(even)

	if odd[0] != 1  {
		t.Error("First element should == 1")
	}

	if even[0] != 2  {
		t.Error("First element should == 2")
	}
}

// func baseIterator(in []reflect.Value) []reflect.Value {
// 	return []reflect.Value{in[0]}
// }
// var intIterator func(int) int
	// handler := func(fnPtr interface{}) {
	// 	fmt.Println(fnPtr)
	// 	fn := reflect.ValueOf(fnPtr).Elem()
	// 	fmt.Println(fn)
	// 	v := reflect.MakeFunc(fn.Type(), baseIterator)
	// 	// Assign it to the value fn represents.
	// 	fn.Set(v)
	// }
	// handler(&intIterator)
	// fmt.Println("intIterator: ", intIterator(1))


	// fn := func(s string) string {
	// 	return s
	// }
	// res := StringMap(slice, fn)
	// fmt.Println(res)
	// slice := []int{1, 1, 3, 5, 8, 13}
	// s := reflect.ValueOf(slice)
	// p := reflect.Pointer(slice)
	// fmt.Println(slice)
	// fmt.Println(s.Index(1))
	// fmt.Println("slice")
	// fmt.Println("type of p:", s.Type())
	// fmt.Println("settability of p:", s.CanSet())
	// v := s.Index(0).Interface()
	// fmt.Println(&v)
