package app

import (
	"fmt"
	"kbot/bot"
	"kbot/config"
	fb "kbot/fbposts"
	"time"

	. "github.com/ahmetb/go-linq/v3"
)

func PostLunchOffers(time time.Time, config config.Config, posts *bot.Posts) {
	interval := fb.CheckInterval{
		Min: config.PostCheckIntervalBeforeMin,
		Max: config.PostCheckIntervalAfterMax,
	}

	if !fb.IsTimeToCheck(time, config.PostTime, interval) {
		return
	}

	fmt.Printf("Post check %s", time)
	fbPosts, err := fb.FindPosts(config.FbPageId, config.FbAccessToken)
	if err != nil {
		fmt.Print(err)
	}

	fbPost := From(fbPosts).
		FirstWithT(func(p fb.FbPost) bool {
			return fb.ContainsWord(p, config.PostPhraseToSearch) && fb.IsPostedToday(p.CreatedTime)
		}).(fb.FbPost)

	if err := posts.Create(fbPost.Text); err != nil {
		fmt.Print(err)
	}

	if err := posts.Create(fbPost.Picture); err != nil {
		fmt.Print(err)
	}
}
