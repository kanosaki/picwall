package source

import (
	"github.com/Sirupsen/logrus"
	"github.com/kanosaki/picwall/server/core"
	"reflect"
)

type BottomSource struct {
	upstreams  []core.Source
	output     chan *core.Entry
	bufferSize int
}

func NewBottomSource(upstreams []core.Source, bufferSize int) *BottomSource {
	ret := &BottomSource{
		output:     make(chan *core.Entry, bufferSize),
		upstreams:  upstreams,
		bufferSize: bufferSize,
	}
	go ret.runPump()
	return ret
}

func (bs *BottomSource) runPump() {
	cases := make([]reflect.SelectCase, len(bs.upstreams))
	for i, src := range bs.upstreams {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(src.Faucet())}
	}
	for {
		_, value, ok := reflect.Select(cases)
		if !ok {
			// channel closed
			continue
		}
		// src := bs.upstreams[chosen]
		val, ok := value.Interface().(*core.Entry)
		if ok {
			bs.output <- val
		} else {
			logrus.Errorf("Invalid value: %v", val)
		}
	}
}

func (bs *BottomSource) Faucet() <-chan *core.Entry {
	return bs.output
}
