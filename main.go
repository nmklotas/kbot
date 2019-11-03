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

	posts, err := bot.SubscribeToPosts(apiClient, botChannel.Bot, botChannel.Channel)
	panicOnError(err)
	unsubscribeFromPostsOnInterupt(posts)

	fbLunch := app.NewFbLunch(config, posts)
	logger := log.NewLogger()
	botCommands := app.NewBotCommands(
		createCommands(ordersStore, posts, users),
		logger)

	go func() {
		posts.Subscribe(func(post *model.Post) {
			botCommands.Execute(post)
		})
	}()

	go func() {
		fb.StartTicking(func(time time.Time) {
			logger.Info("Tick")
			fbLunch.PostOffers(time)
		}, config.PostCheckIntervalMin)
	}()

	select {}
}

func createCommands(ordersStore *command.OrdersStore, posts *bot.Posts, users *bot.Users) []command.Command {
	commands := []command.Command{
		command.NewOrderCommand(ordersStore, posts, users),
		command.NewForgetOrderCommand(ordersStore, posts, users),
		command.NewListOrdersCommand(ordersStore, posts),
	}
	helpCommand := command.NewHelpCommand(posts, commands)
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

func unsubscribeFromPostsOnInterupt(posts *bot.Posts) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			posts.Close()
			os.Exit(0)
		}
	}()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
