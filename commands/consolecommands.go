package commands

type ConsoleCommand struct {
	Command
	description string
	arguments   []string
}

type ConsoleCommander interface {
	Call()
	GetDescr() string
	GetArguments() []string
}

func (c *ConsoleCommand) GetDescr() string {
	return c.description
}

func (c *ConsoleCommand) GetArguments() []string {
	return c.arguments
}

func (c *ConsoleCommand) Constructor(name string, f interface{}, descr string, args []string) {
	c.Command.Constructor(name, f)
	c.description = descr
	c.arguments = args
}

/*
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
}*/
