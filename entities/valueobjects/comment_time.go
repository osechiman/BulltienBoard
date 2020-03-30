package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用
import (
	"time"
)

type CommentTime struct {
	unixTime int64
}

func NewCommentTime(unixTime int64) (CommentTime, error) {
	ct := CommentTime{
		unixTime: unixTime,
	}
	if unixTime == 0 {
		ct.unixTime = time.Now().Unix()
	}
	return ct, nil
}

func (c CommentTime) Get() CommentTime {
	return c
}

func (c CommentTime) ToUnixTime() int64 {
	return c.unixTime
}

func (c CommentTime) Equals(other CommentTime) bool {
	if c.Get() == other.Get() {
		return true
	}
	return false
}
