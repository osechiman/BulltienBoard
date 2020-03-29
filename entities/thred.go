package entities

import (
	"vspro/entities/valueobjects"
)

type Thread struct {
	ID              ThreadID
	BulletinBoardID BulletinBoardID
	Title           string
	Comment         []*Comment
}

func NewThread(ID ThreadID, bID BulletinBoardID, title string) Thread {
	return Thread{ID: ID, BulletinBoardID: bID, Title: title}
}

type ThreadID interface {
	Get() valueobjects.ThreadID
	String() string
}
