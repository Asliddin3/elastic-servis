package service

import (
	"github.com/Asliddin3/poll-servis/internal/controller/storage"
	"github.com/Asliddin3/poll-servis/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type PollService struct {
	Logger  *logger.Logger
	Storage storage.IStorage
}

func NewPollService(l *logger.Logger, stg *mongo.Client) *PollService {
	return &PollService{
		Logger:  l,
		Storage: storage.NewStoragePg(stg),
	}
}
