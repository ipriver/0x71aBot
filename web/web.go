package web

import (
	"net/http"
)

func WebService() {
	//consoleCM := make(chan string)

	mux := http.NewServeMux()
	mux.HandleFunc("/", RunBotHandler)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
