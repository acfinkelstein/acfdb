package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/acfinkelstein/acfdb/internal/app"
)

func main() {
	fmt.Println("Welcome to ACF Database. Please enter a command.")

	commandScanner := bufio.NewScanner(os.Stdin)

	var command string

	for command != "END" {
		fmt.Print(">>")
		commandScanner.Scan()

		command = commandScanner.Text()

		if command != "END" {
			response, err := app.Interpret(command)

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Println(response)
		}
	}
}
