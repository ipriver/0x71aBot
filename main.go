package main

import (
	//"bufio"
	"flag"
	//"fmt"
	//"os"
	//"strconv"
	//"strings"
	"./web"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "config":
		UpdateConfig()
	default:
		web.WebService()
	}
}

//Upgrades data in config.json
func UpdateConfig() {
	/*scanner := bufio.NewScanner(os.Stdin)
	newconf := config.GlobalConfig{}
	newconf.Load()

	//working with user input
	fmt.Printf("1)%s Enter new host adress:\n", newconf.HostAddr)
	scanner.Scan()
	newconf.HostAddr = scanner.Text()

	fmt.Printf("2)%d Enter new port:\n", newconf.Port)
	scanner.Scan()
	newconf.Port, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("3)%s Enter new Login for Bot:\n", newconf.LoginBotName)
	scanner.Scan()
	newconf.LoginBotName = scanner.Text()

	fmt.Printf("4)%s Enter new OATH:\n", newconf.LogOath)
	scanner.Scan()
	newconf.LogOath = scanner.Text()

	fmt.Println(newconf)
	fmt.Println("Save new config? yes/no")
	scanner.Scan()
	switch strings.ToLower(scanner.Text()) {
	case "yes":
		fallthrough
	case "1":
		err = newconf.Save()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Saved")
	}
	os.Exit(0)
	*/
}
