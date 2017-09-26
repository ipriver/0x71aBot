package bot

import (
	"../commands"
	"../config"
	"../monitor"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Bot struct {
	Bot_id       int
	Config       *config.UserConfig
	Commands     []commands.Command
	UpTime       time.Time
	InnerChannel chan interface{}
}

//function creates connection to Twitch and starts listening to the channel (runs as a goroutine)
func (b *Bot) StartBot() {
	//check in cache bot_id
	//check in db bot_id
	//if not create it in id
	//run bot
	conn, err := b.CreateConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()
	monitor.MonitorChannel(conn, b.Config)
}

//twitch tcp\ip connection
func (b *Bot) CreateConnection() (net.Conn, error) {
	addr := strings.Join([]string{b.Config.HostAddr, strconv.Itoa(b.Config.Port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("connection error occured %v", err)
	}
	conn.Write([]byte("PASS " + b.Config.LogOath + "\r\n"))
	conn.Write([]byte("NICK " + b.Config.LoginBotName + "\r\n"))
	conn.Write([]byte("JOIN #" + b.Config.Channel + " \r\n"))
	return conn, err
}
