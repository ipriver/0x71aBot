package main

import (
	"./config"
	"./monitor"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	_ "reflect"
	"strconv"
	"strings"
)

func Bot(conf *config.UserConfig) {
	conn, err := CreateConnection(conf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		fmt.Println("Connection closed")
		conn.Close()
	}()
	monitor.MonitorChannel(conn, conf)
}

func CreateConnection(conf *config.UserConfig) (net.Conn, error) {
	addr := strings.Join([]string{conf.HostAddr, strconv.Itoa(conf.Port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("connection error occured %v", err)
	}
	conn.Write([]byte("PASS " + conf.LogOath + "\r\n"))
	conn.Write([]byte("NICK " + conf.LoginBotName + "\r\n"))
	conn.Write([]byte("JOIN #" + conf.Channel + " \r\n"))
	return conn, err
}

func RunBot(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		userDecoder := json.NewDecoder(req.Body)
		defer req.Body.Close()

		userConfig, err := config.LoadUserConfig(userDecoder)
		if err != nil {
			fmt.Println(err)
			return
		}
		go Bot(userConfig)
		rw.WriteHeader(200)
	default:
		rw.WriteHeader(404)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RunBot)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
