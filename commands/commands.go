package commands

import (
	"reflect"
)

type Command struct {
	Name      string
	innerFunc interface{}
}

func (c *Command) Call(args ...interface{}) {
	fv := reflect.ValueOf(c.innerFunc)
	ft := fv.Type()
	margs := ft.NumIn()
	inv := make([]reflect.Value, margs)
	for n := 0; n < margs; n++ {
		if n < len(args) {
			inv[n] = reflect.ValueOf(args[n])
		} else {
			inv[n] = reflect.Zero(ft.In(n))
		}
	}
	fv.Call(inv)
}

func (c *Command) Constructor(name string, f interface{}) {
	c.Name = name
	c.innerFunc = f
}
