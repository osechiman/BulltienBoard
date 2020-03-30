package usecases

import (
	"vspro/entities"
)

type CommentUsecase struct {
	Repository CommentRepositorer
}

func NewCommentUsecase(r CommentRepositorer) *CommentUsecase {
	return &CommentUsecase{Repository: r}
}

func (cc *CommentUsecase) AddComment(c entities.Comment, threadRepository ThreadRepositorer) error {
	_, err := threadRepository.GetThreadByID(c.ThreadID)
	if err != nil {
		return err
	}
	return cc.Repository.AddComment(c)
}

func (cc *CommentUsecase) ListComment() ([]*entities.Comment, error) {
	return cc.Repository.ListComment()
}

func (cc *CommentUsecase) ListCommentByThreadID(tID entities.ThreadID) ([]*entities.Comment, error) {
	return cc.Repository.ListCommentByThreadID(tID)
}
