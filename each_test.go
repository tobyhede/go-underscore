package un

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEach(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(fn, SLICE_STRING)

	expect := "abcdefghijklmnopqrstuvwxyz"

	if receive := buffer.String(); expect != receive {
		t.Errorf("[TestPartition] Expected %v; Received %v", expect, receive)
	}
}

func TestEachInt(t *testing.T) {
	var receive int

	fn := func(i int) {
		receive += i
	}

	EachInt(fn, SLICE_INT)

	if expect := 45; expect != receive {
		t.Errorf("[TestPartition] Expected %v; Received %v", expect, receive)
	}
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

	fmt.Println(buffer.String())
	fmt.Println("-------")

	equals(t, expect, buffer.String())
}
