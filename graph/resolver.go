package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"fmt"

	"github.com/Asliddin3/elastic-servis/pkg/db"
	"github.com/Asliddin3/elastic-servis/pkg/logger"

	config "github.com/Asliddin3/elastic-servis/configs"
	"github.com/Asliddin3/elastic-servis/internal/controller/service"
)

type Resolver struct {
	Service *service.PollService
}

func Init(cfg *config.Config) *service.PollService {
	l := logger.New(cfg.LogLevel)
	con, err := db.ConnectToDb(cfg)
	if err != nil {
		var someErr error
		fmt.Errorf("error connecting to mongo in resolver", someErr)
		cfg.Logger.Fatal().AnErr(someErr.Error(), err)
		l.Fatal(someErr, err.Error())
	}
	postDb, err := db.ConnectToPostgres(cfg)
	if err != nil {
		var someErr error
		fmt.Errorf("error connecting to postgres in resolver", someErr)
		cfg.Logger.Fatal().AnErr(someErr.Error(), err)
		l.Fatal(someErr, err.Error())
	}
	pollService := service.NewPollService(l, &cfg.Logger, con, postDb)
	return pollService
}
