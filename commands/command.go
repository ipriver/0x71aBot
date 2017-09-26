package commands

//package for Commands which are used by Bots
import (
	"time"
)

type Command struct {
	Name       string
	Message    string
	PrivMSG    bool
	OnCooldown bool
	CDtime     int
}

func (c *Command) Cooldown() {
	c.OnCooldown = true
	go func() {
		time.Sleep(5 * time.Millisecond)
		c.OnCooldown = false
	}()
}
