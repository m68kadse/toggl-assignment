package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"

	"github.com/togglhire/backend-homework/graph/model"
)

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: CreateQuestion - createQuestion"))
}

// UpdateQuestion is the resolver for the updateQuestion field.
func (r *mutationResolver) UpdateQuestion(ctx context.Context, id string, input model.QuestionInput) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: UpdateQuestion - updateQuestion"))
}

// DeleteQuestion is the resolver for the deleteQuestion field.
func (r *mutationResolver) DeleteQuestion(ctx context.Context, id string) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: DeleteQuestion - deleteQuestion"))
}

// CreateOption is the resolver for the createOption field.
func (r *mutationResolver) CreateOption(ctx context.Context, input model.OptionInput) (*model.Option, error) {
	panic(fmt.Errorf("not implemented: CreateOption - createOption"))
}

// UpdateOption is the resolver for the updateOption field.
func (r *mutationResolver) UpdateOption(ctx context.Context, id string, input model.OptionInput) (*model.Option, error) {
	panic(fmt.Errorf("not implemented: UpdateOption - updateOption"))
}

// DeleteOption is the resolver for the deleteOption field.
func (r *mutationResolver) DeleteOption(ctx context.Context, id string) (*model.Option, error) {
	panic(fmt.Errorf("not implemented: DeleteOption - deleteOption"))
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	panic(fmt.Errorf("not implemented: Questions - questions"))
}

// Question is the resolver for the question field.
func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	panic(fmt.Errorf("not implemented: Question - question"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }