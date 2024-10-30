package internal

import "go.uber.org/zap"

var (
	log = zap.Logger{}
)

type App struct {
}

func New(logger zap.Logger) *App {
	log = logger
	return &App{}
}
