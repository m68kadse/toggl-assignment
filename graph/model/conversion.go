package model

import (
	"strconv"
	"strings"

	"github.com/m68kadse/toggl-assignment/dto"
)

// While keeping separate DTOs is not strictly necessary for the scope of this assignment
// it makes the codebase more maintainable insofar as the GraphQL model is not the only representation
// of our data and GraphQL can easily be replaced and amended with other APIs.

func OptionFromDTO(o *dto.Option) *Option {
	return &Option{
		Body:    strings.Clone(o.Body),
		Correct: o.Correct,
	}
}

func QuestionFromDTO(q *dto.Question) *Question {
	mq := new(Question)
	mq.ID = strconv.FormatInt(q.ID, 10)
	mq.Body = strings.Clone(q.Body)
	mq.Options = make([]*Option, len(q.Options))
	for i, o := range q.Options {
		mq.Options[i] = OptionFromDTO(o)
	}

	return mq
}

func DTOFromQuestionInput(input *QuestionInput) *dto.Question {
	questionDTO := &dto.Question{
		Body:    input.Body,
		Options: make([]*dto.Option, len(input.Options)),
	}

	for i, optInput := range input.Options {
		questionDTO.Options[i] = &dto.Option{
			Body:    optInput.Body,
			Correct: optInput.Correct,
		}
	}

	return questionDTO
}
