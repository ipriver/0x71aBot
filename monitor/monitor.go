package monitor

import (
	"../config"
	"fmt"
	"net"
	"strings"
)

func MonitorChannel(conn net.Conn, conf *config.UserConfig) {
	var buff []byte = make([]byte, 1024)
	for {
		rb, _ := conn.Read(buff)
		bStri := string(buff[:rb])
		fmt.Println(bStri)
		d := strings.Split(bStri, ":")
		mes := d[len(d)-1]
		LookForCommands(mes, conn)
	}

}
