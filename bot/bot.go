package bot

import (
	"../commands"
	"errors"
	"net"
	"time"
)

var (
	OnlineBots = make(map[string]*Bot)
)

type Bot struct {
	Id       int
	Channel  string
	Commands []*commands.ChatCommand
	UpTime   time.Time
	Conn     net.Conn
	CMchan   chan string
	Quit     chan bool
}

func Constructor(id int, channel string) *Bot {
	b := new(Bot)
	b.Id = id
	b.Channel = channel
	b.Commands = make([]*commands.ChatCommand, 0)
	b.CMchan = make(chan string)
	b.Quit = make(chan bool)
	return b
}

func (b *Bot) Start() error {
	err := b.CreateConnection()
	if err != nil {
		return errors.New("Cannot establish connection")
	}
	go b.MonitorChannel()
	OnlineBots[b.Channel] = b
	return err
}

func (b *Bot) Kill() {
	b.Conn.Close()
	delete(OnlineBots, b.Channel)
}

func (b *Bot) AddCommand(cm *commands.ChatCommand) {
	b.Commands = append(b.Commands, cm)
}
