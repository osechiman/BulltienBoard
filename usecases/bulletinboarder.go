package usecases

import (
	"bulltienboard/entities"
	"bulltienboard/entities/valueobjects"
)

// BulletinBoardRepositorer は外部データソースに存在するentities.BulletinBoardを操作する際に利用するインターフェースです。
type BulletinBoardRepositorer interface {
	// GetBulletinBoardByID は指定されたvalueobjects.BulletinBoardIDを持つentities.BulletinBoardを取得します。
	GetBulletinBoardByID(ID valueobjects.BulletinBoardID) (entities.BulletinBoard, error)
	// ListBulletinBoard はentities.BulletinBoardの一覧を取得します。
	ListBulletinBoard() ([]entities.BulletinBoard, error)
	// AddBulletinBoard はentities.BulletinBoardを追加します。
	AddBulletinBoard(bb entities.BulletinBoard) error
	// DeleteBulletinBoard はentities.BulletinBoardを全て削除します。
	DeleteBulletinBoardAll() error
}
