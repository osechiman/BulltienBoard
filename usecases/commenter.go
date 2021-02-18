package usecases

import (
	"bulltienboard/entities"
	"bulltienboard/entities/valueobjects"
)

// CommentRepositorer は外部データソースに存在するentities.Commentを操作する際に利用するインターフェースです。
type CommentRepositorer interface {
	// ListComment はentities.Commentの一覧を取得します。
	ListComment() ([]entities.Comment, error)
	// ListCommentByThreadID は指定されたvalueobjects.ThreadIDを持つentities.Commentの一覧を取得します。
	ListCommentByThreadID(tID valueobjects.ThreadID) ([]entities.Comment, error)
	// AddComment はentities.Comment を追加します。
	AddComment(c entities.Comment) error
}
