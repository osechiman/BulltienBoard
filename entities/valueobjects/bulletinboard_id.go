package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用。
import (
	"github.com/google/uuid"
)

type bulletinBoardID uuid.UUID

// BulletinBoardID は掲示板のIDです。
type BulletinBoardID struct {
	id  bulletinBoardID
	str string
}

// NewBulletinBoardID はBulletinBoardのIDを生成します。
func NewBulletinBoardID(ID string) (BulletinBoardID, error) {
	uid, err := uuid.Parse(ID)
	if err != nil {
		uid, err = uuid.NewRandom()
		if err != nil {
			return BulletinBoardID{}, NewInternalServerError(err.Error())
		}
	}

	bbid := BulletinBoardID{id: bulletinBoardID(uid), str: uid.String()}
	return bbid, nil
}

func (bb BulletinBoardID) Get() BulletinBoardID {
	return bb
}

func (bb BulletinBoardID) String() string {
	return bb.str
}

func (bb BulletinBoardID) Equals(other BulletinBoardID) bool {
	if (bb.Get() == other.Get()) && (bb.String() == other.String()) {
		return true
	}
	return false
}
