package app

import (
	"fmt"
	"kbot/command"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/mattermost/mattermost-server/model"
)

func ExecuteCommands(commands []command.Command, post *model.Post) {
	if !command.IsBotCommand(post.Message) {
		return
	}

	From(commands).ForEachT(func(c command.Command) {
		message := command.Message{Text: post.Message, UserId: post.UserId}
		if !c.CanHandle(message) {
			return
		}

		if err := c.Handle(message); err != nil {
			fmt.Println(err)
		}
	})
}
