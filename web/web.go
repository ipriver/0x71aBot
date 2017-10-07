package web

import (
	"../monitor"
	"fmt"
	"net/http"
)

func WebService() {
	//listens user input and calls functions
	go monitor.ListenCMD()

	mux := http.NewServeMux()
	mux.HandleFunc("/", RunBotHandler)
	mux.HandleFunc("/bot_info/", InfoBotHandler)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()

	//closing connection to DBs
	defer func() {
		fmt.Println("closed db connections")
		//config.Rc.Close()
		//config.Db.Close()
	}()
}
