package usecases

import (
	"bulltienboard/entities"
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"
)

const ThreadLimit = 50

// ThreadUsecase はThreadに対するUsecaseを定義するものです。
type ThreadUsecase struct {
	ThreadRepository        ThreadRepositorer        // ThreadRepositorer は外部データソースに存在するentities.Threadを操作する際に利用するインターフェースです。
	BulletinBoardRepository BulletinBoardRepositorer // BulletinBoardRepository は外部データソースに存在するentities.BulletinBoardを操作する際に利用するインターフェースです。
	CommentRepository       CommentRepositorer       // CommentRepositorer は外部データソースに存在するentities.Commentを操作する際に利用するインターフェースです。
}

// NewThreadUsecase はThreadUsecaseを初期化します。
func NewThreadUsecase(tr ThreadRepositorer, br BulletinBoardRepositorer, cr CommentRepositorer) *ThreadUsecase {
	return &ThreadUsecase{ThreadRepository: tr, BulletinBoardRepository: br, CommentRepository: cr}
}

// GetThreadByID は指定されたvalueobjects.ThreadIDを持つentities.Threadを取得します。
func (tu *ThreadUsecase) GetThreadByID(ID valueobjects.ThreadID) (entities.Thread, error) {
	cl, err := tu.CommentRepository.ListCommentByThreadID(ID)
	if err != nil {
		switch err.(type) {
		case *errorobjects.NotFoundError:
			cl = make([]entities.Comment, 0)
		default:
			return entities.Thread{}, err
		}
	}

	t, err := tu.ThreadRepository.GetThreadByID(ID)
	if err != nil {
		return entities.Thread{}, err
	}

	t.Comments = cl
	return t, nil
}

// AddThread は自信が持つBulletinBoardIDのBulletinBoardが存在するかをチェックし、
// 現在登録されているThreadの数を確認して閾値に達成していなければentities.Threadを追加します。
func (tu *ThreadUsecase) AddThread(t entities.Thread) error {
	_, err := tu.BulletinBoardRepository.GetBulletinBoardByID(t.BulletinBoardID.Get())
	if err != nil {
		return err
	}

	ts, err := tu.ThreadRepository.ListThread()
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

	return tu.ThreadRepository.AddThread(t)
}

// ListThread はentities.Threadの一覧を取得します。
func (tu *ThreadUsecase) ListThread() ([]entities.Thread, error) {
	return tu.ThreadRepository.ListThread()
}

// ListThreadByBulletinBoardID は指定されたvalueobjects.BulletinBoardIDを持つentities.Threadの一覧を取得します。
func (tu *ThreadUsecase) ListThreadByBulletinBoardID(bID valueobjects.BulletinBoardID) ([]entities.Thread, error) {
	return tu.ThreadRepository.ListThreadByBulletinBoardID(bID)
}
