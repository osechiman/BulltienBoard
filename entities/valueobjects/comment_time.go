package valueobjects

import (
	"time"
)

// CommentTime はコメントが作成された時刻を表します。
type CommentTime struct {
	unixTime int64 // unixTime はunixTimeです。
}

// NewCommentTime はCommentの作成時刻を生成します。
func NewCommentTime(unixTime int64) (CommentTime, error) {
	ct := CommentTime{
		unixTime: unixTime,
	}
	if unixTime < 0 {
		ct.unixTime = time.Now().Unix()
	}
	return ct, nil
}

// Get は自分自身を返却します。
func (c CommentTime) Get() CommentTime {
	return c
}

// ToUnixTime はCommentTimeがunixTimeに変換されたものを返却します。
func (c CommentTime) ToUnixTime() int64 {
	return c.unixTime
}

// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
func (c CommentTime) Equals(other CommentTime) bool {
	if c.Get() == other.Get() {
		return true
	}
	return false
}
