package app

import (
	"errors"
	"fmt"
	"strconv"
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
		err := checkArgs(args, 2, "[name]")

		if err != nil {
			return "", err
		}

		deleteValue(args[1])
	case "COUNT":
		err := checkArgs(args, 2, "[value]")

		if err != nil {
			return "", err
		}

		response = strconv.Itoa(countValues(args[1]))
	case "BEGIN":
		err := checkArgs(args, 1, "")

		if err != nil {
			return "", err
		}

		beginTransaction()
	case "ROLLBACK":
		err := checkArgs(args, 1, "")

		if err != nil {
			return "", err
		}

		if !rollbackTransaction() {
			response = "TRANSACTION NOT FOUND"
		}
	case "COMMIT":
		err := checkArgs(args, 1, "")

		if err != nil {
			return "", err
		}

		commitTransaction()
	default:
		return "", errors.New("unknown command")
	}

	return response, nil
}

func checkArgs(args []string, expected int, usage string) error {
	var expectedUsageNotice = "%s expects %d params. Usage: %s"

	if len(args) != expected {
		return fmt.Errorf(expectedUsageNotice, args[0], expected-1, args[0]+" "+usage)
	}

	return nil
}
