package usecases

import (
	"vspro/entities"
)

// ThreadRepositorer はThreadRepositoryのインターフェースです。
type ThreadRepositorer interface {
	GetThreadByID(ID entities.ThreadID) (*entities.Thread, error)
	ListThread() ([]*entities.Thread, error)
	ListThreadByBulletinBoardID(bID entities.BulletinBoardID) ([]*entities.Thread, error)
	AddThread(t entities.Thread) error
}
