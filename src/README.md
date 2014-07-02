CamelCase is like under_score, but for Go
==========================================

See what I did there?

Right now we have a generic Contains function for slices.

```
  o := "a"
  s := []string{"a", "b", "c"}

  b := camelcase.Contains(s, o)
  fmt.Println(b) //true
```