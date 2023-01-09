package repo

import (
	// "github.com/Asliddin3/graph/model"
	"github.com/Asliddin3/elastic-servis/graph/model"
)

type PollStorageI interface {
	CreatePoll(*model.NewPoll) (*model.Poll, error)
	ChoiceFromPoll(*model.UserChoice) (*model.Poll, error)
	GetPoll(id *string) (*model.Poll, error)
	GetPolls() ([]*model.Poll, error)
}
type PostStorageI interface {
	CreatePost(*model.NewPost) (*model.Post, error)
	UpdatePost(*model.UpdatedPost) (*model.Post, error)
	GetPost(id int) (*model.Post, error)
	GetPosts() ([]*model.Post, error)
}
