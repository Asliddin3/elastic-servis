package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
	"template/voice-servis/graph/model"
)

// CreatePoll is the resolver for the CreatePoll field.
func (r *mutationResolver) CreatePoll(ctx context.Context, input *model.NewPoll) (*model.Poll, error) {
	panic(fmt.Errorf("not implemented: CreatePoll - CreatePoll"))
}

// ChoiceFromPoll is the resolver for the ChoiceFromPoll field.
func (r *mutationResolver) ChoiceFromPoll(ctx context.Context, input *model.UserChoice) (*model.Poll, error) {
	panic(fmt.Errorf("not implemented: ChoiceFromPoll - ChoiceFromPoll"))
}

// Poll is the resolver for the poll field.
func (r *queryResolver) Poll(ctx context.Context, pollID string) (*model.Poll, error) {
	panic(fmt.Errorf("not implemented: Poll - poll"))
}

// Polls is the resolver for the polls field.
func (r *queryResolver) Polls(ctx context.Context) ([]*model.Poll, error) {
	panic(fmt.Errorf("not implemented: Polls - polls"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.