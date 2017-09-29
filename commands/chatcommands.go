package commands

import (
	"time"
)

type ChatCommand struct {
	Command
	Message    string
	PrivMSG    bool
	OnCooldown bool
	CDtime     int
}

func (c *ChatCommand) Cooldown() {
	c.OnCooldown = true
	go func() {
		time.Sleep(5 * time.Millisecond)
		c.OnCooldown = false
	}()
}

func (c *ChatCommand) Constructor(name string, mes string, privmsg bool, cd int) {
	c.Name = name
	c.Message = mes
	c.PrivMSG = privmsg
	c.OnCooldown = false
	c.CDtime = cd
}

func (c *ChatCommand) Call() {
	if c.OnCooldown == true {
		return
	}
	c.Cooldown()
	c.innerFunc()
}
