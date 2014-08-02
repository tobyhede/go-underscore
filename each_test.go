package un

import (
	"bytes"
	"strconv"
	"testing"
)

func init() {
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

func TestEachWithIndex(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s, i interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(fn, SLICE_STRING)

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}

func TestEachWithMap(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(v, k interface{}) {
		buffer.WriteString(k.(string))
		buffer.WriteString(strconv.Itoa(v.(int)))
	}

	Each(fn, MAP_STRING_TO_INT)

	expect := "abcdefghijklmnopqrstuvwxyz1234567891011121314151617181920212223242526"
	receive := buffer.String()

	equals(t, len(expect), len(receive))
}

func TestEachWithMapValueOnly(t *testing.T) {

	var buffer bytes.Buffer

	fn := func(v interface{}) {
		buffer.WriteString(strconv.Itoa(v.(int)))
	}

	Each(fn, MAP_STRING_TO_INT)

	expect := "1234567891011121314151617181920212223242526"
	receive := buffer.String()

	equals(t, len(expect), len(receive))
}

func TestEachStringInt(t *testing.T) {
	var receive int

	fn := func(v int, k string) {
		receive += v
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	EachStringInt(fn, m)

	expect := 15
	equals(t, expect, receive)
}

func TestEachInt(t *testing.T) {
	var receive int

	fn := func(v, i int) {
		receive += v
	}

	EachInt(fn, SLICE_INT)

	expect := 45
	equals(t, expect, receive)
}

func TestEachP(t *testing.T) {
	var buffer bytes.Buffer

	ch := make(chan string)

	fn := func(s string) {
		ch <- s
	}

	go func() {
		EachP(fn, SLICE_STRING)
		close(ch)
	}()

	for s := range ch {
		buffer.WriteString(s)
	}

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}

func TestEachPMap(t *testing.T) {
	var buffer bytes.Buffer

	ch := make(chan int)

	fn := func(i int) {
		ch <- i
	}

	go func() {
		EachP(fn, MAP_STRING_TO_INT)
		close(ch)
	}()

	for i := range ch {
		buffer.WriteString(strconv.Itoa(i))
	}

	// expect := "abcdefghijklmnopqrstuvwxyz"
	expect := "1234567891011121314151617181920212223242526"
	receive := buffer.String()
	equals(t, len(expect), len(receive))
}

func TestRefEach(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s string) {
		buffer.WriteString(s)
	}

	refEach(SLICE_STRING, fn)

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
		refPEach(SLICE_STRING, fn)
		close(ch)
	}()

	for s := range ch {
		buffer.WriteString(s)
	}

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}
