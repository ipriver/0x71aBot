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
	Bot_id int
	Config *config.UserConfig
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

type UserJSONdata struct {
	Bot_id  int    `json:"user_id"`
	Channel string `json:"channel"`
}

func CreateBotHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		defer req.Body.Close()
		data := UserJSONdata{}
		err := decoder.Decode(&data)

		bot := checkBotInCache(data.Bot_id)
		if bot == nil {
			fmt.Println("no cache")
		}

		userConfig, err := config.LoadUserConfig(data.Channel)
		if err != nil {
			fmt.Println(err)
			return
		}
		bot = &Bot{data.Bot_id, userConfig}
		go bot.startBot()
		rw.WriteHeader(200)
	default:
		rw.WriteHeader(404)
	}
}
