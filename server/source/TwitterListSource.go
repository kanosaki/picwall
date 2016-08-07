package source

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/kanosaki/picwall/server/core"
	"net/url"
)

type TwitterListSource struct {
	output      chan *core.Entry
	sourceCh    chan *anaconda.Tweet
	twitter     *core.Twitter
	listName    string
	seeker      *core.TweetSeeker
	stopped     bool
	transformer *core.TweetTransformer
}

func NewTwitterListSource(tw *core.Twitter, listName string, transformer *core.TweetTransformer) (*TwitterListSource, error) {
	ch := make(chan *core.Entry)
	sourceCh := make(chan *anaconda.Tweet)
	seeker := core.NewTweetSeeker(func(v url.Values) ([]anaconda.Tweet, error) {
		logrus.Infof("Twitter GetList: %s %v", listName, v)
		return tw.GetListTweets(listName, v)
	}, sourceCh)
	ret := &TwitterListSource{
		output:      ch,
		sourceCh:    sourceCh,
		twitter:     tw,
		listName:    listName,
		transformer: transformer,
		seeker:      seeker,
	}
	go ret.convertTweetPump()
	seeker.Start()
	return ret, nil
}

func (tl *TwitterListSource) convertTweetPump() {
	for !tl.stopped {
		tweet := <-tl.sourceCh
		entry, ok, err := tl.transformer.ToEntry(tweet)
		if err != nil {
			logrus.Errorf("Error while converting tweet: %s", err)
		} else {
			if ok {
				tl.output <- entry
			} else {
				logrus.Debugf("Skipped: %d @%s %s", tweet.Id, tweet.User.ScreenName, tweet.Text)
			}
		}
	}
}

func (tl *TwitterListSource) Faucet() <-chan *core.Entry {
	return tl.output
}

func (tl *TwitterListSource) Drain(count int) {
	tl.seeker.Drain(count)
}

func (tl *TwitterListSource) Shutdown() {
	tl.stopped = true
	tl.seeker.Stop()
	close(tl.output)
}
