package mongo

import (
	"context"
	"time"

	"github.com/Asliddin3/poll-servis/graph/model"

	// "github.com/Asliddin3/pkg/db"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PollRepo struct {
	Db *mongo.Client
}

func NewPollRepo(db *mongo.Client) *PollRepo {
	return &PollRepo{Db: db}
}

func (db *PollRepo) CreatePoll(pollReq *model.NewPoll) (*model.Poll, error) {
	collection := db.Db.Database("poll_service").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	poll := model.Poll{}
	poll.ID = uuid.New().String()
	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "id", Value: poll.ID},
		{Key: "email", Value: &pollReq.UserEmail},
		{Key: "text", Value: &pollReq.Text},
		{Key: "choises", Value: &pollReq.Choises},
	})
	for _, val := range pollReq.Choises {
		poll.Choises = append(poll.Choises, &model.Choice{
			ID:   val.ID,
			Name: val.Name,
		})
	}
	if err != nil {
		return &model.Poll{}, err
	}
	return &poll, nil
}

func (db *PollRepo) ChoiceFromPoll(choiceReq *model.UserChoice) (*model.Poll, error) {
	collection := db.Db.Database("poll_service").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	poll := &model.Poll{}
	_, err := collection.UpdateOne(ctx, choiceReq.PollID, bson.D{
		{"$push", bson.M{"results": bson.M{"choiceid": &choiceReq.ChoiceID, "email": &choiceReq.UserEmail}}},
	})
	if err != nil {
		return &model.Poll{}, err
	}
	filter := bson.D{{"id", &choiceReq.PollID}}
	err = collection.FindOne(ctx, filter).Decode(&poll)
	if err != nil {
		return &model.Poll{}, err
	}
	return poll, nil
}

// func (db *db.DB) FindMovieById(id string) *model.Movie {
// 	ObjectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	movieColl := db.client.Database("graphql-mongodb-api-db").Collection("movie")
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
// 	res := movieColl.FindOne(ctx, bson.M{"_id": ObjectID})

// 	movie := model.Movie{ID: id}

// 	res.Decode(&movie)

// 	return &movie
// }
