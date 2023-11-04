package dao

type PaginationParams struct {
	Limit  int // Number of items per page
	Offset int // Starting index for pagination
}

type QuestionDao interface {
}
