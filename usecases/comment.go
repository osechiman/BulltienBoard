package usecases

import (
	"bulltienboard/entities"
	"bulltienboard/entities/errorobjects"
	"bulltienboard/entities/valueobjects"
)

// CommentLimit は作成出来るCommentの上限値です。
const CommentLimit = 1000

// CommentUsecase はCommentに対するUsecaseを定義するものです。
type CommentUsecase struct {
	CommentRepository CommentRepositorer // Repositorer は外部データソースに存在するentities.Commentを操作する際に利用するインターフェースです。
	ThreadRepository  ThreadRepositorer  // ThreadRepositorer は外部データソースに存在するentities.Threadを操作する際に利用するインターフェースです。
}

// NewCommentUsecase はCommentUsecaseを初期化します。
func NewCommentUsecase(cr CommentRepositorer, tr ThreadRepositorer) *CommentUsecase {
	return &CommentUsecase{CommentRepository: cr, ThreadRepository: tr}
}

// AddComment は自信が持つThreadIDのThreadが存在するかをチェックし、
// 現在登録されているCommentの数を確認して閾値に達成していなければentities.Commentを追加します。
func (cc *CommentUsecase) AddComment(c entities.Comment) error {
	_, err := cc.ThreadRepository.GetThreadByID(c.ThreadID.Get())
	if err != nil {
		return err
	}

	cs, err := cc.CommentRepository.ListComment()
	if err != nil {
		switch err.(type) {
		// AddCommentにおいては一覧が取得出来なくても登録できる仕様なのでNotFoundErrorは無視します。
		case *errorobjects.NotFoundError:
		default:
			return err
		}
	}

	if len(cs) >= CommentLimit {
		return errorobjects.NewResourceLimitedError("maximum number of comment exceeded. comment limit is " + string(CommentLimit))
	}

	return cc.CommentRepository.AddComment(c)
}

// ListComment はentities.Commentの一覧を取得します。
func (cc *CommentUsecase) ListComment() ([]entities.Comment, error) {
	return cc.CommentRepository.ListComment()
}

// ListCommentByThreadID は指定されたvalueobjects.ThreadIDを持つentities.Commentの一覧を取得します。
func (cc *CommentUsecase) ListCommentByThreadID(tID valueobjects.ThreadID) ([]entities.Comment, error) {
	return cc.CommentRepository.ListCommentByThreadID(tID)
}
