package command

import "strings"

type Message struct {
	Text   string
	UserId string
}

type Command interface {
	CanHandle(message Message) bool
	Handle(message Message) error
	Help() string
}

func IsBotCommand(message string) bool {
	return strings.HasPrefix(message, "kbot")
}
