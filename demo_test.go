package un

import (
	"reflect"
	"testing"
)

func init() {
	suite("Demo")
}

var i int = 42
var s string = "42"

func TestValueOf(t *testing.T) {
	v := reflect.ValueOf(i)

	inspect(v)
	// inspect(v.Int())
	// inspect(v.Interface())
}
