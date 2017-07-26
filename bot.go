package main

import (
	"config"
	"fmt"
	"monitor"
	"net"
	_ "reflect"
	"strconv"
	"strings"
)

var (
	conf *config.Config
	err  error
)

func main() {
	conf, err = config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
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

func CreateConnection(conf *config.Config) (net.Conn, error) {
	addr := strings.Join([]string{conf.HostAddr, strconv.Itoa(conf.Port)}, ":")
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("connection error occured %v", err)
	}
	conn.Write([]byte("PASS " + conf.LogOath + "\r\n"))
	conn.Write([]byte("NICK " + conf.LoginName + "\r\n"))
	conn.Write([]byte("JOIN #" + conf.Channel + " \r\n"))
	return conn, err
}
