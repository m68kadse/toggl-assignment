package graph

import "github.com/m68kadse/toggl-assignment/dao"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	QuestionDAO dao.QuestionDAO
}
