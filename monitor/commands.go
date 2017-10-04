package monitor

/*import (
	"../config"
	_ "fmt"
	"net"
)

func SendToChannel(command string, conn net.Conn, conf *config.UserConfig) {
	//PRIVMSG # + conf.Channel
	conn.Write([]byte("PRIVMSG #" + conf.Channel + " : " + command + "\r\n"))
}

func LookForCommands(message string, conn net.Conn, conf *config.UserConfig) {
	command := ""

	switch message {
	case "!bot\r\n":
		command = "Hello, i'm your bot :)"

	case "!hi\r\n":
		command = "@ipriver hi"
	}
	if command != "" {
		SendToChannel(command, conn, conf)
	}

}
*/
