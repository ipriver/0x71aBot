package monitor

import (
	"../commands/operations"
	"bufio"
	"fmt"
	"os"
)

func ListenCMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		res := operations.FindConsoleCommand(scanner.Text())
		if res != nil {
			res.Call()
		} else {
			fmt.Println("Command not found")
		}
	}
}
