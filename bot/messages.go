package bot

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

type MessageCallback func(*model.Post)

type Messages struct {
	client   *model.Client4
	wsClient *model.WebSocketClient
	user     *model.User
	channel  *model.Channel
}

func ListenMessages(client *model.Client4, user *model.User, channel *model.Channel) (*Messages, error) {
	wsClient, err := model.NewWebSocketClient4(createWebSocketServerUrl(client), client.AuthToken)
	if err != nil {
		return nil, err
	}

	wsClient.Listen()

	posts := Messages{client, wsClient, user, channel}
	return &posts, nil
}

func (p Messages) Close() {
	p.wsClient.Close()
}

func (p Messages) Send(message string) error {
	post := &model.Post{
		ChannelId: p.channel.Id,
		Message:   message,
	}

	if _, resp := p.client.CreatePost(post); resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (p Messages) Subscribe(callback MessageCallback) {
	for event := range p.wsClient.EventChannel {
		p.onMessage(event, callback)
	}
}

func (p Messages) onMessage(event *model.WebSocketEvent, callback MessageCallback) {
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post == nil {
		return
	}

	if post.UserId == p.user.Id {
		return
	}

	callback(post)
}

func createWebSocketServerUrl(client *model.Client4) string {
	address := strings.Replace(client.Url, "https://", "", 1)
	websocketAddress := fmt.Sprintf("wss://%s", address)
	return websocketAddress
}
