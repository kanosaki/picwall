package core

import (
	"github.com/gorilla/websocket"
)

type App struct {
	WebSocketUpgrader *websocket.Upgrader
	center            *Center
	config            *Config
	Twitter           *Twitter
	Tumblr            *Tumblr
}

func NewApp(config *Config) (*App, error) {
	twitter, err := NewTwitter(config.Twitter)
	if err != nil {
		return nil, err
	}
	tumblr, err := NewTumblr(config.Tumblr)
	if err != nil {
		return nil, err
	}
	return &App{
		Tumblr:  tumblr,
		Twitter: twitter,
		config:  config,
		center:  &Center{},
		WebSocketUpgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}, nil
}
