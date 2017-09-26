package main

import (
	"./bot"
	"./config"
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var err error

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "start":
		Start()
	case "config":
		UpdateConfig()
	default:
		StartWebService()
	}
}

func UpdateConfig() {
	scanner := bufio.NewScanner(os.Stdin)
	newconf := config.GlobalConfig{}
	newconf.Load()

	fmt.Printf("1)%s Enter new host adress:\n", newconf.HostAddr)
	scanner.Scan()
	newconf.HostAddr = scanner.Text()

	fmt.Printf("2)%d Enter new port:\n", newconf.Port)
	scanner.Scan()
	newconf.Port, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("3)%s Enter new Login for Bot:\n", newconf.LoginBotName)
	scanner.Scan()
	newconf.LoginBotName = scanner.Text()

	fmt.Printf("4)%s Enter new OATH:\n", newconf.LogOath)
	scanner.Scan()
	newconf.LogOath = scanner.Text()

	fmt.Println(newconf)
	fmt.Println("Save new config? yes/no")
	scanner.Scan()
	switch strings.ToLower(scanner.Text()) {
	case "yes":
		fallthrough
	case "1":
		err = newconf.Save()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Saved")
	}
	os.Exit(0)

}

func StartWebService() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", bot.CreateBotHandler)
	mux.HandleFunc("/bot_info/", bot.BotInfo)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
	defer func() {
		config.Rc.Close()
		config.Db.Close()
	}()
}

func Start() {

}
