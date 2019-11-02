package app

import (
	"kbot/command"
	"strings"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/mattermost/mattermost-server/model"
	"github.com/sirupsen/logrus"
)

type BotCommands struct {
	commands []command.Command
	logger   *logrus.Logger
}

func NewBotCommands(c []command.Command, l *logrus.Logger) *BotCommands {
	return &BotCommands{c, l}
}

func (b BotCommands) Execute(post *model.Post) {
	if !b.isBotCommand(post) {
		return
	}

	From(b.commands).ForEachT(func(c command.Command) {
		message := command.Message{Text: post.Message, UserId: post.UserId}
		if !c.CanHandle(message) {
			return
		}

		if err := c.Handle(message); err != nil {
			b.logger.Log(logrus.ErrorLevel, err)
		}
	})
}

func (b BotCommands) isBotCommand(p *model.Post) bool {
	return strings.HasPrefix(p.Message, "kbot")
}
