package valueobjects

// ValueObjectは外部ライブラリに依存して良いというポリシーで運用。
import (
	"vspro/entities/errorobjects"

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
			return BulletinBoardID{}, errorobjects.NewInternalServerError(err.Error())
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

func (bb BulletinBoardID) Equals(id BulletinBoardID) bool {
	if (bb.Get() == id.Get()) && (bb.String() == id.String()) {
		return true
	}
	return false
}
