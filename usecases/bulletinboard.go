package usecases

import (
	"bulltienboard/entities"
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"
)

// BulletinBoardLimit は作成出来るBulletinBoardの上限値です。
const BulletinBoardLimit = 50

// BulletinBoardUsecase はBulletinBoardに対するUsecaseを定義するものです。
type BulletinBoardUsecase struct {
	BulletinBoardRepository BulletinBoardRepositorer // BulletinBoardRepository は外部データソースに存在するentities.BulletinBoardを操作する際に利用するインターフェースです。
	ThreadRepository        ThreadRepositorer        // ThreadRepositorer は外部データソースに存在するentities.Threadを操作する際に利用するインターフェースです。
}

// NewBulletinBoardUsecase はBulletinBoardUsecaseを初期化します。
func NewBulletinBoardUsecase(br BulletinBoardRepositorer, tr ThreadRepositorer) *BulletinBoardUsecase {
	return &BulletinBoardUsecase{BulletinBoardRepository: br, ThreadRepository: tr}
}

// GetBulletinBoardByID は指定されたvalueobjects.BulletinBoardIDを持つentities.BulletinBoardを取得します。
func (bbu *BulletinBoardUsecase) GetBulletinBoardByID(ID valueobjects.BulletinBoardID) (entities.BulletinBoard, error) {
	tl, err := bbu.ThreadRepository.ListThreadByBulletinBoardID(ID)
	if err != nil {
		switch err.(type) {
		case *errorobjects.NotFoundError:
			tl = make([]entities.Thread, 0)
		default:
			return entities.BulletinBoard{}, errorobjects.NewInternalServerError(err.Error())
		}
	}

	b, err := bbu.BulletinBoardRepository.GetBulletinBoardByID(ID)
	if err != nil {
		return entities.BulletinBoard{}, err
	}

	b.Threads = tl
	return b, nil
}

// AddBulletinBoard は現在登録されているBulletinBoardの数を確認して、閾値に達成していなければentities.BulletinBoardを追加します。
func (bbu *BulletinBoardUsecase) AddBulletinBoard(bb entities.BulletinBoard) error {
	bbs, err := bbu.BulletinBoardRepository.ListBulletinBoard()
	if err != nil {
		switch err.(type) {
		// NotFoundErrorの場合は処理を中断しない
		case *errorobjects.NotFoundError:
		default:
			return err
		}
	}

	if len(bbs) >= BulletinBoardLimit {
		return errorobjects.NewResourceLimitedError("maximum number of bulletin boards exceeded")
	}

	return bbu.BulletinBoardRepository.AddBulletinBoard(bb)
}

// ListBulletinBoard はentities.BulletinBoardの一覧を取得します。
func (bbu *BulletinBoardUsecase) ListBulletinBoard() ([]entities.BulletinBoard, error) {
	return bbu.BulletinBoardRepository.ListBulletinBoard()
}
