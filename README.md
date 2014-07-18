__ Underscore.go __
==========================================

Like underscore.js, but for Go
------------------------------------------

And yes, I am aware that the whole idea is not particularly very TheGoWay™, but it is useful as a learning exercise, and it is useful for moving fast and optimising later.


### Map ###
---------------------------------------------------------------------------

Map func([]A, func(A) B) []B

Map accepts a slice or map and a function to produce a new collection.

The base Map function accepts interface{} types and returns []interface{}

```
  // Map func(interface{}, func(interface{}) interface{}) []interface{}

  s := []string{"a", "b", "c", "d"}

  fn := func(s interface{}) interface{} {
    return s.(string) + "!"
  }

  m := __.Map(ToI(slice), fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
```

Typed Maps can be defined using a function type and the *MakeMap* helper.

```
  Map func([]A, func(A) B) []B

  var SMap func([]string, func(string) string) []string
  __.MakeMap(&SMap)

  m := __.SMap(s, fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
```

Of note is the return value of Map is a slice of the return type of the operant function.


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

  odd, even := __.Partition(s, fn)

  fmt.Println(odd)  //[1, 3, 5, 7, 9]
  fmt.Println(even) //[2, 4, 6, 8, 10]
```

Typed Partitions can be defined using a function type and the *MakeMap* helper.

```
  // Partition func([]A, func(A) bool) ([]A []A)

  var IPartition func([]int, func(int) bool) ([]int, []int)

  __.MakePartition(&IPartition)

  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  fn := func(i int) bool {
    return (i % 2) == 1
  }

  odd, even := __.IPartition(s, fn)

  fmt.Println(odd)  //[1, 3, 5, 7, 9]
  fmt.Println(even) //[2, 4, 6, 8, 10]
```


Contains returns true if an object is in a slice.
```
  o := "a"
  s := []string{"a", "b", "c"}

  b := __.Contains(s, o)
  fmt.Println(b) //true
```


ToI converts a slice of arbitrary type []T into a slice of []interfaces{}

```
  s := []int{1, 1, 3, 5, 8, 13}
  i := __.ToI(s)
```