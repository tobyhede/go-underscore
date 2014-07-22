package un

import (
	"testing"
)

var SLICE_STRING = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

var SLICE_INT = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var MAP = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

func TestToI(t *testing.T) {
	i := ToI(SLICE_STRING)

	if expected, received := len(SLICE_STRING), len(i); expected != received {
		t.Errorf("[ToI] Expected %v; Received %v", expected, received)
	}

	if expected, received := SLICE_STRING[0], i[0]; expected != received {
		t.Errorf("[ToI] Expected %v; Received %v", expected, received)
	}
}
