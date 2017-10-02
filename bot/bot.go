package bot

import (
	"../commands"
	"../config"
	/*"../monitor"
	"fmt"
	"net"
	"strconv"
	"strings"*/
	"time"
)

var OnlineBots = make(map[int]*Bot)

type Bot struct {
	id           int
	config       *config.UserConfig
	commands     []commands.ChatCommand
	upTime       time.Time
	connection   net.Conn
	innerChannel chan interface{}
}

func (b *Bot) GetId() int {
	return b.id
}

func (b *Bot) SetId(id int) {
	b.id = id
}

func (b *Bot) GetConfig() *config.UserConfig {
	return b.config
}

func (b *Bot) SetConfig(config *config.UserConfig) {
	b.config = config
}

func (b *Bot) GetUptime() time.Time {
	return b.upTime
}

func (b *Bot) Constructor(id int, channel string) {
	b.id = id
	conf := &config.UserConfig{}
	conf.SetChannel(channel)
	conf.Load()
	b.config = conf
}

//function creates connection to Twitch and starts listening to the channel (runs as a goroutine)
func (b *Bot) StartBot() {
	botId := b.GetId()
	_, ok := OnlineBots[botId]
	if ok == true {
		fmt.Println("Bot is already online")
		return
	}
	//check in cache bot_id
	//check in db bot_id
	//if not create it in id
	//run bot
	err = b.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println("Connection closed")
		b.connection.Close()
	}()
	OnlineBots[botId] = b
	monitor.MonitorChannel(conn, b.Config)
}

//twitch tcp\ip connection
func (b *Bot) CreateConnection() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Cannot establish connectoin")
			fmt.Println(err)
			return error
		}
	}()

	conf := b.GetConfig()
	host := conf.GetHost()
	port := conf.GetPort()
	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	oath := conf.GetOath()
	login := conf.GetLoginBotName()
	channel := conf.GetChannel()
	conn.Write([]byte("PASS " + oath + "\r\n"))
	conn.Write([]byte("NICK " + login + "\r\n"))
	conn.Write([]byte("JOIN #" + channel + " \r\n"))
	if err != nil {
		panic(err)
	}
	b.connection = conn
	return nil
}

func (b *Bot) AddCommand() {
	cm := commands.ChatCommand{}
	//TODO: custom commands
	cm.Constructor("exit", func() {}, "None", false, 3000)
	b.Commands = append(b.Commands, cm)
}
