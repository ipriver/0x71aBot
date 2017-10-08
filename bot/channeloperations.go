package bot

import (
	"../config"
	"fmt"
	"net"
	"strconv"
	"strings"
)

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

func (b *Bot) MonitorChannel() {
	var buff []byte = make([]byte, 1024)
	for {
		ReadBuffer(b.Conn, buff, b.CMchan)

		select {
		case resp := <-b.CMchan:
			b.Conn.Write([]byte("PRIVMSG #" + b.Channel + " : " + resp))
		case <-b.Quit:
			b.Kill()
			return
		default:
		}
	}
}

func ReadBuffer(conn net.Conn, buff []byte, rcChan chan string) {
	//fmt.Println("BUFFER: ", buff)
	rb, _ := conn.Read(buff)
	bStri := string(buff[:rb])
	fmt.Println("bSTRINNG: ", bStri)

	d := strings.Split(bStri, ":")
	mes := d[len(d)-1]
	go ReadMessage(mes, rcChan)
}

func ReadMessage(mes string, rcChan chan string) {
	defer func() {
		recover()
	}()
	if true {
		rcChan <- mes
	}
	fmt.Println("MESSAGE: ", mes)
}
