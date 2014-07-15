__ Underscore.go __
==========================================

Like underscore.js, but for Go
------------------------------------------

And yes, I am aware that the whole idea is not particularly very TheGoWay™, but it is useful as a learning exercise, and it is useful for moving fast and optimising later.


### Map ###
---------------------------------------------------------------------------

Map accepts a slice or map and a function to produce a new collection.

The base Map function accepts interface{} types and returns []interfaces{}

'''
  var Map func(interface{}, func(interface{}) interface{}) []interface{}

  s := []string{"a", "b", "c", "d"}

  fn := func(s interface{}) interface{} {
    return s.(string) + "!"
  }

  m := __.Map(ToI(slice), fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
'''

Typed Maps can be defined using a function type and the *MakeMap* helper.

'''
  Map func([]A, func(A) B) []B

  var SMap func([]string, func(string) string) []string
  __.MakeMap(&SMap)

  m := __.SMap(s, fn)
  fmt.Println(m) //["a!", "b!", "c!", "d!"]
'''

Of note is the return value of Map is a slice of the return type of the operant function




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