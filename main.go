package main

import (
	"./bot"
	"./config"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", bot.CreateBotHandler)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
	defer func() {
		config.Rc.Close()
		config.Db.Close()
	}()
}
