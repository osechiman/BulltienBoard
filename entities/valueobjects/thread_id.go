package valueobjects

import (
	"github.com/google/uuid"
)

// threadID はuuid.UUIDを独自のTypeに再定義したものです。
type threadID uuid.UUID

// ThreadID はスレッドのIDです。
type ThreadID struct {
	id  threadID // id はcommentIDです。
	str string   // str はuuid.UUIDを文字列に変換したものです。
}

// NewThreadID はThreadのIDを生成します。
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

// Get は自分自身を返却します。
func (t ThreadID) Get() ThreadID {
	return t
}

// String はThreadIDが文字列に変換されたものを返却します。
func (t ThreadID) String() string {
	return t.str
}

// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
func (t ThreadID) Equals(other ThreadID) bool {
	if (t.Get() == other.Get()) && (t.String() == other.String()) {
		return true
	}
	return false
}
