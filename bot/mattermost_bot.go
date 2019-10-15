package bot

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

type Connection struct {
	ServerUrl string
	Email     string
	Password  string
	Team      string
	Channel   string
}

type MatterMostBot struct {
	client     *model.Client4
	connection Connection
}

func NewMatterMostBot(c Connection) *MatterMostBot {
	return &MatterMostBot{model.NewAPIv4Client(c.ServerUrl), c}
}

func (b MatterMostBot) JoinChannel() (*Posts, error) {
	botUser, resp := b.client.Login(b.connection.Email, b.connection.Password)
	if resp.Error != nil {
		return nil, resp.Error
	}

	botTeam, resp := b.client.GetTeamByName(b.connection.Team, "")
	if resp.Error != nil {
		return nil, resp.Error
	}

	botChannel, resp := b.client.GetChannelByName(b.connection.Channel, botTeam.Id, "")
	if resp.Error != nil {
		return nil, resp.Error
	}

	_, resp = b.client.AddChannelMember(botChannel.Id, botUser.Id)
	if resp.Error != nil {
		return nil, resp.Error
	}

	webSocketClient, err := model.NewWebSocketClient4(b.createWebSocketServerUrl(), b.client.AuthToken)
	if err != nil {
		return nil, err
	}

	webSocketClient.Listen()
	return NewPosts(b.client, webSocketClient, botUser, botChannel), nil
}

func (b MatterMostBot) createWebSocketServerUrl() string {
	address := strings.Replace(b.connection.ServerUrl, "https://", "", 1)
	websocketAddress := fmt.Sprintf("wss://%s", address)
	return websocketAddress
}
