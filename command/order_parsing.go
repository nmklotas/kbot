package command

import (
	"errors"
	"strings"
)

func ParseOrderToRemove(message string) (*string, error) {
	parts := strings.Split(message, " ")

	if len(parts) != 3 {
		return nil, errors.New("Invalid command")
	}

	if !strings.EqualFold(parts[0], "kbot") {
		return nil, errors.New("Invalid command")
	}

	if !strings.EqualFold(parts[1], "forget") {
		return nil, errors.New("Invalid command")
	}

	if parts[2] == "" || len(parts[2]) != 1 {
		return nil, errors.New("Invalid command")
	}

	return &parts[2], nil
}

func ParseOrderToAdd(message string) (*string, error) {
	parts := strings.Split(message, " ")

	if len(parts) != 3 {
		return nil, errors.New("Invalid command")
	}

	if !strings.EqualFold(parts[0], "kbot") {
		return nil, errors.New("Invalid command")
	}

	if !strings.EqualFold(parts[1], "order") {
		return nil, errors.New("Invalid command")
	}

	if parts[2] == "" || len(parts[2]) != 1 {
		return nil, errors.New("Invalid command")
	}

	return &parts[2], nil
}

func ParseListOrders(message string) bool {
	parts := strings.Split(message, " ")

	if len(parts) != 2 {
		return false
	}

	if !strings.EqualFold(parts[0], "kbot") {
		return false
	}

	if !strings.EqualFold(parts[1], "list") {
		return false
	}

	return true
}
