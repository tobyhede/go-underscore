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

// func TestPartitionWithMap(t *testing.T) {
// 	var buffer bytes.Buffer

// 	fn := func(v, k interface{}) {
// 		buffer.WriteString(k.(string))
// 		buffer.WriteString(strconv.Itoa(v.(int)))
// 	}

// 	Partition(fn, MAP_STRING_TO_INT)

// 	expect := "abcdefghijklmnopqrstuvwxyz1234567891011121314151617181920212223242526"
// 	receive := buffer.String()

// 	equals(t, len(expect), len(receive))
// }
