package main

type App struct {
	Port          string
	AdminAddress  string
	DeployAddress string
}

type Config struct {
}

func NewApp() *App {
	return &App{}
}
