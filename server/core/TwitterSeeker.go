package core

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"net/url"
	"strconv"
	"time"
)

type TweetSeeker struct {
	fetchFunc func(v url.Values) ([]anaconda.Tweet, error)
	minId     int64 // most oldest ID // TODO: should be big.Int?
	maxId     int64 // most recent ID
	output    chan *anaconda.Tweet
	controlCh ControlChannel
	canceled  bool
}

func NewTweetSeeker(fetchFunc func(v url.Values) ([]anaconda.Tweet, error), output chan *anaconda.Tweet) *TweetSeeker {
	return &TweetSeeker{
		output:    output,
		controlCh: make(ControlChannel),
		fetchFunc: fetchFunc,
	}
}

func (ts *TweetSeeker) Stop() {
	if ts.canceled {
		return
	}
	ts.canceled = true
	close(ts.controlCh)
}

func (ts *TweetSeeker) Start() {
	go ts.runMessagePump()
	go ts.runTicker()
}

func (ts *TweetSeeker) updateMinMaxID(tweet *anaconda.Tweet) {
	if tweet.Id > ts.maxId || ts.minId == 0 {
		ts.maxId = tweet.Id
	}
	if tweet.Id < ts.minId || ts.minId == 0 {
		ts.minId = tweet.Id
	}
}

func (ts *TweetSeeker) runMessagePump() {
	for !ts.canceled {
		msg := <-ts.controlCh
		if ts.canceled {
			return
		}
		switch msg {
		case STREAM_TICK:
			err := ts.Refresh()
			if err != nil {
				logrus.Errorf("TweetSeeker Stopped: %s", err)
				return
			}
		default:
			if msg.IsDrainRequest() {
				res, err := ts.Previous(msg.AsDrainRequest())
				if err != nil {
					logrus.Errorf("TweetSeeker Stopped: %s", err)
					return
				}
				for _, tweet := range res {
					ts.updateMinMaxID(&tweet)
					ts.output <- &tweet
				}
			} else {
				logrus.Errorf("Unknown channel control message: %v", msg)
			}
		}
	}
}

func (ts *TweetSeeker) runTicker() {
	// Future: intelligent ticker
	// Now: periodic ticker
	interval := 60 * time.Second
	ticker := time.NewTicker(interval)
	for !ts.canceled {
		ts.controlCh <- STREAM_TICK
		<-ticker.C
	}
}

func (ts *TweetSeeker) Previous(count int) ([]anaconda.Tweet, error) {
	vals := url.Values{}
	vals.Set("count", strconv.Itoa(count))
	if ts.minId != 0 {
		vals.Set("max_id", strconv.FormatInt(ts.minId, 10))
	}
	result, err := ts.fetchFunc(vals)

	return result, err
}

func (ts *TweetSeeker) Drain(count int) {
	if !ts.canceled {
		ts.controlCh <- StreamControlMessage(count)
	}
}

func (ts *TweetSeeker) Refresh() error {
	count := 20
	vals := url.Values{}
	tryCount := 0
	maxTry := 3
	for tryCount < maxTry {
		if ts.maxId != 0 {
			count = 20
			vals.Set("since_id", strconv.FormatInt(ts.maxId, 10))
		} else {
			count = 100 // initial fetch
		}
		vals.Set("count", strconv.Itoa(count))
		result, err := ts.fetchFunc(vals)
		if err != nil {
			return err
		}
		for i := len(result) - 1; i >= 0; i-- {
			tweet := &result[i]
			ts.updateMinMaxID(tweet)
			ts.output <- tweet
		}
		if len(result) < count {
			break
		}
		tryCount += 1
	}
	return nil
}
