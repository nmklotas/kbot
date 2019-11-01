package main

import (
	"fmt"
	"kbot/bot"
	"kbot/command"
	"kbot/config"
	"kbot/fbposts"
	"os"
	"os/signal"
	"time"

	. "github.com/ahmetb/go-linq/v3"
	"github.com/mattermost/mattermost-server/model"
)

func main() {
	config := readConfig()
	connection := createConnection(config)
	apiClient := model.NewAPIv4Client(connection.ServerUrl)
	matterMostBot := bot.NewMatterMostBot(apiClient, connection)
	users := bot.NewUsers(apiClient)

	ordersStore, err := command.NewOrdersStore()
	panicOnError(err)
	defer ordersStore.Close()

	botChannel, err := matterMostBot.JoinChannel()
	panicOnError(err)

	posts, err := bot.NewPosts(apiClient, botChannel.Bot, botChannel.Channel)
	panicOnError(err)
	stopListeningForPostsOnInterupt(posts)

	commands := createCommands(ordersStore, posts, users)

	go func() {
		posts.Subscribe(func(post *model.Post) {
			if !command.IsBotCommand(post.Message) {
				return
			}

			executeCommands(commands, post)
		})
	}()

	go func() {
		fbposts.StartTicking(func(time time.Time) {
			fmt.Printf("Post check %s", time)
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

func executeCommands(commands []command.Command, post *model.Post) {
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

func stopListeningForPostsOnInterupt(posts *bot.Posts) {
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
