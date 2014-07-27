package un

import (
	"bytes"
	"testing"
)


func init() {
	display("Testing Each")
}

func TestEach(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(fn, SLICE_STRING)

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}

func TestEachWithMap(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(fn, MAP_STRING_TO_INT)

	expect := "abcdefghijklmnopqrstuvwxyz"
	receive := buffer.String()
	equals(t, len(expect), len(receive))

}

func TestEachInt(t *testing.T) {
	var receive int

	fn := func(i int) {
		receive += i
	}

	EachInt(fn, SLICE_INT)

	expect := 45
	equals(t, expect, receive)
}

func TestRefEach(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s string) {
		buffer.WriteString(s)
	}

	RefEach(SLICE_STRING, fn)

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}

func TestRefPEach(t *testing.T) {
	var buffer bytes.Buffer

	ch := make(chan string)

	fn := func(s string) {
		ch <- s
	}

	go func() {
		RefPEach(SLICE_STRING, fn)
		close(ch)
	}()

	for s := range ch {
		buffer.WriteString(s)
	}

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}
