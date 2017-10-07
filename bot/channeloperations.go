package bot

import (
	"../config"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func (b *Bot) MonitorChannel() {
	conn := b.Conn
	var buff []byte = make([]byte, 1024)
	for {
		rb, _ := conn.Read(buff)
		bStri := string(buff[:rb])
		fmt.Println(bStri)
		d := strings.Split(bStri, ":")
		mes := d[len(d)-1]
		fmt.Println(mes)
		if mes == "hi\r\n" {
			b.Quit <- true
		}
		select {
		case <-b.Quit:
			return
		default:
		}
	}
}

func (b *Bot) CreateConnection() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Cannot establish connection")
			fmt.Println(err)
		}
	}()
	conf := config.Config
	host := conf.HostAddr
	port := conf.Port
	oath := conf.LogOath
	login := conf.LoginBotName
	channel := b.Channel

	addr := strings.Join([]string{host, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	conn.Write([]byte("PASS " + oath + "\r\n"))
	conn.Write([]byte("NICK " + login + "\r\n"))
	conn.Write([]byte("JOIN #" + channel + " \r\n"))
	if err != nil {
		panic(err)
	}
	b.Conn = conn
	return err
}
