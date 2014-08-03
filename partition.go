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

type partitioner struct {
	fn  reflect.Value
	col reflect.Value
	t   reflect.Value
	f   reflect.Value
}

func partition(values []reflect.Value) []reflect.Value {

	fn := interfaceToValue(values[0])
	col := interfaceToValue(values[1])
	kind := values[1].Kind()

	p := newPartitioner(fn, col, kind)

	return p.partition()
}

func newPartitioner(fn, col reflect.Value, kind reflect.Kind) *partitioner {
	t, f := makePartitions(col, kind)
	return &partitioner{fn: fn, col: col, t: t, f: f}
}

func (p *partitioner) partition() []reflect.Value {
	if p.isSlice() {
		p.partitionSlice()
	}
	return []reflect.Value{p.t, p.f}
}

func (p *partitioner) isSlice() bool {
	return p.col.Kind() == reflect.Slice
}

func (p *partitioner) isMap() bool {
	return p.col.Kind() == reflect.Map
}

func (p *partitioner) partitionSlice() {
	for i := 0; i < p.col.Len(); i++ {
		v := p.col.Index(i)
		if ok := predicate(p.fn, v, reflect.ValueOf(i)); ok {
			p.t = reflect.Append(p.t, v)
		} else {
			p.f = reflect.Append(p.f, v)
		}
	}
}

// 	for _, k := range m.MapKeys() {
// 		v := m.MapIndex(k)
// 		partitionCall(fn, v, k)
// 	}
// }

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
