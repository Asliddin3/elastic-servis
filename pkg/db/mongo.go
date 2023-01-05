package db

import (
	"context"
	"fmt"

	"github.com/Asliddin3/poll-servis/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb(cfg *config.Config) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", &cfg.MONGOHost, &cfg.MONGOPort)))
}
