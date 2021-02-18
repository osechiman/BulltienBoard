package entities

import (
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"

	"github.com/go-playground/validator"
)

// Thread はThreadのエンティティです。
type Thread struct {
	ID              ThreadIDer        // ID はThreadIDインターフェースです。
	BulletinBoardID BulletinBoardIDer // BulletinBoardIDer はBulletinBoardIDインターフェースです。
	Title           string            `validate:"required,min=1,max=50"` // Title はThreadのTitleです。
	Comments        []Comment         // Comments はCommentの一覧です。
}

// NewThread はThreadを初期化します。
func NewThread(ID ThreadIDer, bID BulletinBoardIDer, title string) (Thread, error) {
	t := Thread{
		ID:              ID,
		BulletinBoardID: bID,
		Title:           title,
	}

	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return Thread{}, errorobjects.NewParameterBindingError(err.Error())
	}
	return t, nil
}

// ThreadIDer はThreadエンティティに実装されるインターフェースを定義しています。
// valueobjectsはこのインターフェースを満たす様に実装する必要があります。
type ThreadIDer interface {
	// Get は自分自身を返却します。
	Get() valueobjects.ThreadID
	// String はThreadIDが文字列に変換されたものを返却します。
	String() string
	// Equals は自分自身と引数に渡された値オブジェクトが同一のものか判定します。
	Equals(id valueobjects.ThreadID) bool
}
