package monitor

import (
	"../commands"
	"bufio"
	"fmt"
	"os"
)

func ListenCMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res := commands.FindConsoleCommand(scanner.Text())
		if res != nil {
			res.Call()
		} else {
			fmt.Println("Command not found")
		}
	}
}
