package entities

import (
	"vspro/entities/valueobjects"
)

type BulletinBoard struct {
	ID      BulletinBoardID
	Title   string
	Threads []*Thread
}

func NewBulletinBoard(ID BulletinBoardID, title string) BulletinBoard {
	return BulletinBoard{ID: ID, Title: title}
}

type BulletinBoardID interface {
	Get() valueobjects.BulletinBoardID
	String() string
	Equals(id valueobjects.BulletinBoardID) bool
}
