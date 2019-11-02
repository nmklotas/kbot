package bot

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

type PostCallback func(*model.Post)

type Posts struct {
	client          *model.Client4
	webSocketClient *model.WebSocketClient
	user            *model.User
	channel         *model.Channel
}

func SubscribeToPosts(client *model.Client4, user *model.User, channel *model.Channel) (*Posts, error) {
	webSocketClient, err := model.NewWebSocketClient4(createWebSocketServerUrl(client), client.AuthToken)
	if err != nil {
		return nil, err
	}

	webSocketClient.Listen()

	posts := Posts{client, webSocketClient, user, channel}
	return &posts, nil
}

func (p Posts) Close() {
	p.webSocketClient.Close()
}

func (p Posts) Create(message string) error {
	post := &model.Post{}
	post.ChannelId = p.channel.Id
	post.Message = message
	post.RootId = ""

	_, resp := p.client.CreatePost(post)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (p Posts) Subscribe(callback PostCallback) {
	for event := range p.webSocketClient.EventChannel {
		p.onMessage(event, callback)
	}
}

func (p Posts) onMessage(event *model.WebSocketEvent, callback PostCallback) {
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post == nil {
		return
	}

	if post.Id == p.user.Id {
		return
	}

	callback(post)
}

func createWebSocketServerUrl(client *model.Client4) string {
	address := strings.Replace(client.Url, "https://", "", 1)
	websocketAddress := fmt.Sprintf("wss://%s", address)
	return websocketAddress
}
