package service

import (
	"database/sql"

	"github.com/Asliddin3/elastic-servis/internal/controller/storage"
	"github.com/Asliddin3/elastic-servis/pkg/logger"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type PollService struct {
	Logstash *zerolog.Logger
	Logger   *logger.Logger
	Storage  storage.IStorage
}

func NewPollService(l *logger.Logger, lz *zerolog.Logger, stg *mongo.Client, postDb *sql.DB) *PollService {
	return &PollService{
		Logger:   l,
		Logstash: lz,
		Storage:  storage.NewStoragePg(stg, postDb, lz),
	}
}
