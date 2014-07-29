package un

import (
	"bytes"
	"strconv"
	"testing"
)

func init() {
	// suite("Each")
}

func TestEach(t *testing.T) {
	title("Each with []interface{}")
	var buffer bytes.Buffer

	fn := func(s, i interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(fn, SLICE_STRING)

	expect := "abcdefghijklmnopqrstuvwxyz"

	equals(t, expect, buffer.String())
}

func TestEachWithMap(t *testing.T) {
	title("Each with map interface{}")
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
	title("Each with map interface{} and value iterator")

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
	title("Each with map[string]int")
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
	title("Each with []int")
	var receive int

	fn := func(v, i int) {
		receive += v
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
