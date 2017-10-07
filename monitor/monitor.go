package monitor

/*import (
	"../bot"
	//"../config"
	//"fmt"
	"net"
	"strings"
)

//MonitorChannel(conn net.Conn, conf *config.UserConfig)
func MonitorChannel(bot *bot.Bot) {
	conn := bot.connection
	var buff []byte = make([]byte, 1024)
	for {
		rb, _ := conn.Read(buff)
		bStri := string(buff[:rb])
		fmt.Println(bStri)
		d := strings.Split(bStri, ":")
		mes := d[len(d)-1]
		fmt.Println(mes, conn, conf)
		//LookForCommands(mes, conn, conf)
	}

}
*/
