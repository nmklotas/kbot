package fbposts

import (
    "fmt"
    fb "github.com/huandu/facebook"
    . "github.com/ahmetb/go-linq/v3"
)

type FbPost struct {
    Text string
    CreatedTime string
}

func FindPosts(pageId string, accessToken string) (*[]FbPost, error) {  
	res, err := fb.Get("/" + pageId + "/posts?fields=message,created_time", fb.Params{
		"access_token": accessToken,
	})

	if err != nil {
		if e, ok := err.(*fb.Error); ok {
			return nil, fmt.Errorf(
                "Facebook error. [message:%v] [type:%v] [code:%v] [subcode:%v] [trace:%v]",
                e.Message, e.Type, e.Code, e.ErrorSubcode, e.TraceID)
		}
    }

    var items []fb.Result
    err = res.DecodeField("data", &items)
    if err != nil {
        return nil, err 
    }

    posts := []FbPost {}
	From(items).
		SelectT(func(r fb.Result) FbPost {
            return FbPost{
                Text: r["message"].(string), 
                CreatedTime: r["created_time"].(string),
            }
		}).
		ToSlice(&posts)

    return &posts, nil
}
