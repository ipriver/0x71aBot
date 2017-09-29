package cmd

import (
	"../bot"
	"../commands"
	"../config"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var cmdList = commands.ConsoleCmdList

func ListenCMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		switch scanner.Text() {
		case cmdList[0]:
			os.Exit(0)
		case cmdList[1]:
			getOnlineBots()
		case cmdList[5]:
			help()
		}
		run, _ := regexp.MatchString(cmdList[4], scanner.Text())
		if run == true {
			channel := strings.Fields(scanner.Text())
			if len(channel) > 1 {
				runCustomBot(channel[1])
			}
		}
		run, _ = regexp.MatchString(cmdList[3], scanner.Text())
		if run == true {
			channel := strings.Fields(scanner.Text())
			if len(channel) > 1 {
				getBotInfo(channel[1])
			}
		}
	}
}
