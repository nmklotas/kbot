package bot

import (
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

func NewPosts(client *model.Client4, webSocketClient *model.WebSocketClient, user *model.User, channel *model.Channel) *Posts {
	posts := Posts{client, webSocketClient, user, channel}
	return &posts
}

func (p Posts) Disconnect() {
	p.webSocketClient.Close()
}

func (p Posts) Create(message string, messageToReplyId string) error {
	post := &model.Post{}
	post.ChannelId = p.channel.Id
	post.Message = message
	post.RootId = messageToReplyId

	_, resp := p.client.CreatePost(post)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}

func (b Posts) Subscribe(callback PostCallback) {
	go func() {
		for event := range b.webSocketClient.EventChannel {
			b.onMessage(event, callback)
		}
	}()
}

func (b Posts) onMessage(event *model.WebSocketEvent, callback PostCallback) {
	if event.Event != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	post := model.PostFromJson(strings.NewReader(event.Data["post"].(string)))
	if post == nil {
		return
	}

	if post.Id == b.user.Id {
		return
	}

	callback(post)
}
