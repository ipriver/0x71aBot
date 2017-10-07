package bot

import (
	"fmt"
	"strings"
)

func (b *Bot) MonitorChannel() {
	conn := b.connection
	var buff []byte = make([]byte, 1024)
	for {
		rb, _ := conn.Read(buff)
		bStri := string(buff[:rb])
		fmt.Println(bStri)
		d := strings.Split(bStri, ":")
		mes := d[len(d)-1]
		fmt.Println(mes)
		//LookForCommands(mes, conn, conf)
	}
}

func (b *Bot) CreateConnection() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Cannot establish connection")
			fmt.Println(err)
		}
	}()
	conf := b.GetConfig()
	host := conf.GetHost()
	port := conf.GetPort()
	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	oath := conf.GetOath()
	login := conf.GetLoginBotName()
	channel := conf.GetChannel()
	conn.Write([]byte("PASS " + oath + "\r\n"))
	conn.Write([]byte("NICK " + login + "\r\n"))
	conn.Write([]byte("JOIN #" + channel + " \r\n"))
	if err != nil {
		panic(err)
	}
	b.connection = conn
	return nil
}
