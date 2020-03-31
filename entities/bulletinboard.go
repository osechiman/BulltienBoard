package entities

import (
	"vspro/entities/valueobjects"
)

// BulletinBoard はBulletinBoardのエンティティです。
type BulletinBoard struct {
	ID      BulletinBoardID // ID はBulletinBoardIDインターフェースです。
	Title   string          // Title はBulletinBoardのTitleです。
	Threads []*Thread       // Thread はThreadの一覧です。
}

// NewBulletinBoard はBulletinBoardを初期化します。
func NewBulletinBoard(ID BulletinBoardID, title string) BulletinBoard {
	return BulletinBoard{ID: ID, Title: title}
}

// BulletinBoardID はBulletinBoardエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type BulletinBoardID interface {
	// Get は自分自身を返却します。
	Get() valueobjects.BulletinBoardID
	// String はBulletinBoardIDが文字列に変換されたものを返却します。
	String() string
	// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
	Equals(id valueobjects.BulletinBoardID) bool
}
