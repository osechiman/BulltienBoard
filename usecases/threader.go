package usecases

import (
	"vspro/entities"
	"vspro/entities/valueobjects"
)

// ThreadRepositorer は外部データソースに存在するentities.Threadを操作する際に利用するインターフェースです。
type ThreadRepositorer interface {
	// GetThreadByID は指定されたvalueobjects.ThreadIDを持つentities.Threadを取得します。
	GetThreadByID(ID valueobjects.ThreadID) (*entities.Thread, error)
	// ListThread はentities.Threadの一覧を取得します。
	ListThread() ([]*entities.Thread, error)
	// ListThreadByBulletinBoardID は指定されたvalueobjects.BulletinBoardIDを持つentities.Threadの一覧を取得します。
	ListThreadByBulletinBoardID(bID valueobjects.BulletinBoardID) ([]*entities.Thread, error)
	// AddThread はentities.Threadを追加します。
	AddThread(t entities.Thread) error
}
