package entities

import (
	"vspro/entities/valueobjects"
)

// Thread はThreadのエンティティです。
type Thread struct {
	ID              ThreadID        // ID はThreadIDインターフェースです。
	BulletinBoardID BulletinBoardID // BulletinBoardID はBulletinBoardIDインターフェースです。
	Title           string          // Title はThreadのTitleです。
	Comments        []*Comment      // Comments はCommentの一覧です。
}

// NewThread はThreadを初期化します。
func NewThread(ID ThreadID, bID BulletinBoardID, title string) Thread {
	return Thread{ID: ID, BulletinBoardID: bID, Title: title}
}

// ThreadID はThreadエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type ThreadID interface {
	// Get は自分自身を返却します。
	Get() valueobjects.ThreadID
	// String はThreadIDが文字列に変換されたものを返却します。
	String() string
	// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
	Equals(id valueobjects.ThreadID) bool
}
