package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用
import (
	"github.com/google/uuid"
)

type threadID uuid.UUID

type ThreadID struct {
	id  threadID
	str string
}

func NewThreadID(ID string) (ThreadID, error) {
	uid, err := uuid.Parse(ID)
	if err != nil {
		uid, err = uuid.NewRandom()
		if err != nil {
			return ThreadID{}, err
		}
	}

	tid := ThreadID{id: threadID(uid), str: uid.String()}
	return tid, nil
}

func (t ThreadID) Get() ThreadID {
	return t
}

func (t ThreadID) String() string {
	return t.str
}

func (t ThreadID) Equals(other ThreadID) bool {
	if (t.Get() == other.Get()) && (t.String() == other.String()) {
		return true
	}
	return false
}
