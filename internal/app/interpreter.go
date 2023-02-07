package app

import (
	"fmt"
	"strings"
)

func Interpret(command string) (string, error) {
	var response string

	args := strings.Split(command, " ")

	switch args[0] {
	case "GET":
		err := checkArgs(args, 2, "[name]")

		if err != nil {
			return "", err
		}

		response = getValue(args[1])

	case "SET":
		err := checkArgs(args, 3, "[name] [value]")

		if err != nil {
			return "", err
		}

		setValue(args[1], args[2])

	case "DELETE":
	case "COUNT":
	default:
		response = "Unknown Command"
	}

	return response, nil
}

func checkArgs(args []string, expected int, usage string) error {
	var expectedUsageNotice = "Expected %s usage: %s"

	if len(args) != expected {
		return fmt.Errorf(expectedUsageNotice, args[0], args[0]+" "+usage)
	}

	return nil
}
