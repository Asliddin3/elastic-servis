package repo

import (
	// "github.com/Asliddin3/graph/model"
	"github.com/Asliddin3/poll-servis/graph/model"
)

type PollStorageI interface {
	CreatePoll(*model.NewPoll) (*model.Poll, error)
	ChoiceFromPoll(*model.UserChoice) (*model.Poll, error)
	GetPoll(id *string) (*model.Poll, error)
	GetPolls() ([]*model.Poll, error)
}
