package usecases

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
)

const ThreadLimit = 50

// ThreadUsecase はThreadに対するUsecaseを定義するものです。
type ThreadUsecase struct {
	Repository ThreadRepositorer // Repositorer は外部データソースに存在するentities.Threadを操作する際に利用するインターフェースです。
}

// NewThreadUsecase はThreadUsecaseを初期化します。
func NewThreadUsecase(r ThreadRepositorer) *ThreadUsecase {
	return &ThreadUsecase{Repository: r}
}

// GetThreadByID は指定されたvalueobjects.ThreadIDを持つentities.Threadを取得します。
func (tu *ThreadUsecase) GetThreadByID(ID valueobjects.ThreadID, commentRepository CommentRepositorer) (entities.Thread, error) {
	cl, err := commentRepository.ListCommentByThreadID(ID)
	if err != nil {
		switch err.(type) {
		case *errorobjects.NotFoundError:
			cl = make([]entities.Comment, 0)
		default:
			return entities.Thread{}, err
		}
	}

	t, err := tu.Repository.GetThreadByID(ID)
	if err != nil {
		return entities.Thread{}, err
	}

	t.Comments = cl
	return t, nil
}

// AddThread は自信が持つBulletinBoardIDのBulletinBoardが存在するかをチェックし、
// 現在登録されているThreadの数を確認して閾値に達成していなければentities.Threadを追加します。
func (tu *ThreadUsecase) AddThread(t entities.Thread, bulletinBoardRepository BulletinBoardRepositorer) error {
	_, err := bulletinBoardRepository.GetBulletinBoardByID(t.BulletinBoardID.Get())
	if err != nil {
		return err
	}

	ts, err := tu.Repository.ListThread()
	if err != nil {
		switch err.(type) {
		// AddThreadにおいては一覧が取得出来なくても登録できる仕様なのでNotFoundErrorは無視します。
		case *errorobjects.NotFoundError:
		default:
			return err
		}
	}

	if len(ts) > ThreadLimit {
		return errorobjects.NewResourceLimitedError("maximum number of thread exceeded. thread limit is " + string(ThreadLimit))
	}

	return tu.Repository.AddThread(t)
}

// ListThread はentities.Threadの一覧を取得します。
func (tu *ThreadUsecase) ListThread() ([]entities.Thread, error) {
	return tu.Repository.ListThread()
}

// ListThreadByBulletinBoardID は指定されたvalueobjects.BulletinBoardIDを持つentities.Threadの一覧を取得します。
func (tu *ThreadUsecase) ListThreadByBulletinBoardID(bID valueobjects.BulletinBoardID) ([]entities.Thread, error) {
	return tu.Repository.ListThreadByBulletinBoardID(bID)
}
