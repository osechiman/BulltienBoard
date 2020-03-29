package usecases

import (
	"vspro/adapters/gateways"
	"vspro/entities"
)

type ThreadUsecase struct {
	Repository ThreadRepositorer
}

func NewThreadUsecase(r ThreadRepositorer) *ThreadUsecase {
	return &ThreadUsecase{Repository: r}
}

func (tu *ThreadUsecase) GetThreadByID(ID entities.ThreadID) (*entities.Thread, error) {
	return tu.Repository.GetThreadByID(ID.Get())
}

func (tu *ThreadUsecase) AddThread(t entities.Thread) error {
	br := gateways.GetInMemoryRepositoryInstance()
	_, err := br.GetBulletinBoardByID(t.BulletinBoardID)
	if err != nil {
		return err
	}
	return tu.Repository.AddThread(t)
}

func (tu *ThreadUsecase) ListThread() ([]*entities.Thread, error) {
	return tu.Repository.ListThread()
}

func (tu *ThreadUsecase) ListThreadByBulletinBoardID(bID entities.BulletinBoardID) ([]*entities.Thread, error) {
	return tu.Repository.ListThreadByBulletinBoard(bID)
}
