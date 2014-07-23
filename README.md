Underscore.go
==========================================

Move Fast; Optimize Late
------------------------------------------

A useful collection of Go utilities. Designed for programmer happiness.

TL;DR Sort-of like underscore.js, but for Go


:warning: Alert :warning:
------------------------------------------
This package is in heavy flux at the moment as I work to incorporate feedback from various sources.



### Each ###
---------------------------------------------------------------------------

Each func([]A, func(A))

Applies the given function to each element of a slice,
```
  // Each func(interface{}, func(interface{}))

  var buffer bytes.Buffer

  s := []string{"a", "b", "c", "d"}

  fn := func(s interface{}) {
    buffer.WriteString(s.(string))
  }

  e := un.Each(s, fn)
  fmt.Println(e) //["abcde"]
```

Typed Each can be defined using a function type and the *MakeEach* helper.

```
  var sum int

  fn := func(i int) {
    sum += i
  }

  i := []int{1, 2, 3, 4, 5}
  EachInt(i, fn)

  fmt.Printf("%#v\n", sum) //15
```

Of note is the ability to close over variables within the calling scope.


### Map ###
---------------------------------------------------------------------------

Map func([]A, func(A) B) []B

Applies the given function to each element of a slice, returning a slice of results

The base Map function accepts interface{} types and returns []interface{}

```
  // Map func(interface{}, func(interface{}) interface{}) []interface{}

  s := []string{"a", "b", "c", "d"}

  fn := func(s interface{}) interface{} {
    return s.(string) + "!"
  }

  m := un.Map(ToI(s), fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
```

Typed Maps can be defined using a function type and the *MakeMap* helper.

```
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


```
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

```
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
```
  o := "a"
  s := []string{"a", "b", "c"}

  b := un.Contains(s, o)
  fmt.Println(b) //true
```


ToI converts a slice of arbitrary type []T into a slice of []interfaces{}

```
  s := []int{1, 1, 3, 5, 8, 13}
  i := un.ToI(s)
```




Notes
------------------------------------------

I am aware that the whole idea is not particularly very TheGoWay™, but it is useful as a learning exercise, and it is useful for moving fast and optimising later.
