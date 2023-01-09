package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "github.com/Asliddin3/elastic-servis/configs"

	"github.com/Asliddin3/elastic-servis/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/Asliddin3/elastic-servis/pkg/db"

	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

func Run(cfg *config.Config) {
	// l := logger.New(cfg.LogLevel)
	// mongoDbUrl := fmt.Sprintf("%s://%s:%s@%s:%d/?maxPoolSize=20&w=majority", &cfg.MONGO_DATABASE,
	// 	&cfg.MONGO_USER, &cfg.MONGO_PASSWORD, &cfg.MONGO_HOST, &cfg.MONGO_PORT)
	_, err := db.ConnectToDb(cfg)
	fmt.Println(err)
	if err != nil {
		return
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://%s:3000", cfg)},
		AllowCredentials: true,
	})
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	cfg.Logger = logger
	logger.Info().Interface("config", cfg).Msg("config:")
	graphResolver := graph.Init(cfg)
	// pollService := service.NewPollService(l, mongoConnect)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Service: graphResolver}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://%s:%s/ for GraphQL playground", cfg.POLLServiceHost, cfg.POLLServicePort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.POLLServicePort), nil))

}
