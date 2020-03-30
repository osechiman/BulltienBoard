package usecases

import (
	"vspro/entities"
)

type CommentRepositorer interface {
	ListComment() ([]*entities.Comment, error)
	ListCommentByThreadID(tID entities.ThreadID) ([]*entities.Comment, error)
	AddComment(c entities.Comment) error
}
