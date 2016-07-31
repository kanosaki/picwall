package core

import (
	"github.com/gorilla/websocket"
)

type App struct {
	WebSocketUpgrader *websocket.Upgrader
	center            *Center
}

func NewApp(config *Config) *App {
	return &App{
		center: &Center{},
		WebSocketUpgrader: websocket.Upgrader{
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
	}
}

