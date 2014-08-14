package un

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestMapWithSliceInterface(t *testing.T) {
	fn := func(s interface{}) interface{} {
		return s.(string) + "!"
	}
	receive := Map(fn, SLICE_STRING)

	expect := "a!"
	equals(t, expect, receive[0])
}

func TestMapString(t *testing.T) {
	fn := func(s string) string {
		return s + "!"
	}

	receive := MapString(fn, SLICE_STRING)

	expect := "a!"
	equals(t, expect, receive[0])
}

func TestMapInt(t *testing.T) {
	fn := func(i int) int {
		return i + 1
	}

	receive := MapInt(fn, SLICE_INT)

	expect := 1
	equals(t, expect, receive[0])
}

func TestMakePMapWithSlice(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fn := func(s string) string {
		return s + "!"
	}
	receive := MapPString(fn, SLICE_STRING)

	assert(t, strings.Contains(receive[0], "!"), "should contain !")
}

func TestMakePMapWithMap(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fn := func(v, k interface{}) interface{} {
		return k.(string) + strconv.Itoa(v.(int)) + "!"
	}

	receive := MapP(fn, MAP_STRING_TO_INT, 20)

	display(receive)
	// assert(t, strings.Contains(receive[0], "!"), "should contain !")
}
