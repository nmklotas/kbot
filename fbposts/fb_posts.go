package fbposts

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/ahmetb/go-linq/v3"
	fb "github.com/huandu/facebook"
)

type FbPost struct {
	Text        string
	CreatedTime time.Time
	Picture     string
}

func ContainsWord(post FbPost, word string) bool {
	return strings.Contains(strings.ToUpper(post.Text), strings.ToUpper(word))
}

func FindPosts(pageId int, accessToken string) (*[]FbPost, error) {
	res, err := fb.Get("/"+strconv.Itoa(pageId)+"/posts?fields=message,created_time", fb.Params{
		"access_token": accessToken,
	})

	if err != nil {
		if e, ok := err.(*fb.Error); ok {
			return nil, fmt.Errorf(
				"Facebook error. [message:%v] [type:%v] [code:%v] [subcode:%v] [trace:%v]",
				e.Message, e.Type, e.Code, e.ErrorSubcode, e.TraceID)
		}
	}

	return createFbPosts(res)
}

func createFbPosts(res fb.Result) (*[]FbPost, error) {
	var items []fb.Result
	err := res.DecodeField("data", &items)
	if err != nil {
		return nil, err
	}

	var timeParseErr error
	posts := []FbPost{}
	From(items).
		SelectT(func(r fb.Result) FbPost {
			createdTime, err := time.Parse(time.RFC3339, r["created_time"].(string))
			if err != nil {
				timeParseErr = err
			}
			return FbPost{
				Text:        r["message"].(string),
				CreatedTime: createdTime,
				Picture:     r["full_picture"].(string),
			}
		}).
		ToSlice(&posts)

	if timeParseErr != nil {
		return nil, err
	}

	return &posts, nil
}
