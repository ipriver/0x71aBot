package monitor

import (
	"fmt"
	"net"
)

func SendToChannel(command string, conn net.Conn) {
	//PRIVMSG # + conf.Channel
	conn.Write([]byte("PRIVMSG #" + "ipriver" + " : " + command + "\r\n"))
}

func LookForCommands(message string, conn net.Conn) {
	command := ""

	switch message {
	case "!bot\r\n":
		command = "Hello, i'm your bot :)"

	case "!hi\r\n":
		command = "@ipriver hi"
	}
	if command != "" {
		SendToChannel(command, conn)
	}

}
