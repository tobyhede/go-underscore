package un

import (
	"testing"
)

func TestPartition(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i interface{}) bool {
		return (i.(int) % 2) == 1
	}

	odd, even := Partition(slice, fn)

	if expect, receive := 1, odd[0]; expect != receive {
		t.Errorf("[TestPartition] Expected %v; Received %v", expect, receive)
	}

	if expect, receive := 2, even[0]; expect != receive {
		t.Errorf("[TestPartition] Expected %v; Received %v", expect, receive)
	}
}

func TestPartitionInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fn := func(i int) bool {
		return (i % 2) == 1
	}

	odd, even := PartitionInt(slice, fn)

	if expect, receive := 1, odd[0]; expect != receive {
		t.Errorf("[TestPartitionInt] Expected %v; Received %v", expect, receive)
	}

	if expect, receive := 2, even[0]; expect != receive {
		t.Errorf("[TestPartitionInt] Expected %v; Received %v", expect, receive)
	}

}
