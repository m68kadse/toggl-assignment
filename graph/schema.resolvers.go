package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"strconv"

	"github.com/m68kadse/toggl-assignment/dao"
	"github.com/m68kadse/toggl-assignment/graph/model"
)

// CreateQuestion is the resolver for the createQuestion field.
func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	question, err := r.QuestionDAO.CreateQuestion(ctx, model.DTOFromQuestionInput(&input))
	if err != nil {
		ctx.Err()
		return nil, err
	}

	return model.QuestionFromDTO(question), nil
}

// UpdateQuestion is the resolver for the updateQuestion field.
func (r *mutationResolver) UpdateQuestion(ctx context.Context, id string, input model.QuestionInput) (*model.Question, error) {
	q := model.DTOFromQuestionInput(&input)
	var err error
	q.ID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	q, err = r.QuestionDAO.UpdateQuestion(ctx, q)
	if err != nil {
		return nil, err
	}

	return model.QuestionFromDTO(q), nil
}

// DeleteQuestion is the resolver for the deleteQuestion field.
func (r *mutationResolver) DeleteQuestion(ctx context.Context, id string) (*string, error) {
	var err error
	qID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	qID, err = r.QuestionDAO.DeleteQuestion(ctx, qID)
	id = strconv.FormatInt(qID, 10)

	return &id, err
}

// Questions is the resolver for the questions field.
func (r *queryResolver) Questions(ctx context.Context, offset *int) ([]*model.Question, error) {
	const PAGESIZE int = 30

	params := dao.PaginationParams{
		Limit: PAGESIZE,
	}

	if offset != nil {
		params.Offset = *offset
	}

	questions, err := r.QuestionDAO.GetQuestions(ctx, params)
	if err != nil {
		return nil, err
	}

	modelQuestions := make([]*model.Question, len(questions))
	for i, q := range questions {
		modelQuestions[i] = model.QuestionFromDTO(q)
	}

	return modelQuestions, nil
}

// Question is the resolver for the question field.
func (r *queryResolver) Question(ctx context.Context, id string) (*model.Question, error) {
	intID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Invalid id: %s\n%w", id, err)
	}
	q, err := r.QuestionDAO.GetQuestionByID(ctx, intID)
	if err != nil {
		return nil, err
	}

	return model.QuestionFromDTO(q), nil
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
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
const PAGESIZE int = 30
