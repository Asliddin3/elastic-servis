package main

import (
	"github.com/Asliddin3/poll-servis/config"
	"github.com/Asliddin3/poll-servis/internal/app"
)

func main() {
	cfg := config.LoadConfig()
	app.Run(cfg)
}
