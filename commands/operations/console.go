package operations

import (
	//"../../bot"
	"../../commands"
	"fmt"
	"os"
)

func FindConsoleCommand(name string) *commands.ConsoleCommand {
	for _, obj := range consoleCmdList {
		if obj.GetName() == name {
			return obj
		}
	}
	return nil
}

var (
	cmdExit   = &commands.ConsoleCommand{}
	cmdStatus = &commands.ConsoleCommand{}
	cmdStop   = &commands.ConsoleCommand{}
	cmdInfo   = &commands.ConsoleCommand{}
	cmdRun    = &commands.ConsoleCommand{}
	cmdHelp   = &commands.ConsoleCommand{}
)

var consoleCmdList = []*commands.ConsoleCommand{
	cmdExit,
	cmdStatus,
	cmdStop,
	cmdInfo,
	cmdRun,
	cmdHelp,
}

func init() {
	cmdExit.Constructor("exit", programmExit, "use to gently exit the programm", make([]string, 0))
	cmdStatus.Constructor("status", notImplemented, "shows current status information of all bots", make([]string, 0))
	cmdStop.Constructor("stop", notImplemented, "stops user bot goroutine", make([]string, 0))
	cmdInfo.Constructor("info", notImplemented, "detailed information of channel bot", make([]string, 0))
	cmdRun.Constructor("run", notImplemented, "creates new bot goroutine", make([]string, 0))
	cmdHelp.Constructor("help", help, "help information", make([]string, 0))

}

var notImplemented = func(args ...interface{}) {}

func help() {
	fmt.Println("Avalibale commands:")
	for i, c := range consoleCmdList {
		fmt.Printf("%v) %s\n", i+1, c.GetName())
	}
}

func programmExit() {
	//TODO: check and close all connections
	os.Exit(0)
}

/*
func getStatus() {
	fmt.Printf("Online bots: %v\n", len(bot.OnlineBots))
	for i, b := range bot.OnlineBots {
		uptime := time.Since(b.UpTime)
		fmt.Printf("id: %d, Channel: %s, Uptime: %v\n", b.GetId(), i, uptime)
	}
}
*/
