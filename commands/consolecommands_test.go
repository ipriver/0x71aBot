package commands

import (
	"reflect"
	"testing"
)

func createConsoleObj() ConsoleCommand {
	c := ConsoleCommand{}
	c.Constructor("test", func() {}, "random descr", []string{"one", "two"})
	return c
}

func TestConsoleCommand(t *testing.T) {
	c := createConsoleObj()
	rc := reflect.TypeOf(c)
	if rc.Name() != "ConsoleCommand" {
		t.Error("incorrect struct name")
	}
}

func TestConsoleGetDescr(t *testing.T) {
	c := createConsoleObj()
	if reflect.TypeOf(c.GetDescr()).String() != "string" {
		t.Error("incorrect return type")
	}
	if c.GetDescr() != "random descr" {
		t.Error("incorrect result")
	}
}

func TestConsoleGetArguments(t *testing.T) {
	c := createConsoleObj()
	if reflect.TypeOf(c.GetArguments()).String() != "[]string" {
		t.Error("incorrect return type")
	}
	if !reflect.DeepEqual(c.GetArguments(), c.arguments) {
		t.Error("incorrect result")
	}
}

func TestConsoleCallWithIncorrectInterface(t *testing.T) {
	defer func() {
		//Check that we are catching panic
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	c := ConsoleCommand{}
	c.Constructor("rand", 111, "", []string{"", ""})
	c.Call()
}

func TestConsoleCall(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	c := createConsoleObj()
	c.Call()
}
