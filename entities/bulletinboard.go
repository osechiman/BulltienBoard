package entities

import (
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"

	"github.com/go-playground/validator"
)

// BulletinBoard はBulletinBoardのエンティティです。
type BulletinBoard struct {
	ID      BulletinBoardIDer // ID はBulletinBoardIDインターフェースです。
	Title   string            `validate:"required,min=1,max=50"` // Title はBulletinBoardのTitleです。
	Threads []*Thread         // Thread はThreadの一覧です。
}

// NewBulletinBoard はBulletinBoardを初期化します。
func NewBulletinBoard(ID BulletinBoardIDer, title string) (BulletinBoard, error) {
	bb := BulletinBoard{
		ID:    ID,
		Title: title,
	}
	validate := validator.New()
	err := validate.Struct(bb)
	if err != nil {
		return BulletinBoard{}, errorobjects.NewParameterBindingError(err.Error())
	}
	return BulletinBoard{ID: ID, Title: title}, nil
}

// BulletinBoardIDer はBulletinBoardエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type BulletinBoardIDer interface {
	// Get は自分自身を返却します。
	Get() valueobjects.BulletinBoardID
	// String はBulletinBoardIDが文字列に変換されたものを返却します。
	String() string
	// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
	Equals(id valueobjects.BulletinBoardID) bool
}
