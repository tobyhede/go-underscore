package un

import (
	"testing"
)

func init() {
	display("Testing Every")
}

func TestEveryS(t *testing.T) {
	fn := func(s interface{}) bool {
		return true
	}

	s := ToI([]int{1, 2, 3, 4, 5})
	result := Every(fn, s)

	equals(t, true, result)
}

func TestEveryMap(t *testing.T) {
	fn := func(s interface{}) bool {
		return true
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Every(fn, m)

	equals(t, true, result)
}

func TestEveryMapInt(t *testing.T) {
	fn := func(i interface{}) bool {
		return i.(int) <= 5
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Every(fn, m)

	equals(t, true, result)
}
