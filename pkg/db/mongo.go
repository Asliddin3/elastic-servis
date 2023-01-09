package db

import (
	"context"
	"database/sql"
	"fmt"

	config "github.com/Asliddin3/elastic-servis/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb(cfg *config.Config) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", cfg.MONGOHost, cfg.MONGOPort)))
}

func ConnectToPostgres(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.POSTGRES_HOST, cfg.POSTGRES_PORT, cfg.POSTGRES_USER, cfg.POSTGRES_PASS, cfg.POSTGRES_DB)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		cfg.Logger.Err(err).AnErr("error connecting to postgres", err)
		return db, err
	}
	return db, nil
}
