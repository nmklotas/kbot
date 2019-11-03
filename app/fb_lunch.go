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
	config          c.Config
	messages        *bot.Messages
	lunchRepository *LunchRepository
	checkInterval   fb.CheckInterval
}

func NewFbLunch(config c.Config, messages *bot.Messages, lunchRepository *LunchRepository) *FbLunch {
	instance := FbLunch{
		config,
		messages,
		lunchRepository,
		fb.CheckInterval{
			Min: config.PostCheckIntervalBeforeMin,
			Max: config.PostCheckIntervalAfterMax,
		}}

	return &instance
}

func (f FbLunch) PostOffers(time time.Time) error {
	if !fb.IsTimeToCheck(time, f.config.PostTime, f.checkInterval) {
		return nil
	}

	postExists := f.lunchRepository.Any(func(l Lunch) bool {
		return fb.IsPostedToday(l.CreatedAt)
	})

	if postExists {
		return nil
	}

	fbPosts, err := fb.FindPosts(f.config.FbPageId, f.config.FbAccessToken)
	if err != nil {
		return err
	}

	todaysPost, err := findTodaysPost(f.config.PostPhraseToSearch, fbPosts)
	if err != nil {
		return err
	}

	if err := f.lunchRepository.Save(Lunch{}); err != nil {
		return err
	}

	return f.SendMessage(*todaysPost)
}

func (f FbLunch) SendMessage(post fb.FbPost) error {
	if post.Text != "" {
		if err := f.messages.Send(post.Text); err != nil {
			return err
		}
	}

	if post.Picture != "" {
		if err := f.messages.Send(post.Picture); err != nil {
			return err
		}
	}

	return nil
}

func findTodaysPost(phraseToSearch string, posts *[]fb.FbPost) (*fb.FbPost, error) {
	fbPost, ok := From(*posts).
		FirstWithT(func(p fb.FbPost) bool {
			return fb.ContainsWord(p, phraseToSearch) && fb.IsPostedToday(p.CreatedTime)
		}).(fb.FbPost)

	if !ok {
		return nil, errors.New("Fb post not found")
	}

	return &fbPost, nil
}
