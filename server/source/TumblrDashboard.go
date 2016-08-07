package source

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/kanosaki/picwall/server/core"
	"github.com/kanosaki/picwall/server/tumblr"
)

type TumblrDashboard struct {
	transformer *core.TumblrTransformer
	output      chan *core.Entry
	sourceCh    chan *tumblr.Post
	tumblr      *core.Tumblr
	seeker      *core.TumblrSeeker
	stopped     bool
}

func NewTumblrDashboard(tblr *core.Tumblr, transformer *core.TumblrTransformer) (*TumblrDashboard, error) {
	output := make(chan *core.Entry)
	sourceCh := make(chan *tumblr.Post)
	seeker := core.NewTumblrSeeker(func(v map[string]string) ([]tumblr.Post, error) {
		logrus.Infof("Tumblr Dashbord: %v", v)
		res, err := tblr.Api.Dashboard(v)
		if err != nil {
			return nil, err
		}
		posts := make([]tumblr.Post, 0, len(res.Posts))
		for _, ps := range res.Posts {
			var post tumblr.Post
			err := json.Unmarshal(ps, &post)
			if err != nil {
				return nil, err
			}
			posts = append(posts, post)
		}
		return posts, nil
	}, sourceCh)
	ret := &TumblrDashboard{
		seeker:      seeker,
		tumblr:      tblr,
		output:      output,
		sourceCh:    sourceCh,
		transformer: transformer,
	}
	go ret.convertTweetPump()
	seeker.Start()
	return ret, nil
}

func (tl *TumblrDashboard) convertTweetPump() {
	for !tl.stopped {
		post := <-tl.sourceCh
		entry, ok, err := tl.transformer.ToEntry(post)
		if err != nil {
			logrus.Errorf("Error while converting tweet: %s", err)
		} else {
			if ok {
				tl.output <- entry
			} else {
				logrus.Debugf("Skipped: %d %s", post.ID, post.Text)
			}
		}
	}
}

func (td *TumblrDashboard) Faucet() <-chan *core.Entry {
	return td.output
}

func (td *TumblrDashboard) Drain(count int) {

}

func (td *TumblrDashboard) Shutdown() {

}
