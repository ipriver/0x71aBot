package main

import (
	"config"
	"fmt"
	"net"
	_ "reflect"
	"strconv"
	"strings"
)

var (
	buff []byte = make([]byte, 1024)
	conf *config.Config
	err  error
)

func main() {
	conf, err = config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	addr := strings.Join([]string{conf.HostAddr, strconv.Itoa(conf.Port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		fmt.Println("connection error occured %v", err)
	}
	conn.Write([]byte("PASS " + conf.LogOath + "\r\n"))
	conn.Write([]byte("NICK " + conf.LoginName + "\r\n"))
	conn.Write([]byte("JOIN #" + conf.Channel + " \r\n"))
	for {
		rb, _ := conn.Read(buff)
		bStri := string(buff[:rb])
		fmt.Println(bStri)
		k := strings.Split(bStri, " ")

		if k[0] == "PING" {
			answer := "PONG " + k[1] + "\r\n"
			conn.Write([]byte(answer))
			fmt.Println("We answered " + answer)
		}
		d := strings.Split(bStri, ":")
		fmt.Println(d[len(d)-1])
		if d[len(d)-1] == "!bot"+"\r\n" {
			conn.Write([]byte("PRIVMSG #" + conf.Channel + " : Hello, i'm your bot :)\r\n"))
		}
	}
}
