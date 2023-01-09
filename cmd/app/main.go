package main

import (
	_ "github.com/99designs/gqlgen"
	config "github.com/Asliddin3/elastic-servis/configs"
	"github.com/Asliddin3/elastic-servis/internal/app"
)

func main() {
	cfg := config.LoadConfig()
	app.Run(cfg)
}
