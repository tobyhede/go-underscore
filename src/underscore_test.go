package __

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var SLICE = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var MAP = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

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

func TestMapWithInterface(t *testing.T) {
	fn := func(s interface{}) interface{} {
		return s.(string) + "!"
	}
	m := Map(ToI(SLICE), fn)
	if m[0] != "a!" {
		t.Error("First element should == a!")
	}
}

func TestMapWithMaps(t *testing.T) {
	fn := func(s string, i int) string {
		n := strconv.Itoa(i)
		return s + n
	}

	m := collectMap(MAP, fn)
	fmt.Println(m)
	if m[0] != "a!" {
		t.Error("First element should == a!")
	}
}

func TestStringToBoolMap(t *testing.T) {
	b := StringToBoolMap(SLICE, func(s string) bool {
		return s == "z"
	})

	if b[0] {
		t.Error("First element should == false")
	}

	if !b[len(b)-1] {
		t.Error("Last element should == true")
	}
}

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i interface{}) bool {
		return (i.(int) % 2) == 1
	}

	odd, even := Partition(slice, fn)

	if odd[0] != 1 {
		t.Error("First element should == 1")
	}

	if even[0] != 2 {
		t.Error("First element should == 2")
	}
}

func TestPartitionInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i int) bool {
		return (i % 2) == 1
	}

	odd, even := PartitionInt(slice, fn)

	if odd[0] != 1 {
		t.Error("First element should == 1")
	}

	if even[0] != 2 {
		t.Error("First element should == 2")
	}
}

func TestReduce(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(m, e int) int {
		return m + e
	}

	result := ReduceInt(slice, fn, 0)

	if result != 55 {
		t.Error("Expected 55 Received: ", result)
	}
}

func TestReduceR(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}

	fn := func(m, e int) int {
		return m - e
	}
	result := ReduceRInt(slice, fn, 15)

	if result != 0 {
		t.Error("Expected 0 Received: ", result)
	}
}

func contains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}

func with_types(count int) {
	in := "z"
	start := time.Now()

	for i := 0; i < count; i++ {
		contains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("Typed Contains: ", elapsed)
}

func with_interface(count int) {
	in := "z"
	start := time.Now()

	for i := 0; i < count; i++ {
		Contains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("Interface Contains: ", elapsed)
}

func with_magic(count int) {
	in := "z"
	start := time.Now()

	for i := 0; i < count; i++ {
		StringContains(SLICE, in)
	}

	elapsed := time.Since(start)
	fmt.Println("StringContains: ", elapsed)
}

func TestBench(t *testing.T) {
	count := 10000
	with_types(count)
	with_interface(count)
	with_magic(count)
}
