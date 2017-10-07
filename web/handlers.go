package web

import (
	"../bot"
	"encoding/json"
	"net/http"
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

		if err != nil {
			panic(err)
		}
		err = bot.LoadBot(data.Bot_id, data.Channel)
		if err != nil {
			rw.WriteHeader(500)
		} else {
			rw.WriteHeader(200)
		}

	default:
		rw.WriteHeader(403)
	}
}
