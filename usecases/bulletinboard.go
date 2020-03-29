package usecases

import (
	"vspro/entities"
	"vspro/entities/valueobjects"
)

type BulletinBoardUsecase struct {
	Repository BulletinBoardRepositorer
}

func NewBulletinBoardUsecase(r BulletinBoardRepositorer) *BulletinBoardUsecase {
	return &BulletinBoardUsecase{Repository: r}
}

func (bbu *BulletinBoardUsecase) GetBulletinBoardByID(ID entities.BulletinBoardID, threadRepository ThreadRepositorer) (*entities.BulletinBoard, error) {
	tl, err := threadRepository.ListThreadByBulletinBoard(ID)
	if err != nil {
		switch err.(type) {
		case *valueobjects.NotFoundError:
			tl = make([]*entities.Thread, 0)
		default:
			return nil, err
		}
	}

	b, err := bbu.Repository.GetBulletinBoardByID(ID.Get())
	if err != nil {
		return nil, err
	}

	b.Threads = tl
	return b, nil
}

func (bbu *BulletinBoardUsecase) AddBulletinBoard(q entities.BulletinBoard) error {
	return bbu.Repository.AddBulletinBoard(q)
}

func (bbu *BulletinBoardUsecase) ListBulletinBoard() ([]*entities.BulletinBoard, error) {
	return bbu.Repository.ListBulletinBoard()
}
