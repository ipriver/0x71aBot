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
	config := config.LoadConfig()
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
	conn.Write([]byte("JOIN #" + CHANNEL + " \r\n"))
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
			conn.Write([]byte("PRIVMSG #" + CHANNEL + " : Hello, i'm your bot :)\r\n"))
		}
	}
}
