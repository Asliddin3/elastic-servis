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
	Service *service.PollService
}

func Connect(cfg *config.Config) *service.PollService {
	l := logger.New(cfg.LogLevel)
	con, err := db.ConnectToDb(cfg)
	fmt.Println("-------con", err)
	if err != nil {
		var someErr error
		fmt.Errorf("error connecting to mongo in resolver", someErr)
		l.Fatal(someErr, err.Error())
	}
	pollService := service.NewPollService(l, con)
	return pollService
}
