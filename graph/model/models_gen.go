// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Option struct {
	ID      string `json:"id"`
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type OptionInput struct {
	Body    string `json:"body"`
	Correct bool   `json:"correct"`
}

type Question struct {
	ID      string    `json:"id"`
	Body    string    `json:"body"`
	Options []*Option `json:"options"`
}

type QuestionInput struct {
	Body    string         `json:"body"`
	Options []*OptionInput `json:"options"`
}