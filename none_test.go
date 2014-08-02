package un

import (
	"strings"
	"testing"
)

func init() {
}

func TestNoneSlice(t *testing.T) {

	fn := func(s interface{}) bool {
		return false
	}

	s := ToI([]int{1, 2, 3, 4, 5})
	result := None(fn, s)

	equals(t, true, result)
}

func TestNoneMap(t *testing.T) {

	fn := func(s interface{}) bool {
		return false
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := None(fn, m)

	equals(t, true, result)
}

func TestNoneSliceWithInt(t *testing.T) {

	fn := func(i int) bool {
		return i > 10
	}

	s := []int{1, 2, 3, 4, 5}
	result := NoneInt(fn, s)

	equals(t, true, result)
}

func TestNoneSliceWithString(t *testing.T) {

	fn := func(s string) bool {
		return strings.Contains(s, "z")
	}

	s := []string{"a!", "b!", "c!", "d!", "e!"}
	result := NoneString(fn, s)

	equals(t, true, result)
}

func TestNoneMapWithInt(t *testing.T) {
	fn := func(i interface{}) bool {
		return i.(int) >= 211
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := None(fn, m)

	equals(t, true, result)
}
