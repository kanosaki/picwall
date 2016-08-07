package core

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"net/url"
)

type Twitter struct {
	Api     *anaconda.TwitterApi
	me      anaconda.User
	myLists map[string]anaconda.List
}

func NewTwitter(config *TwitterConfig) (*Twitter, error) {
	// init twitter
	anaconda.SetConsumerKey(config.ConsumerKey)
	anaconda.SetConsumerSecret(config.ConsumerSecret)
	twitterAccount := config.DefaultAccount
	api := anaconda.NewTwitterApi(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret)
	//api.EnableThrottling(time.Duration(config.InitialThrottling.Seconds) * time.Second, config.InitialThrottling.BucketSize)
	logrus.Infof("Checking Twitter...")
	ok, err := api.VerifyCredentials()
	if !ok {
		return nil, err
	}
	logrus.Infof("Twitter authorized")
	me, err := api.GetSelf(url.Values{})
	if err != nil {
		return nil, err
	}
	logrus.Infof("Twitter logeed in as @%s", me.ScreenName)
	myLists, err := api.GetListsOwnedBy(me.Id, nil)
	if err != nil {
		return nil, err
	}
	logrus.Infof("Twitter you(@%s) have %d lists (%v)", me.ScreenName, len(myLists))
	myListsMap := make(map[string]anaconda.List)
	for i, l := range myLists {
		logrus.Infof("Twitter list %d: %s(%s)", i, l.Name, l.FullName)
		myListsMap[l.Name] = l
	}
	return &Twitter{
		myLists: myListsMap,
		Api:     api,
		me:      me,
	}, nil
}

func (t *Twitter) GetListTweets(listName string, v url.Values) ([]anaconda.Tweet, error) {
	if id, ok := t.myLists[listName]; ok {
		tweets, err := t.Api.GetListTweets(id.Id, true, v)
		return tweets, err
	} else {
		return nil, fmt.Errorf("No such list: %s", listName)
	}
}

type TweetTransformer struct {
	twitter *Twitter
}

func NewTweetTransformer(twitter *Twitter) *TweetTransformer {
	return &TweetTransformer{
		twitter: twitter,
	}
}

func (tt *TweetTransformer) ToEntry(tweet *anaconda.Tweet) (*Entry, bool, error) {
	if len(tweet.Entities.Media) > 0 {
		e := NewEntry()
		e.Set("text", tweet.Text)
		e.Set("id", fmt.Sprintf("twitter-tweet-%d", tweet.Id))
		for _, media := range tweet.Entities.Media {
			if media.Type != "photo" {
				continue
			}
			e.Set("thumbnail", fmt.Sprintf("%s:thumb", media.Media_url))
		}
		return e, true, nil
	} else {
		return nil, false, nil
	}
}
