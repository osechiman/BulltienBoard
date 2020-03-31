package valueobjects

import (
	"github.com/google/uuid"
)

// commentID はuuid.UUIDを独自のTypeに再定義したものです。
type commentID uuid.UUID

// CommentID はコメントのIDです。
type CommentID struct {
	id  commentID // id はcommentIDです。
	str string    // str はuuid.UUIDを文字列に変換したものです。
}

// NewCommentID はCommentのIDを生成します。
func NewCommentID(ID string) (CommentID, error) {
	uid, err := uuid.Parse(ID)
	if err != nil {
		uid, err = uuid.NewRandom()
		if err != nil {
			return CommentID{}, err
		}
	}

	cid := CommentID{id: commentID(uid), str: uid.String()}
	return cid, nil
}

// Get は自分自身を返却します。
func (c CommentID) Get() CommentID {
	return c
}

// String はCommentIDが文字列に変換されたものを返却します。
func (c CommentID) String() string {
	return c.str
}

// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
func (c CommentID) Equals(other CommentID) bool {
	if (c.Get() == other.Get()) && (c.String() == other.String()) {
		return true
	}
	return false
}
