package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/acfinkelstein/acfdb/internal/app"
)

func main() {
	fmt.Println("Welcome to ACF Database. Please enter a command.")

	app.Init()

	commandScanner := bufio.NewScanner(os.Stdin)

	var command string

	for command != "END" {
		fmt.Print(">>")
		commandScanner.Scan()

		command = commandScanner.Text()

		if command != "END" {
			response, err := app.Interpret(command)

			// interpreter errors throw for invalid commands and for command usage errors
			if err != nil {
				fmt.Println("Error: " + err.Error())
			}

			if response != "" {
				fmt.Println(response)
			}
		}
	}
}
