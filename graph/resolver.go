package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"fmt"

	"github.com/Asliddin3/poll-servis/pkg/logger"

	"github.com/Asliddin3/poll-servis/pkg/db"

	"github.com/Asliddin3/poll-servis/config"
	"github.com/Asliddin3/poll-servis/internal/controller/service"
)

type Resolver struct {
	service *service.PollService
}

func (r *Resolver) Connect(cfg *config.Config) *service.PollService {
	l := logger.New(cfg.LogLevel)
	con, err := db.ConnectToDb(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("error connecting to mongo in resolver", err.Error()))
	}
	pollService := service.NewPollService(l, con)
	return pollService
}
