package command

import (
	"kbot/bot"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
)

type HelpCommand struct {
	messages *bot.Messages
	commands []Command
}

func NewHelpCommand(messages *bot.Messages, commands []Command) *HelpCommand {
	return &HelpCommand{messages, commands}
}

func (p HelpCommand) CanHandle(message Message) bool {
	return strings.EqualFold(message.Text, "kbot help")
}

func (p HelpCommand) Handle(message Message) error {
	return p.messages.Send(p.createHelpText())
}

func (p HelpCommand) Help() string {
	return "returns help of all commands"
}

func (p HelpCommand) createHelpText() string {
	var commandsTexts []string

	From(p.commands).
		SelectT(func(c Command) string {
			return c.Help()
		}).
		Concat(From([]string{"kbot help. " + p.Help()})).
		ToSlice(&commandsTexts)

	return strings.Join(commandsTexts, "\r\n")
}
