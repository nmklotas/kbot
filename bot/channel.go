package bot

import (
	"github.com/mattermost/mattermost-server/model"
)

type Connection struct {
	ServerUrl string
	Email     string
	Password  string
	Team      string
	Channel   string
}

type Channel struct {
	client     *model.Client4
	connection Connection
}

type BotChannel struct {
	Bot     *model.User
	Channel *model.Channel
}

func NewChannel(client *model.Client4, c Connection) *Channel {
	return &Channel{client, c}
}

func (b Channel) Join() (*BotChannel, error) {
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

	return &BotChannel{
		Bot:     botUser,
		Channel: botChannel,
	}, nil
}
