package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"strconv"
	"github.com/jaysonmulwa/sheng-app/graph/generated"
	"github.com/jaysonmulwa/sheng-app/graph/model"
	"github.com/jaysonmulwa/sheng-app/server/graph/repository"
)
var wordRepo repository.WordRepository = repository.New()

func (r *mutationResolver) CreateWord(ctx context.Context, input model.NewWord) (*model.Word, error) {
	word := &model.Word{
		ID:      strconv.Itoa(rand.Int()),
		Word:    input.Word,
		Meaning: input.Meaning,
		Author:  &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	// r.words = append(r.words, word)

	wordRepo.Save(word)

	return word, nil
}

func (r *queryResolver) Words(ctx context.Context) ([]*model.Word, error) {
	//return r.words, nil
	return wordRepo.FindAll(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
