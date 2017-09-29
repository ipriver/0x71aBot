package commands

import (
	"fmt"
	"os"
	"time"
)

type ConsoleCommand struct {
	Command
	description string
	argument    string
}

func (c *ConsoleCommand) Constructor(name string, f func(args ...interface{}), descr string, arg string) {
	c.Command.Constructor(name, f)
	c.description = descr
	c.argument = arg
}

var cmdExit = &ConsoleCommand{Command{"exit", programmExit}, "use to gently exit the programm", ""}
var cmdStatus = &ConsoleCommand{Command{"status", getStatus}, "shows current status information of all bots", ""}
var cmdStop = &ConsoleCommand{Command{"stop", notImplemented}, "stops user bot goroutine", ""}
var cmdInfo = &ConsoleCommand{Command{"info", notImplemented}, "detailed information of channel bot", ""}
var cmdRun = &ConsoleCommand{Command{"run", runCustomBot}, "creates new bot goroutine", ""}
var cmdHelp = &ConsoleCommand{Command{"help", help}, "help information", ""}

var ConsoleCmdList = []*ConsoleCommand{
	cmdExit,
	cmdStatus,
	cmdStop,
	cmdInfo,
	cmdRun,
	cmdHelp,
}

var notImplemented = func(args ...interface{}) {}

func getStatus(args ...interface{}) {
	fmt.Printf("Online bots: %v\n", len(bot.OnlineBots))
	for i, b := range bot.OnlineBots {
		uptime := time.Since(b.UpTime)
		fmt.Printf("id: %d, Channel: %s, Uptime: %v\n", b.GetId(), i, uptime)
	}
}

func runCustomBot(args ...interface{}) {
	channel := args[0]
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

func help(args ...interface{}) {
	fmt.Println("Avalibale commands:")
	for i, c := range ConsoleCmdList {
		fmt.Printf("%v) %s\n", i+1, c)
	}
}

func programmExit(args ...interface{}) {
	//TODO: check and close all connections
	os.Exit(0)
}
