package commands

import (
	"reflect"
	"testing"
)

func TestCommand(t *testing.T) {
	c := Command{}
	rc := reflect.TypeOf(c)
	if rc.Name() != "Command" {
		t.Error("incorrect struct name")
	}
}

func TestGetName(t *testing.T) {
	n := "Max"
	c := Command{name: n}
	if c.GetName() != n {
		t.Error("incorrect value")
	}
	if reflect.TypeOf(c.GetName()).String() != "string" {
		t.Error("incorrect return type")
	}
}

func TestCallWithIncorrectInterface(t *testing.T) {
	defer func() {
		//Check that we are catching panic
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	c := Command{"hi", 111}
	c.Call()
}

func TestCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	c := Command{"hi", func() {}}
	c.Call()
	k := Command{"hi", func(x int) int { return x * 20 }}
	k.Call(24)
	d := Command{"hi", func(x string) int { return len(x) }}
	d.Call("test")
}
