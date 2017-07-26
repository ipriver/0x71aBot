package monitor

import (
	"config"
	"fmt"
	"net"
	"strings"
)

func MonitorChannel(conn net.Conn, conf *config.Config) {
	var buff []byte = make([]byte, 1024)
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
