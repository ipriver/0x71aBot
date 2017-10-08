package commands

type ConsoleCommand struct {
	Command
	Description string
	Arguments   []string
}

func Constructor(name string, f interface{}, descr string, args []string) *ConsoleCommand {
	c := new(ConsoleCommand)
	c.Command = *NewCommand(name, f)
	c.Description = descr
	c.Arguments = args
	return c
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
