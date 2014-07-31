package un

import "testing"

func init() {
	// suite("Every")
}

func TestEverySlice(t *testing.T) {
	title("Slice")

	fn := func(s interface{}) bool {
		return true
	}

	s := ToI([]int{1, 2, 3, 4, 5})
	result := Every(fn, s)

	equals(t, true, result)
}

func TestEveryMap(t *testing.T) {
	title("Map")

	fn := func(s interface{}) bool {
		return true
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Every(fn, m)

	equals(t, true, result)
}

func TestEverySliceWithInt(t *testing.T) {

	fn := func(i int) bool {
		return i <= 5
	}

	s := []int{1, 2, 3, 4, 5}
	result := EveryInt(fn, s)

	equals(t, true, result)
}

func TestEverySliceWithString(t *testing.T) {

	fn := func(s string) bool {
		// return strings.Contains(s, "!")
		return true
	}

	s := []string{"a!", "b!", "c!", "d!", "e!"}
	result := EveryString(fn, s)

	equals(t, true, result)
}

func TestEveryMapWithInt(t *testing.T) {
	title("Map With Int")
	fn := func(i interface{}) bool {
		return i.(int) <= 5
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	result := Every(fn, m)

	equals(t, true, result)
}
