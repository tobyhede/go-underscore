package un

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var SLICE_STRING = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var SLICE_INT = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var MAP = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

var MAP_STRING_TO_INT = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

// Test functions From https://github.com/benbjohnson/testing
// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}

func display(i interface{}) {
	fmt.Printf("\033[32m%v\033[39m\n", i)
}

func inspect(i interface{}) {
	fmt.Printf("\033[32m%#v\033[39m\n", i)
}

func TestValueize(t *testing.T) {
	i := reflect.ValueOf(42)
	s := reflect.ValueOf("42")
	res := Valueize(i, s)
	equals(t, res, []reflect.Value{i, s})
}

func TestToI(t *testing.T) {
	i := ToI(SLICE_STRING)

	if expected, received := len(SLICE_STRING), len(i); expected != received {
		t.Errorf("[ToI] Expected %v; Received %v", expected, received)
	}

	if expected, received := SLICE_STRING[0], i[0]; expected != received {
		t.Errorf("[ToI] Expected %v; Received %v", expected, received)
	}
}
