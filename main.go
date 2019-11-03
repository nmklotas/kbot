package main

import (
	"kbot/app"
	"kbot/bot"
	"kbot/command"
	"kbot/config"
	fb "kbot/fbposts"
	"kbot/log"
	"os"
	"os/signal"
	"time"

	"github.com/mattermost/mattermost-server/model"
)

func main() {
	config := readConfig()
	connection := createConnection(config)
	apiClient := model.NewAPIv4Client(connection.ServerUrl)
	users := bot.NewUsers(apiClient)

	ordersStore, err := command.OpenOrdersStore()
	panicOnError(err)
	defer ordersStore.Close()

	botChannel, err := bot.NewChannel(apiClient, connection).Join()
	panicOnError(err)

	messages, err := bot.ListenMessages(apiClient, botChannel.Bot, botChannel.Channel)
	panicOnError(err)
	unsubscribeFromPostsOnInterupt(messages)

	fbLunch := app.NewFbLunch(config, messages)
	logger := log.NewLogger()
	botCommands := app.NewBotCommands(
		createCommands(ordersStore, messages, users),
		logger)

	go func() {
		messages.Subscribe(func(m *model.Post) {
			botCommands.Execute(m)
		})
	}()

	go func() {
		fb.StartTicking(func(t time.Time) {
			logger.Info("Tick")
			if err := fbLunch.PostOffers(t); err == nil {
				messages.Send("Orders available. Type 'kbot order {letter}' to order!")
			}
		}, config.PostCheckIntervalMin)
	}()

	select {}
}

func createCommands(ordersStore *command.OrdersStore, messages *bot.Messages, users *bot.Users) []command.Command {
	commands := []command.Command{
		command.NewOrderCommand(ordersStore, messages, users),
		command.NewForgetOrderCommand(ordersStore, messages, users),
		command.NewListOrdersCommand(ordersStore, messages),
	}
	helpCommand := command.NewHelpCommand(messages, commands)
	return append(commands, helpCommand)
}

func readConfig() config.Config {
	config, err := config.ReadConfig("config")
	panicOnError(err)
	return config
}

func createConnection(c config.Config) bot.Connection {
	return bot.Connection{
		ServerUrl: c.ServerUrl,
		Channel:   c.Channel,
		Email:     c.Email,
		Password:  c.Password,
		Team:      c.Team,
	}
}

func unsubscribeFromPostsOnInterupt(messages *bot.Messages) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			messages.Close()
			os.Exit(0)
		}
	}()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
