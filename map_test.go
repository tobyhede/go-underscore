package un

import (
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

// func TestMapInt(t *testing.T) {
// 	fn := func(i int) int {
// 		return i + 1
// 	}

// 	m := MapInt(SLICE_INT, fn)

// 	if expect, receive := m[0], 1; expect != receive {
// 		t.Errorf("[TestMapInt] Expected %v; Recieved %v", expect, receive)
// 	}
// }

// func TestMapMapStringToBool(t *testing.T) {
// 	b := MapStringToBool(SLICE_STRING, func(s string) bool {
// 		return s == "z"
// 	})

// 	if expect := b[0]; expect {
// 		t.Errorf("[TestMapStringToBool] Expected false; Received %v", expect)
// 	}

// 	last := len(b) - 1
// 	if expect := b[last]; !expect {
// 		t.Errorf("[TestMapStringToBool] Expected true; Received %v", expect)
// 	}
// }

// // func TestRefMapMap(t *testing.T) {
// // 	fn := func(s string, i int) string {
// // 		n := strconv.Itoa(i)
// // 		return s + n
// // 	}

// // 	m := RefMapMap(MAP, fn)

// // 	if m[0] != "a!" {
// // 		t.Error("First element should == a!")
// // 	}
// // }
