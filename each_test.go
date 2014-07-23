package un

import (
	"bytes"
	"testing"
)

func TestEach(t *testing.T) {
	var buffer bytes.Buffer

	fn := func(s interface{}) {
		buffer.WriteString(s.(string))
	}

	Each(SLICE_STRING, fn)

	expect := "abcdefghijklmnopqrstuvwxyz"

	if receive := buffer.String(); expect != receive {
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

	if receive := buffer.String(); expect != receive {
		t.Errorf("[TestPartition] Expected %v; Received %v", expect, receive)
	}
}
