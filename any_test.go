package un

import (
	"strings"
	"testing"
)

func init() {
	// suite("Any")
}

func TestAnySlice(t *testing.T) {

	fn := func(s interface{}) bool {
		return true
	}

	s := ToI([]int{1, 2, 3, 4, 5})
	result := Any(fn, s)

	equals(t, true, result)
}

func TestAnyMap(t *testing.T) {

	fn := func(s interface{}) bool {
		return true
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Any(fn, m)

	equals(t, true, result)
}

func TestAnySliceWithInt(t *testing.T) {

	fn := func(i int) bool {
		return i == 3
	}

	s := []int{1, 2, 3, 4, 5}
	result := AnyInt(fn, s)

	equals(t, true, result)
}

func TestAnySliceWithString(t *testing.T) {

	fn := func(s string) bool {
		return strings.Contains(s, "d")
	}

	s := []string{"a!", "b!", "c!", "d!", "e!"}
	result := AnyString(fn, s)

	equals(t, true, result)
}

func TestAnyMapWithInt(t *testing.T) {
	fn := func(i interface{}) bool {
		return i.(int) == 2
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Any(fn, m)

	equals(t, true, result)
}
