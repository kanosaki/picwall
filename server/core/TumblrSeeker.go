package core

import (
	"github.com/Sirupsen/logrus"
	"github.com/kanosaki/picwall/server/tumblr"
	"strconv"
	"time"
)

type TumblrSeeker struct {
	fetchFunc func(v map[string]string) ([]tumblr.Post, error)
	minId     int64 // most oldest ID // TODO: should be big.Int?
	maxId     int64 // most recent ID
	output    chan *tumblr.Post
	controlCh ControlChannel
	canceled  bool
}

func NewTumblrSeeker(fetchFunc func(v map[string]string) ([]tumblr.Post, error), output chan *tumblr.Post) *TumblrSeeker {
	return &TumblrSeeker{
		output:    output,
		controlCh: make(ControlChannel),
		fetchFunc: fetchFunc,
	}
}

func (ts *TumblrSeeker) Stop() {
	if ts.canceled {
		return
	}
	ts.canceled = true
	close(ts.controlCh)
}

func (ts *TumblrSeeker) Start() {
	go ts.runMessagePump()
	go ts.runTicker()
}

func (ts *TumblrSeeker) updateMinMaxID(post *tumblr.Post) {
	if post.ID > ts.maxId || ts.minId == 0 {
		ts.maxId = post.ID
	}
	if post.ID < ts.minId || ts.minId == 0 {
		ts.minId = post.ID
	}
}

func (ts *TumblrSeeker) runMessagePump() {
	for !ts.canceled {
		msg := <-ts.controlCh
		if ts.canceled {
			return
		}
		switch msg {
		case STREAM_TICK:
			err := ts.Refresh()
			if err != nil {
				logrus.Errorf("TumblrSeeker Stopped: %s", err)
				return
			}
		default:
		}
	}
}

func (ts *TumblrSeeker) runTicker() {
	// Future: intelligent ticker
	// Now: periodic ticker
	interval := 60 * time.Second
	ticker := time.NewTicker(interval)
	for !ts.canceled {
		ts.controlCh <- STREAM_TICK
		<-ticker.C
	}
}

//func (ts *TumblrSeeker) Previous(count int) ([]anaconda.Tweet, error) {
//	vals := url.Values{}
//	vals.Set("count", strconv.Itoa(count))
//	if ts.minId != 0 {
//		vals.Set("max_id", strconv.FormatInt(ts.minId, 10))
//	}
//	result, err := ts.fetchFunc(vals)
//
//return result, err
//}

func (ts *TumblrSeeker) Drain(count int) {
	if !ts.canceled {
		ts.controlCh <- StreamControlMessage(count)
	}
}

func (ts *TumblrSeeker) Refresh() error {
	count := 20
	vals := make(map[string]string)
	tryCount := 0
	maxTry := 3
	for tryCount < maxTry {
		if ts.maxId != 0 {
			count = 20
			vals["since_id"] = strconv.FormatInt(ts.maxId, 10)
		} else {
			count = 100 // initial fetch
		}
		vals["limit"] = strconv.Itoa(count)
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
