package commands

import (
	"time"
)

type ChatCommand struct {
	Command
	Message    string
	PrivMSG    bool
	OnCooldown bool
	CdTime     int
}

func (c *ChatCommand) SetCooldown() {
	c.OnCooldown = true
	go func() {
		time.Sleep(time.Duration(c.CdTime) * time.Millisecond)
		c.OnCooldown = false
	}()
}

func NewChatCommand(name string, f interface{}, mes string, privmsg bool, cd int) *ChatCommand {
	c := new(ChatCommand)
	c.Command = *NewCommand(name, f)
	c.Message = mes
	c.PrivMSG = privmsg
	c.OnCooldown = false
	c.CdTime = cd
	return c
}

func (c *ChatCommand) Call() {
	if c.OnCooldown == true {
		return
	}
	c.SetCooldown()
	c.Command.Call()
}

/*func FindChatCommand(name string) *ChatCommand {
	for _, obj := range consoleCmdList {
		if obj.GetName() == name {
			return obj
		}
	}
	return nil
}*/
