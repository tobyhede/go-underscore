package un

import (
	"runtime"
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

func TestMakePMap(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fn := func(s string) string {
		return s + "!"
	}

	receive := MapPString(fn, SLICE_STRING)

	// expect := "a!"
	display(receive)
	// equals(t, expect, receive[0])
}

func TestRefPSliceMap(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fn := func(s string) string {
		return s + "!"
	}

	receive := refPSliceMap(fn, SLICE_STRING)

	// expect := "a!"
	display(receive)
	// equals(t, expect, receive[0])
}
