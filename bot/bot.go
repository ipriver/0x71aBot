package bot

import (
	"../config"
	"../monitor"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type Bot struct {
	BotUserId int
	Config    *config.UserConfig
}

func (b *Bot) startBot() {
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

func CreateBotHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		userDecoder := json.NewDecoder(req.Body)
		defer req.Body.Close()

		userConfig, err := config.LoadUserConfig(userDecoder)
		if err != nil {
			fmt.Println(err)
			return
		}
		//TODO: replace botid
		bot := Bot{0, userConfig}
		go bot.startBot()
		rw.WriteHeader(200)
	default:
		rw.WriteHeader(404)
	}
}
