package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UpdateConfig() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(Config)

	fmt.Printf("1)%s Enter new host adress:\n", Config.HostAddr)
	scanner.Scan()
	Config.HostAddr = scanner.Text()

	fmt.Printf("2)%d Enter new port:\n", Config.Port)
	scanner.Scan()
	Config.Port, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("3)%s Enter new Login for Bot:\n", Config.LoginBotName)
	scanner.Scan()
	Config.LoginBotName = scanner.Text()

	fmt.Printf("4)%s Enter new OATH:\n", Config.LogOath)
	scanner.Scan()
	Config.LogOath = scanner.Text()

	fmt.Println("Save new config? yes/no")
	scanner.Scan()
	switch strings.ToLower(scanner.Text()) {
	case "yes":
		fallthrough
	case "1":
		err = Config.Save()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Saved")
	}
	os.Exit(0)
}
