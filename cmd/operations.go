package cmd

import (
	"bufio"
	"fmt"
	"os"
)

var cmdList = []string{"hi", "lol"}

func ListenCMD() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if cmd[0] == scanner.Text() {
			fmt.Println("WORKS")
		}
	}
}
