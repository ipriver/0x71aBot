package bot

import (
	"../config"
	"fmt"
	_ "github.com/garyburd/redigo/redis"
)

func checkBotInCache(id int) *Bot {
	_, err := config.Rc.Do("GET", fmt.Sprintf("bot:bots:%d", id))
	if err != nil {
		fmt.Println("no bot in cache")
	}
	return nil
}
