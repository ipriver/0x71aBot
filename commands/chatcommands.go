package commands

import (
	"time"
)

type ChatCommand struct {
	Command
	message    string
	privMSG    bool
	onCooldown bool
	cdTime     int
}

type ChatCommander interface {
	SetCooldown()
	IsOnCooldown() bool
	GetMessage() string
	Call()
}

func (c *ChatCommand) setCooldown() {
	c.onCooldown = true
	go func() {
		time.Sleep(time.Duration(c.cdTime) * time.Millisecond)
		c.onCooldown = false
	}()
}

func (c *ChatCommand) IsOnCooldown() bool {
	return c.onCooldown
}

func (c *ChatCommand) IsPrivate() bool {
	return c.privMSG
}

func (c *ChatCommand) GetMessage() string {
	return c.message
}

func (c *ChatCommand) Constructor(name string, f interface{}, mes string, privmsg bool, cd int) {
	c.Command.Constructor(name, f)
	c.message = mes
	c.privMSG = privmsg
	c.onCooldown = false
	c.cdTime = cd
}

func (c *ChatCommand) Call() {
	if c.onCooldown == true {
		return
	}
	c.setCooldown()
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
