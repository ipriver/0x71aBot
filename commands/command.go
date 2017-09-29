package commands

//package for Commands which are used by Bots
import (
	"time"
)

var GlobalCD int = 5000

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

func (c *Command) Constructor(name string, mes string, privmsg bool) {
	c.Name = name
	c.Message = mes
	c.PrivMSG = privmsg
	c.OnCooldown = false
	c.CDtime = GlobalCD
}
