package usecases

import (
	"vspro/entities"
)

type BulletinBoardRepositorer interface {
	GetBulletinBoardByID(ID entities.BulletinBoardID) (*entities.BulletinBoard, error)
	ListBulletinBoard() ([]*entities.BulletinBoard, error)
	AddBulletinBoard(q entities.BulletinBoard) error
}
