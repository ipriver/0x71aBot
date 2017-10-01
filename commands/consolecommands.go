package commands

type ConsoleCommand struct {
	Command
	description string
	argument    string
}

func (c *ConsoleCommand) Constructor(name string, f interface{}, descr string, arg string) {
	c.Command.Constructor(name, f)
	c.description = descr
	c.argument = arg
}

var (
	cmdExit   = &ConsoleCommand{}
	cmdStatus = &ConsoleCommand{}
	cmdStop   = &ConsoleCommand{}
	cmdInfo   = &ConsoleCommand{}
	cmdRun    = &ConsoleCommand{}
	cmdHelp   = &ConsoleCommand{}
)

var consoleCmdList = []*ConsoleCommand{
	cmdExit,
	cmdStatus,
	cmdStop,
	cmdInfo,
	cmdRun,
	cmdHelp,
}

func init() {
	cmdExit.Constructor("exit", notImplemented, "use to gently exit the programm", "")
	cmdStatus.Constructor("status", notImplemented, "shows current status information of all bots", "")
	cmdStop.Constructor("stop", notImplemented, "stops user bot goroutine", "")
	cmdInfo.Constructor("info", notImplemented, "detailed information of channel bot", "")
	cmdRun.Constructor("run", notImplemented, "creates new bot goroutine", "")
	cmdHelp.Constructor("help", notImplemented, "help information", "")

}

func FindConsoleCommand(name string) *ConsoleCommand {
	for _, v := range consoleCmdList {
		if v.Name == name {
			return v
		}
	}
	return nil
}

var notImplemented = func(args ...interface{}) {}

/*func getStatus(args ...interface{}) {
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
*/
