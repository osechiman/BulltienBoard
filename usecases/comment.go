package usecases

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
)

const CommentLimit = 1000

// CommentUsecase はCommentに対するUsecaseを定義するものです。
type CommentUsecase struct {
	Repository CommentRepositorer // Repositorer は外部データソースに存在するentities.Commentを操作する際に利用するインターフェースです。
}

// NewCommentUsecase はCommentUsecaseを初期化します。
func NewCommentUsecase(r CommentRepositorer) *CommentUsecase {
	return &CommentUsecase{Repository: r}
}

// AddComment は自信が持つThreadIDのThreadが存在するかをチェックし、
// 現在登録されているCommentの数を確認して閾値に達成していなければentities.Commentを追加します。
func (cc *CommentUsecase) AddComment(c entities.Comment, threadRepository ThreadRepositorer) error {
	_, err := threadRepository.GetThreadByID(c.ThreadID.Get())
	if err != nil {
		return err
	}

	cs, err := cc.Repository.ListComment()
	if err != nil {
		return err
	}

	if len(cs) > CommentLimit {
		return errorobjects.NewResourceLimitedError("maximum number of comment exceeded. comment limit is " + string(CommentLimit))
	}

	return cc.Repository.AddComment(c)
}

// ListComment はentities.Commentの一覧を取得します。
func (cc *CommentUsecase) ListComment() ([]entities.Comment, error) {
	return cc.Repository.ListComment()
}

// ListCommentByThreadID は指定されたvalueobjects.ThreadIDを持つentities.Commentの一覧を取得します。
func (cc *CommentUsecase) ListCommentByThreadID(tID valueobjects.ThreadID) ([]entities.Comment, error) {
	return cc.Repository.ListCommentByThreadID(tID)
}
