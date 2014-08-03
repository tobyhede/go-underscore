package un

import "testing"

func init() {
}

func TestPartitionWithSliceInterface(t *testing.T) {

	fn := func(s interface{}) bool {
		return true
	}

	receive, _ := Partition(fn, SLICE_STRING)
	// display(fn)
	display(receive)
	// i := 99
	// v := reflect.ValueOf(i)
	// vp := v.Convert(reflect.Value)
	// display(v.Type().Implements(reflect.Value))

	// if v, ok := v.(reflect.Value); ok {
	// display("hello")
	// }

	// expect := SLICE_STRING
	// equals(t, expect, receive)
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
