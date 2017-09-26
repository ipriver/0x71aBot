package handlers

import (
	"../bot"
	"../commands"
	"../config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserJSONdata struct {
	Bot_id  int    `json:"user_id"`
	Channel string `json:"channel"`
}

func RunBotHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		//parse data from request json
		decoder := json.NewDecoder(req.Body)
		defer req.Body.Close()
		data := UserJSONdata{}
		err := decoder.Decode(&data)

		//check bot in redis by its id
		/*bot := checkBotInCache(data.Bot_id)
		if bot == nil {
			fmt.Println("no cache")
		}*/

		//creates userconfig from data
		userConfig, err := config.LoadUserConfig(data.Channel)
		if err != nil {
			fmt.Println(err)
			return
		}
		//creates bot and runs it in a goroutine
		ch := make(chan interface{})
		currentTime := time.Now()
		listOfUserCommands := make([]commands.Command, 10)
		newBot := &bot.Bot{data.Bot_id, userConfig, listOfUserCommands, currentTime, ch}
		go newBot.StartBot()
		//send to the client response code
		rw.WriteHeader(200)
	default:
		rw.WriteHeader(404)
	}
}

func InfoBotHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		defer req.Body.Close()
		data := UserJSONdata{}
		err := decoder.Decode(&data)
		if err != nil {

		}
		/* check in cache, connect to DB, response with bot information; */
		rw.WriteHeader(200)
	default:
		rw.WriteHeader(404)
	}
}
