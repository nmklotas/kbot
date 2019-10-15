package main

import (
	"fmt"
	"kbot/bot"
	"kbot/command"
	"kbot/config"
	"os"
	"os/signal"

	"github.com/mattermost/mattermost-server/model"
	. "github.com/ahmetb/go-linq/v3"
)

func main() {
    config, err := config.ReadConfig("config")
    if err != nil {
        panic("Failed to read config")
    }
    
	bot := bot.NewMatterMostBot(bot.Connection{
		ServerUrl: config.ServerUrl,
		Channel:   config.Channel,
		Email:     config.Email,
		Password:  config.Password,
		Team:      config.Team,
	})

	posts, err := bot.JoinChannel()
	if err != nil {
		panic("Failed to join channel")
	}

	setupDisconnectOnOsInterrupt(posts)

	posts.Subscribe(func(post *model.Post) {
		if !command.IsBotCommand(post.Message) {
			return
		}

		ordersStore, err := command.NewOrdersStore()
		if err != nil {
			panic("Can't create orders store")
        }
        
        commands := []command.Command{
            command.NewSaveOrderCommand(ordersStore, posts),
            command.NewListOrdersCommand(ordersStore, posts),
        }
        commands = append(commands, command.NewHelpCommand(posts, commands))

		executeCommands(commands, post)
		defer ordersStore.Close()
	})

	select {}
}

func executeCommands(commands []command.Command, post *model.Post) {
    From(commands).ForEachT(func(c command.Command) {
        message := command.Message{Text: post.Message, UserId: post.UserId}
        if c.CanHandle(message) {
            commandErr := c.Handle(message)
            if commandErr != nil {
                fmt.Println(commandErr.Error())
            }
        }
    })
}

func setupDisconnectOnOsInterrupt(posts *bot.Posts) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			posts.Disconnect()
			os.Exit(0)
		}
	}()
}
