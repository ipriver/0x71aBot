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
		err := AnswerToTwitch(strings.Split(bStri, " "), conn)
		if err != nil {
			fmt.Println("PING-PONG answer error occured")
		}
		d := strings.Split(bStri, ":")
		mes := d[len(d)-1]
		LookForCommands(mes, conn)
	}

}

func AnswerToTwitch(mesString []string, conn net.Conn) error {
	if mesString[0] == "PING" {
		answer := "PONG " + mesString[1] + "\r\n"
		conn.Write([]byte(answer))
		fmt.Println("We answered " + answer)
	}
	return nil
}
