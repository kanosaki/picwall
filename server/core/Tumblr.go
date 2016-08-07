package core

import (
	"fmt"
	"github.com/kanosaki/picwall/server/tumblr"
	"github.com/spbr/gotumblr"
)

type Tumblr struct {
	Api *gotumblr.TumblrRestClient
}

func NewTumblr(config *TumblrConfig) (*Tumblr, error) {
	api := gotumblr.NewTumblrRestClient(
		config.ConsumerKey, config.ConsumerSecret,
		config.DefaultAccount.AccessToken, config.DefaultAccount.AccessTokenSecret,
		"https://localhost:4499", "https://api.tumblr.com")
	return &Tumblr{
		Api: api,
	}, nil
}

type TumblrTransformer struct {
	tumblr *Tumblr
}

func NewTumblrTransformer(tumblr *Tumblr) *TumblrTransformer {
	return &TumblrTransformer{
		tumblr: tumblr,
	}
}

func (tt *TumblrTransformer) ToEntry(post *tumblr.Post) (*Entry, bool, error) {
	if len(post.Photos) > 0 {
		e := NewEntry()
		e.Set("id", fmt.Sprintf("tumblr-post-%d", post.ID))
		firstPhoto := post.Photos[len(post.Photos)-1]
		if len(firstPhoto.AltSizes) > 0 {
			e.Set("thumbnail", firstPhoto.AltSizes[len(firstPhoto.AltSizes)-1].URL)
			return e, true, nil
		}
	}
	return nil, false, nil
}
