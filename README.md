Underscore.go
==========================================

Move Fast; Optimize Late
------------------------------------------

A useful collection of Go utilities. Designed for programmer happiness.

TL;DR Sort-of like underscore.js, but for Go

API Documention
------------------------------------------
[GoDoc is WorkInProgress](https://godoc.org/github.com/tobyhede/go-underscore)


:warning: Warning
------------------------------------------
This package is in heavy flux at the moment as I work to incorporate feedback from various sources.


:squirrel: Todo
------------------------------------------

- [ ] godoc
- [ ] contains
- [ ] indexOf
- [ ] worker pools
- [x] parallel each
- [x] parallel map with worker pool
- [x] refactor to make functions first parameter (eg Each func(func(A), []A))
- [x] handle maps & slices
- [x] all
- [x] any
- [x] none



------------------------------------------
* [Notes on Typed Functions](#typed)
* [Any](#any)
* [Each](#each)
* [Every](#every)
* [Map](#map)
* [None](#none)


### <a name="typed"></a>Typed Functions ###
---------------------------------------------------------------------------


### <a name="any"></a>Any ###
---------------------------------------------------------------------------


### <a name="each"></a>Each ###
---------------------------------------------------------------------------

Each func(func(A int), []A)
Each func(func(A B), []A)


Applies the given iterator function to each element of a collection (slice or map).

If the collection is a Slice, the iterator function arguments are *value, index*

If the collection is a Map, the iterator function arguments are *value, key*

EachP is a Parallel implementation of Each and *concurrently* applies the given iterator function to each element of a collection (slice or map).


``` go
  // var Each func(func(value interface{}, i interface{}), interface{})

  var buffer bytes.Buffer

  fn := func(s, i interface{}) {
    buffer.WriteString(s.(string))
  }

  s := []string{"a", "b", "c", "d", "e"}
  Each(fn, s)

  expect := "abcde"

  e := un.Each(fn, s)

  fmt.Printf("%#v\n", e) //"abcde"
```

Typed Each can be defined using a function type and the *MakeEach* helper.

Using a Typed Slice

``` go
  var EachInt func(func(value, i int), []int)
  MakeEach(&EachInt)

  var sum int

  fn := func(v, i int) {
    sum += v
  }

  i := []int{1, 2, 3, 4, 5}
  EachInt(fn, i)

  fmt.Printf("%#v\n", sum) //15
```

Using a Typed Map
``` go
  var EachStringInt func(func(key string, value int), map[string]int)
  var sum int

  fn := func(v int, k string) {
    sum += v
  }

  m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
  EachStringInt(fn, m)

  fmt.Printf("%#v\n", sum) //15
```

Of note is the ability to close over variables within the calling scope.


### <a name="every"></a>Every ###
---------------------------------------------------------------------------




### Map ###
---------------------------------------------------------------------------

``` go
Map func([]A, func(A) B) []B
```

Applies the given function to each element of a slice, returning a slice of results

The base Map function accepts interface{} types and returns []interface{}

``` go
  // Map func(interface{}, func(interface{}) interface{}) []interface{}

  s := []string{"a", "b", "c", "d"}

  fn := func(s interface{}) interface{} {
    return s.(string) + "!"
  }

  m := un.Map(ToI(s), fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
```

Typed Maps can be defined using a function type and the *MakeMap* helper.

``` go
  Map func([]A, func(A) B) []B

  var SMap func([]string, func(string) string) []string
  un.MakeMap(&SMap)

  m := un.SMap(s, fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
```

Of note is the return value of Map is a slice of the return type of the applied function.


### Partition ###
---------------------------------------------------------------------------

Partition func([]A, func(A) bool) ([]A []A)

Partition splits a slice or map based on the evaluation of the supplied function

The base Partition function accepts interface{} types and returns []interface{}


``` go
  // Partition func(interface{}, func(interface{}) bool) ([]interface{}, []interface{})

  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  fn := func(i interface{}) bool {
    return (i.(int) % 2) == 1
  }

  odd, even := un.Partition(s, fn)

  fmt.Println(odd)  //[1, 3, 5, 7, 9]
  fmt.Println(even) //[2, 4, 6, 8, 10]
```

Typed Partitions can be defined using a function type and the *MakePartition* helper.

``` go
  // Partition func([]A, func(A) bool) ([]A []A)

  var IPartition func([]int, func(int) bool) ([]int, []int)

  un.MakePartition(&IPartition)

  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  fn := func(i int) bool {
    return (i % 2) == 1
  }

  odd, even := un.IPartition(s, fn)

  fmt.Println(odd)  //[1, 3, 5, 7, 9]
  fmt.Println(even) //[2, 4, 6, 8, 10]
```


Contains returns true if an object is in a slice.

``` go
  o := "a"
  s := []string{"a", "b", "c"}

  b := un.Contains(s, o)
  fmt.Println(b) //true
```


ToI converts a slice of arbitrary type []T into a slice of []interfaces{}

``` go
  s := []int{1, 1, 3, 5, 8, 13}
  i := un.ToI(s)
```




Notes
------------------------------------------

I am aware that the whole idea is not particularly very TheGoWay™, but it is useful as a learning exercise, and it is useful for moving fast and optimising later.
