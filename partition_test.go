package un

import "testing"

func init() {
}

func TestPartitionWithSliceInterface(t *testing.T) {

	fn := func(s interface{}) bool {
		return true
	}

	receive, _ := Partition(fn, SLICE_STRING)

	expect := SLICE_STRING
	equals(t, ToI(expect), receive)
}

func TestPartitionWithMapInterface(t *testing.T) {

	fn := func(s interface{}) bool {
		return true
	}

	receive, _ := Partition(fn, MAP_STRING_TO_INT)

	expect := MAP_STRING_TO_INT
	equals(t, len(expect), len(receive))

	display(receive)
}

func TestPartitionWithSliceInt(t *testing.T) {

	fn := func(i int) bool {
		return i < 5
	}

	under, over := Partition(fn, SLICE_INT)

	equals(t, 5, len(under))
	equals(t, 5, len(over))

	equals(t, 0, under[0])
	equals(t, 5, over[0])
}
