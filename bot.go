package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	HOST string = "irc.twitch.tv"
	NICK string = "ipriver"
	PORT int    = 6667
	PASS string = "oauth:i67xovmub2w6tudnuglv8dogg1rarp"
)

var (
	buff []byte = make([]byte, 1024)
)

func main() {
	addr := strings.Join([]string{HOST, strconv.Itoa(PORT)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		fmt.Println("connection error occured %v", err)
	}
	conn.Write([]byte("PASS " + PASS + "\r\n"))
	conn.Write([]byte("NICK " + NICK + "\r\n"))
	conn.Write([]byte("JOIN #ipriver \r\n"))
	for {
		rb, _ := conn.Read(buff)
		fmt.Println(string(buff[:rb]))
	}
}
