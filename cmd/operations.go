package cmd

import (
	"../bot"
	"../commands"
	"../config"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

//TODO: stop, info
var cmdList = []string{"exit", "status", "stop", "info", "run"}

func ListenCMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case cmdList[0]:
			os.Exit(0)
		case cmdList[1]:
			GetOnlineBots()
		}
		run, _ := regexp.MatchString(cmdList[4], scanner.Text())
		if run == true {
			channel := strings.Fields(scanner.Text())
			if len(channel) > 1 {
				RunCustomBot(channel[1])
			}
		}
	}
}

func GetOnlineBots() {
	fmt.Printf("Online bots: %v\n", len(bot.OnlineBots))
	for i, b := range bot.OnlineBots {
		uptime := time.Since(b.UpTime)
		fmt.Printf("id: %d, Channel: %s, Uptime: %v\n", b.GetId(), i, uptime)
	}
}

func RunCustomBot(channel string) {
	userConfig, err := config.LoadUserConfig(channel)
	if err != nil {
		fmt.Println(err)
		return
	}
	//creates bot and runs it in a goroutine
	ch := make(chan interface{})
	currentTime := time.Now()
	listOfUserCommands := make([]commands.Command, 0)
	newBot := &bot.Bot{-999, userConfig, listOfUserCommands, currentTime, ch}
	newBot.AddCommand()
	fmt.Println(newBot)
	go newBot.StartBot()
}
