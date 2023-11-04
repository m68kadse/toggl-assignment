package dto

type Question struct {
	ID      int64
	Body    string
	Options []*Option
}
