package main

import (
	"fmt"

	"github.com/fahaik/nophish/constants"
)

func PrintBanner() {
	ascii := `
███╗   ██╗ ██████╗ ██████╗ ██╗  ██╗██╗███████╗██╗  ██╗
████╗  ██║██╔═══██╗██╔══██╗██║  ██║██║██╔════╝██║  ██║
██╔██╗ ██║██║   ██║██████╔╝███████║██║███████╗███████║
██║╚██╗██║██║   ██║██╔═══╝ ██╔══██║██║╚════██║██╔══██║
██║ ╚████║╚██████╔╝██║     ██║  ██║██║███████║██║  ██║
╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝  ╚═╝╚═╝╚══════╝╚═╝  ╚═╝
			`
	fmt.Print(constants.Blue + ascii)
	fmt.Println("   by Fabian Ikizoglu (v0.0.1)" + constants.Reset + "\n")
}

func PrintServiceNumbers() {
	// fmt.Println(constants.Green + "[" + constants.White + "::" + constants.Green + "] " + "Pick a service to imitate.\n")
	for index, service := range Services {
		number := fmt.Sprintf("%02d", index)

		if (index+1)%3 == 0 {
			fmt.Println(constants.Green + "[" + constants.White + number + constants.Green + "] " + service + "\t  ")
		} else {
			fmt.Print(constants.Green + "[" + constants.White + number + constants.Green + "] " + service + "\t  ")
		}
	}
	fmt.Println("")
}

func PrintRunning(serverOne bool, serverTwo bool) {
	if serverOne && serverTwo {
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "phishing server: " + constants.Green + "running" + constants.Reset)
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "admin panel: " + constants.Green + "running" + constants.Reset)
	} else if !serverOne && serverTwo {
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "phishing server: " + constants.Red + "stopped" + constants.Reset)
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "admin panel: " + constants.Green + "running" + constants.Reset)
	} else if serverOne && !serverTwo {
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "phishing server: " + constants.Green + "running" + constants.Reset)
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "admin panel: " + constants.Red + "stopped" + constants.Reset)
	} else {
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "phishing server: " + constants.Red + "stopped" + constants.Reset)
		fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + "admin panel: " + constants.Red + "stopped" + constants.Reset)
	}
}

func PrintChoice(message string) {
	fmt.Println("\n" + constants.Yellow + "[" + constants.White + "::" + constants.Yellow + "] " + message + constants.Reset + "\n")
	// fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + constants.Reset)
	// fmt.Println(constants.Blue + "[" + constants.White + ">>" + constants.Blue + "] " + message + constants.Reset)
	// fmt.Println(constants.Blue + "[" + constants.White + "::" + constants.Blue + "] " + constants.Reset)
}
