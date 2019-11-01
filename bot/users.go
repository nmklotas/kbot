package bot

import (
	"github.com/mattermost/mattermost-server/model"
)

type Users struct {
	client *model.Client4
}

func NewUsers(client *model.Client4) *Users {
	users := Users{client}
	return &users
}

func (u Users) Name(userId string) (*string, error) {
	user, resp := u.client.GetUser(userId, "")
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &user.Username, nil
}
