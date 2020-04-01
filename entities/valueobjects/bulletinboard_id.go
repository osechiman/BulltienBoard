package valueobjects

import (
	"vspro/entities/errorobjects"

	"github.com/google/uuid"
)

// bulletinBoardID はuuid.UUIDを独自のTypeに再定義したものです。
type bulletinBoardID uuid.UUID

// BulletinBoardID は掲示板のIDです。
type BulletinBoardID struct {
	id  bulletinBoardID // id はbulletinBoardIDです。
	str string          // str はuuid.UUIDを文字列に変換したものです。
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

// Get は自分自身を返却します。
func (bb BulletinBoardID) Get() BulletinBoardID {
	return bb
}

// String はBulletinBoardIDが文字列に変換されたものを返却します。
func (bb BulletinBoardID) String() string {
	return bb.str
}

// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
func (bb BulletinBoardID) Equals(other BulletinBoardID) bool {
	if (bb.Get() == other.Get()) && (bb.String() == other.String()) {
		return true
	}
	return false
}
