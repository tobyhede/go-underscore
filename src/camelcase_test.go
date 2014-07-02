package camelcase

import (
	"testing"
)

func TestContains(t *testing.T) {
	o := "a"
	s := []string{"a", "b", "c"}

	if b := Contains(s, o); b != true {
		t.Error("Slice should contain obj")
	}
}

func TestToInterfaceString(t *testing.T) {
	s := []string{"a", "b", "c"}
	i := ToInterface(s)

	if len(s) != len(i) {
		t.Error("[]interface{} should be the same length as slice")
	}
}

func TestToInterfaceInt(t *testing.T) {
	s := []int{1, 1, 3, 5, 8, 13}
	i := ToInterface(s)

	if len(s) != len(i) {
		t.Error("[]interface{} should be the same length as slice")
	}
}

func TestToInterfacePanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("ToInterface should panic with non-slice args")
		}
	}()

	ToInterface(99)
}
