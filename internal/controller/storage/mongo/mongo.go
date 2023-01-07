package mongo

import (
	"context"
	"fmt"
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
	for _, val := range pollReq.Choises {
		poll.Choises = append(poll.Choises, &model.Choice{
			ID:   uuid.New().String(),
			Name: val.Name,
		})
	}
	_, err := collection.InsertOne(ctx, bson.D{
		{Key: "id", Value: poll.ID},
		{Key: "email", Value: &pollReq.Email},
		{Key: "text", Value: &pollReq.Text},
		{Key: "choises", Value: &poll.Choises},
		// {Key: "results", Value: []map[string]string{}},
	})
	fmt.Println("in createPoll mongo func", err)
	poll.Text = pollReq.Text
	poll.Email = pollReq.Email
	fmt.Println("after for mongo func")
	return &poll, nil
}

type ObjectForDecode struct {
	id      string
	email   string
	text    string
	choises Choises
	results []Result
}
type Result struct {
	email    string
	choiceid string
}
type Choises struct {
	email    string
	choiceid string
}

func (db *PollRepo) ChoiceFromPoll(choiceReq *model.UserChoice) (*model.Poll, error) {
	collection := db.Db.Database("poll_service").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	params := make(map[string]string)
	params["choiceid"] = choiceReq.ChoiceID
	params["email"] = choiceReq.UserEmail
	// model.PollResult{}
	filterSearch := bson.D{{Key: "id", Value: choiceReq.PollID}}
	// update := bson.D{{"$push", bson.D{{"results", params}}}}
	defer cancel()
	var poll model.Poll
	_, err := collection.UpdateOne(ctx, filterSearch, bson.D{
		{Key: "$push", Value: bson.M{"results": bson.M{"choiceid": choiceReq.ChoiceID, "email": choiceReq.UserEmail}}},
		// update[0],
		//  {Key: "$push", Value:bson.D{"results": bson.D{"choiceid", &choiceReq.UserEmail}}},
	})
	fmt.Println("first quiery err", err)
	if err != nil {
		return &model.Poll{}, err
	}
	filter := bson.D{{"id", &choiceReq.PollID}}
	testObj := &ObjectForDecode{}
	err = collection.FindOne(ctx, filter).Decode(&poll)
	fmt.Println("sercond quiery err", err)
	fmt.Println(testObj)

	if err != nil {
		return &model.Poll{}, err
	}
	return &poll, nil
}
func (db *PollRepo) GetPoll(id *string) (*model.Poll, error) {
	collection := db.Db.Database("poll_service").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	poll := &model.Poll{}
	filter := bson.D{{"id", &id}}
	err := collection.FindOne(ctx, filter).Decode(poll)
	if err != nil {
		return &model.Poll{}, err
	}
	return poll, nil
}

func (db *PollRepo) GetPolls() ([]*model.Poll, error) {
	collection := db.Db.Database("poll_service").Collection("polls")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	polls := []*model.Poll{}
	filter := bson.D{{}}
	// opts := options.Find()
	// opts.SetSort(bson.D{{"duration", -1}})
	res, err := collection.Find(ctx, filter)
	fmt.Println("somer error here",err)
	for res.Next(ctx) {
		poll := &model.Poll{}
		err = res.Decode(&poll)
		if err != nil {
			return nil, err
		}
		polls = append(polls, poll)
	}
	if err != nil {
		return []*model.Poll{}, err
	}
	return polls, nil
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
