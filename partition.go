package un

import "reflect"

func init() {
	MakePartition(&Partition)
	// MakePartition(&PartitionInt)
	// MakePartition(&PartitionString)
	// MakePartition(&PartitionStringInt)
	// MakePartitionP(&PartitionP)
}

// Partition func(func(A, B), []A)
// Applies the given iterator function to partition element of a collection (slice or map).
// If the collection is a Slice, the iterator function arguments are *value, index*
// If the collection is a Map, the iterator function arguments are *value, key*
// Iterator functions accept a value, and the index or key is an optional argument.
// Note: partition does not return a value, you may want un.Map
// var Partition func(func(value, i interface{}), interface{})
var Partition func(fn interface{}, slice_or_map interface{}) ([]interface{}, []interface{})

// var Partition func(interface{}, func(interface{}) bool) ([]interface{}, []interface{})

// // PartitionP Parallel Partition
// // *Concurrently* applies the given iterator function to partition element of a collection (slice or map).
// var PartitionP func(fn interface{}, slice_or_map interface{})

// // PartitionInt
// // Applies the given iterator function to partition element of []int
// // Iterator function arguments are *value, index*
// var PartitionInt func(func(value, i int), []int)

// // PartitionStringInt
// // Applies the given iterator function to partition element of map[string]int
// // Iterator function arguments are *value, key*
// var PartitionStringInt func(func(value int, key string), map[string]int)

// MakePartition implements a typed Partition function in the form Partition func(func(A, B), []A)
func MakePartition(fn interface{}) {
	Maker(fn, partition)
}

func partition(values []reflect.Value) []reflect.Value {

	fn := interfaceToValue(values[0])
	col := interfaceToValue(values[1])
	kind := values[1].Kind()

	t, f := makePartitions(col, kind)
	// display(t)
	// display(f)
	display(kind)
	// display(fn)
	// if col.Kind() == reflect.Map {
	// 	partitionMap(fn, col)
	// }

	if col.Kind() == reflect.Slice {
		partitionSlice(fn, col)
	}

	return Valueize(t, f)
}

func partitionSlice(fn, s reflect.Value) {
	// for i := 0; i < s.Len(); i++ {
	// 	v := s.Index(i)
	// 	partitionCall(fn, v, reflect.ValueOf(i))
	// }

	// if ok := predicate(fn, v); !ok {
	// 	return false
	// }

	// for i := 0; i < slice.Len(); i++ {
	// 	e := slice.Index(i)
	// 	// r := fn.Call([]reflect.Value{e})
	// 	if r[0].Bool() {
	// 		t = reflect.Append(t, e)
	// 	} else {
	// 		f = reflect.Append(f, e)
	// 	}
	// }

}

func partitionMap(fn, m reflect.Value) {
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		partitionCall(fn, v, k)
	}
}

func partitionCall(fn, v, i reflect.Value) {
	args := Valueize(v)
	if in := fn.Type().NumIn(); in == 2 {
		args = append(args, i)
	}
	fn.Call(args)
}

func makePartitions(col reflect.Value, kind reflect.Kind) (reflect.Value, reflect.Value) {
	var t, f reflect.Value

	if kind == reflect.Interface {
		t = reflect.ValueOf(make([]interface{}, 0))
		f = reflect.ValueOf(make([]interface{}, 0))
	} else {
		t = reflect.MakeSlice(col.Type(), 0, 0)
		f = reflect.MakeSlice(col.Type(), 0, 0)
	}
	return t, f
}
