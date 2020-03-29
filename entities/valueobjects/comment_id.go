package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用
import (
	"github.com/google/uuid"
)

type commentID uuid.UUID

type CommentID struct {
	id  commentID
	str string
}

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

func (c CommentID) Get() CommentID {
	return c
}

func (c CommentID) String() string {
	return c.str
}

func (c CommentID) Equals(other CommentID) bool {
	if (c.Get() == other.Get()) && (c.String() == other.String()) {
		return true
	}
	return false
}
