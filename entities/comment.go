package entities

import (
	"vspro/entities/valueobjects"
)

type Comment struct {
	ID       CommentID
	ThreadID ThreadID
	Text     string
	CreateAt CommentTime
}

func NewComment(ID CommentID, tID ThreadID, text string, cTime CommentTime) Comment {
	return Comment{
		ID:       ID,
		ThreadID: tID,
		Text:     text,
		CreateAt: cTime,
	}
}

type CommentID interface {
	Get() valueobjects.CommentID
	String() string
}

type CommentTime interface {
	Get() valueobjects.CommentTime
	ToUnixTime() int64
	Equals(ct valueobjects.CommentTime) bool
}
