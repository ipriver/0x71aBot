package main

import (
	"./config"
	"./web"
	"flag"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "config":
		config.UpdateConfig()
	default:
		web.WebService()
	}
}
