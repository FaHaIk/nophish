package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fahaik/nophish/constants"
)

func Input(message string, wrongMessage string, validate func(string) bool) string {
	repeat := true
	badinput := false
	value := ""

	for repeat {
		if !badinput {
			fmt.Print(constants.Green + "[" + constants.White + "::" + constants.Green + "] " + message + constants.Reset)
		} else {
			fmt.Println(constants.Red + "[" + constants.White + "--" + constants.Red + "] " + wrongMessage + constants.Reset)
			fmt.Print(constants.Green + "[" + constants.White + "::" + constants.Green + "] " + message + constants.Reset)
		}

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		// Validating user input
		if validate(scanner.Text()) {
			value = scanner.Text()
			repeat = false
		} else {
			badinput = true
		}
	}
	return value
}
