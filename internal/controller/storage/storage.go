package storage

import (
	"database/sql"

	repo "github.com/Asliddin3/elastic-servis/internal/controller/storage/repo"
	"github.com/rs/zerolog"

	// "github.com/Asliddin3/elastic-servis/internal/controller/storage/mongo"
	dbmongo "github.com/Asliddin3/elastic-servis/internal/controller/storage/mongo"
	dbpostgres "github.com/Asliddin3/elastic-servis/internal/controller/storage/postgres"

	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Poll() repo.PollStorageI
	Post() repo.PostStorageI
}

type StoragePg struct {
	Db       *mongo.Client
	Logger   *zerolog.Logger
	PollRepo repo.PollStorageI
	PostRepo repo.PostStorageI
}

func NewStoragePg(db *mongo.Client, postDb *sql.DB, lz *zerolog.Logger) *StoragePg {
	return &StoragePg{
		Db:       db,
		Logger:   lz,
		PollRepo: dbmongo.NewPollRepo(db),
		PostRepo: dbpostgres.NewPostRepo(postDb, lz),
	}
}

func (s StoragePg) Poll() repo.PollStorageI {
	return s.PollRepo
}

func (s StoragePg) Post() repo.PostStorageI {
	return s.PostRepo
}
