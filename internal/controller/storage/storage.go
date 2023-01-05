package storage

import (
	repo "github.com/Asliddin3/poll-servis/internal/controller/storage/repo"
	// "github.com/Asliddin3/poll-servis/internal/controller/storage/mongo"
	dbmongo "github.com/Asliddin3/poll-servis/internal/controller/storage/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Poll() repo.PollStorageI
}

type StoragePg struct {
	Db       *mongo.Client
	PollRepo repo.PollStorageI
}

func NewStoragePg(db *mongo.Client) *StoragePg {
	return &StoragePg{
		Db:       db,
		PollRepo: dbmongo.NewPollRepo(db),
	}
}

func (s StoragePg) Poll() repo.PollStorageI {
	return s.PollRepo
}
