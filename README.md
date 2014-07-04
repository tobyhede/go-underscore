__ Underscore.go __
==========================================

Like underscore.js, but for Go
------------------------------------------


Contains returns true if an object is in a slice.
```
  o := "a"
  s := []string{"a", "b", "c"}

  b := __.Contains(s, o)
  fmt.Println(b) //true
```


ToInterface converts a slice of arbitrary type []T into []interface{}

```
  s := []int{1, 1, 3, 5, 8, 13}
  i := __.ToInterface(s)
```