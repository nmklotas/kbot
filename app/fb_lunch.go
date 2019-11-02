package app

import (
	"errors"
	"kbot/bot"
	c "kbot/config"
	fb "kbot/fbposts"
	"time"

	. "github.com/ahmetb/go-linq/v3"
)

type FbLunch struct {
	config c.Config
	posts  *bot.Posts
}

func NewFbLunch(config c.Config, posts *bot.Posts) *FbLunch {
	return &FbLunch{config, posts}
}

func (f FbLunch) PostOffers(time time.Time) error {
	interval := fb.CheckInterval{
		Min: f.config.PostCheckIntervalBeforeMin,
		Max: f.config.PostCheckIntervalAfterMax,
	}

	if !fb.IsTimeToCheck(time, f.config.PostTime, interval) {
		return nil
	}

	fbPosts, err := fb.FindPosts(f.config.FbPageId, f.config.FbAccessToken)
	if err != nil {
		return err
	}

	fbPost, ok := From(*fbPosts).
		FirstWithT(func(p fb.FbPost) bool {
			return fb.ContainsWord(p, f.config.PostPhraseToSearch) && fb.IsPostedToday(p.CreatedTime)
		}).(fb.FbPost)

	if !ok {
		return errors.New("Fb post not found")
	}

	return f.CreatePost(fbPost)
}

func (f FbLunch) CreatePost(post fb.FbPost) error {
	if post.Text != "" {
		if err := f.posts.Create(post.Text); err != nil {
			return err
		}
	}

	if post.Picture != "" {
		if err := f.posts.Create(post.Picture); err != nil {
			return err
		}
	}

	return nil
}
