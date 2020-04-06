package usecases

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
)

// BulletinBoardUsecase はBulletinBoardに対するUsecaseを定義するものです。
type BulletinBoardUsecase struct {
	Repository BulletinBoardRepositorer // Repository は外部データソースに存在するentities.BulletinBoardを操作する際に利用するインターフェースです。
}

// NewBulletinBoardUsecase はBulletinBoardUsecaseを初期化します。
func NewBulletinBoardUsecase(r BulletinBoardRepositorer) *BulletinBoardUsecase {
	return &BulletinBoardUsecase{Repository: r}
}

// GetBulletinBoardByID は指定されたvalueobjects.BulletinBoardIDを持つentities.BulletinBoardを取得します。
func (bbu *BulletinBoardUsecase) GetBulletinBoardByID(ID valueobjects.BulletinBoardID, threadRepository ThreadRepositorer) (entities.BulletinBoard, error) {
	tl, err := threadRepository.ListThreadByBulletinBoardID(ID)
	if err != nil {
		switch err.(type) {
		case *errorobjects.NotFoundError:
			tl = make([]entities.Thread, 0)
		default:
			return entities.BulletinBoard{}, errorobjects.NewInternalServerError(err.Error())
		}
	}

	b, err := bbu.Repository.GetBulletinBoardByID(ID)
	if err != nil {
		return entities.BulletinBoard{}, err
	}

	b.Threads = tl
	return b, nil
}

// AddBulletinBoard はentities.BulletinBoardを追加します。
func (bbu *BulletinBoardUsecase) AddBulletinBoard(bb entities.BulletinBoard) error {
	return bbu.Repository.AddBulletinBoard(bb)
}

// ListBulletinBoard はentities.BulletinBoardの一覧を取得します。
func (bbu *BulletinBoardUsecase) ListBulletinBoard() ([]entities.BulletinBoard, error) {
	return bbu.Repository.ListBulletinBoard()
}
