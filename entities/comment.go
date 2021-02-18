package entities

import (
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"

	"github.com/go-playground/validator"
)

// Comment はCommentのエンティティです。
type Comment struct {
	ID       CommentIDer  // ID はCommentIDインターフェースです。
	ThreadID ThreadIDer   // ThreadIDer はThreadIDインターフェースです。
	Text     string       `validate:"min=1,max=2048"` // Text はCommentの内容です。
	CreateAt CommentTimer // CreatAt はCommentが作成された時間です。
}

// NewComment はCommentを初期化します。
func NewComment(ID CommentIDer, tID ThreadIDer, text string, cTime CommentTimer) (Comment, error) {
	c := Comment{
		ID:       ID,
		ThreadID: tID,
		Text:     text,
		CreateAt: cTime,
	}
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return Comment{}, errorobjects.NewParameterBindingError(err.Error())
	}
	return c, nil
}

// CommentIDer はCommentエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type CommentIDer interface {
	// Get は自分自身を返却します。
	Get() valueobjects.CommentID
	// String はCommentIDが文字列に変換されたものを返却します。
	String() string
}

// CommentTimer はCommentエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type CommentTimer interface {
	Get() valueobjects.CommentTime
	ToUnixTime() int64
	Equals(ct valueobjects.CommentTime) bool
}
