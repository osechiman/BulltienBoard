package entities

import (
	"vspro/entities/valueobjects"
)

// Comment はCommentのエンティティです。
type Comment struct {
	ID       CommentID   // ID はCommentIDインターフェースです。
	ThreadID ThreadID    // ThreadID はThreadIDインターフェースです。
	Text     string      // Text はCommentの内容です。
	CreateAt CommentTime // CreatAt はCommentが作成された時間です。
}

// NewComment はCommentを初期化します。
func NewComment(ID CommentID, tID ThreadID, text string, cTime CommentTime) Comment {
	return Comment{
		ID:       ID,
		ThreadID: tID,
		Text:     text,
		CreateAt: cTime,
	}
}

// CommentID はCommentエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type CommentID interface {
	// Get は自分自身を返却します。
	Get() valueobjects.CommentID
	// String はCommentIDが文字列に変換されたものを返却します。
	String() string
}

// CommentTime はCommentエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type CommentTime interface {
	Get() valueobjects.CommentTime
	ToUnixTime() int64
	Equals(ct valueobjects.CommentTime) bool
}
