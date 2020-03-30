package usecases

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
)

type ThreadUsecase struct {
	Repository ThreadRepositorer
}

func NewThreadUsecase(r ThreadRepositorer) *ThreadUsecase {
	return &ThreadUsecase{Repository: r}
}

func (tu *ThreadUsecase) GetThreadByID(ID entities.ThreadID, commentRepository CommentRepositorer) (*entities.Thread, error) {
	cl, err := commentRepository.ListCommentByThreadID(ID)
	if err != nil {
		switch err.(type) {
		case *errorobjects.NotFoundError:
			cl = make([]*entities.Comment, 0)
		default:
			return nil, err
		}
	}

	t, err := tu.Repository.GetThreadByID(ID)
	if err != nil {
		return nil, err
	}

	t.Comments = cl
	return t, nil
}

func (tu *ThreadUsecase) AddThread(t entities.Thread, bulletinBoardRepository BulletinBoardRepositorer) error {
	_, err := bulletinBoardRepository.GetBulletinBoardByID(t.BulletinBoardID)
	if err != nil {
		return err
	}
	return tu.Repository.AddThread(t)
}

func (tu *ThreadUsecase) ListThread() ([]*entities.Thread, error) {
	return tu.Repository.ListThread()
}

func (tu *ThreadUsecase) ListThreadByBulletinBoardID(bID entities.BulletinBoardID) ([]*entities.Thread, error) {
	return tu.Repository.ListThreadByBulletinBoardID(bID)
}
