package un

import (
	"testing"
)

func TestMapString(t *testing.T) {
	fn := func(s string) string {
		return s + "!"
	}

	m := MapString(SLICE_STRING, fn)

	if expect := m[0]; expect != "a!" {
		t.Error("[TestMapString] Expected %v;", expect)
	}
}

func TestMapInt(t *testing.T) {
	fn := func(i int) int {
		return i + 1
	}

	m := MapInt(SLICE_INT, fn)

	if expect, receive := m[0], 1; expect != receive {
		t.Errorf("[TestMapInt] Expected %v; Recieved %v", expect, receive)
	}
}

func TestMapWithInterface(t *testing.T) {
	fn := func(s interface{}) interface{} {
		return s.(string) + "!"
	}
	m := Map(ToI(SLICE_STRING), fn)

	if expect, receive := m[0], "a!"; expect != receive {
		t.Errorf("[TestMapWithInterface] Expected %v; Recieved %v", expect, receive)
	}

}

func TestMapMapStringToBool(t *testing.T) {
	b := MapStringToBool(SLICE_STRING, func(s string) bool {
		return s == "z"
	})

	if expect := b[0]; expect {
		t.Error("[TestMapStringToBool] Expected false; Received %v", expect)
	}

	last := len(b)-1
	if expect := b[last]; !expect {
		t.Error("[TestMapStringToBool] Expected true; Received %v", expect)
	}
}

// func TestRefMapMap(t *testing.T) {
// 	fn := func(s string, i int) string {
// 		n := strconv.Itoa(i)
// 		return s + n
// 	}

// 	m := RefMapMap(MAP, fn)

// 	if m[0] != "a!" {
// 		t.Error("First element should == a!")
// 	}
// }