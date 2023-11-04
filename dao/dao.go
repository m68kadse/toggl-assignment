package dao

import (
	"context"

	"github.com/m68kadse/toggl-assignment/dto"
)

type PaginationParams struct {
	Limit  int // Number of items per page
	Offset int // Starting index for pagination
}

type QuestionDAO interface {
	CreateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error)
	UpdateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error)
	DeleteQuestion(ctx context.Context, id int64) (int64, error)
	GetQuestionByID(ctx context.Context, id int64) (*dto.Question, error)
	GetQuestions(ctx context.Context, paginationParams PaginationParams) ([]*dto.Question, error)
}
