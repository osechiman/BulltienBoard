package entities

import (
	"vspro/entities/valueobjects"
)

type Comment struct {
	ID       CommentID
	ThreadID ThreadID
	Text     string
}

func NewComment(ID CommentID, text string) Comment {
	return Comment{ID: ID, Text: text}
}

type CommentID interface {
	Get() valueobjects.CommentID
	String() string
}
