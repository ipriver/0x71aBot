package bot

import (
	//	"../commands"
	"errors"
)

func LoadBot(id int, channel string) error {
	_, ok := OnlineBots[channel]
	if ok == true {
		return errors.New("Bot is already online")
	}
	b := Constructor(id, channel)

	err := b.Start()
	return err
}

/*func defaultChatCommands() *commands.ChatCommand {
	cm := commands.NewChatCommand("hi")
}*/
