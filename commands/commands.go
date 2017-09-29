package commands

import ()

type Command struct {
	Name      string
	innerFunc func(args ...interface{})
}

func (c *Command) Call() {
	c.innerFunc()
}

func (c *Command) Constructor(name string, f func(args ...interface{})) {
	c.Name = name
	c.innerFunc = f
}
