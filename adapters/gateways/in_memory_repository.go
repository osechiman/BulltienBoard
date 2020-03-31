package gateways

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
)

// inMemoryRepository はメモリ上でデータを管理する為の変数です。
var inMemoryRepository = NewInMemoryRepository()

// bulletinBoards はentities.BulletinBoardを管理する為の変数です。
var bulletinBoards = make(map[valueobjects.BulletinBoardID]*entities.BulletinBoard)

// threads はentities.Threadを管理する為の変数です。
var threads = make(map[valueobjects.ThreadID]*entities.Thread)

// comments はentities.Commentを管理する為の変数です。
var comments = make(map[valueobjects.CommentID]*entities.Comment)

// InMemoryRepository はメモリ上でデータ管理する為のStructです。
type InMemoryRepository struct{}

// NewInMemoryRepository はInMemoryRepositoryを初期化します。
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

// GetInMemoryRepositoryInstance はパッケージグローバルな変数からInMemoryRepositoryを取得します。
func GetInMemoryRepositoryInstance() *InMemoryRepository {
	return inMemoryRepository
}

// GetBulletinBoardByID は指定されたvalueobjects.BulletinBoardIDを元にentities.BulletinBoardを取得します。
func (i *InMemoryRepository) GetBulletinBoardByID(ID valueobjects.BulletinBoardID) (*entities.BulletinBoard, error) {
	_, exist := bulletinBoards[ID.Get()]
	if !exist {
		return nil, errorobjects.NewNotFoundError(ID.String())
	}
	return bulletinBoards[ID.Get()], nil
}

// AddBulletinBoard はentities.BulletinBoardを追加します。
func (i *InMemoryRepository) AddBulletinBoard(b entities.BulletinBoard) error {
	bulletinBoards[b.ID.Get()] = &b
	return nil
}

// ListBulletinBoard はentities.BulletinBoardの一覧を取得します。
func (i *InMemoryRepository) ListBulletinBoard() ([]*entities.BulletinBoard, error) {
	var bs []*entities.BulletinBoard
	if len(bulletinBoards) == 0 {
		return nil, errorobjects.NewNotFoundError("bulletinBoard not registered,")
	}
	for _, v := range bulletinBoards {
		bs = append(bs, v)
	}
	return bs, nil
}

// ListThread はentities.Threadの一覧を取得します。
func (i *InMemoryRepository) ListThread() ([]*entities.Thread, error) {
	var ts []*entities.Thread
	if len(threads) == 0 {
		return nil, errorobjects.NewNotFoundError("thread not registered,")
	}
	for _, v := range threads {
		ts = append(ts, v)
	}
	return ts, nil
}

// ListThreadByBulletinBoardID は指定されたvalueobjects.BulletinBoardIDを持つentities.Threadの一覧を取得します。
func (i *InMemoryRepository) ListThreadByBulletinBoardID(bID valueobjects.BulletinBoardID) ([]*entities.Thread, error) {
	var ts []*entities.Thread
	if len(threads) == 0 {
		return nil, errorobjects.NewNotFoundError("thread not registered,")
	}

	for _, v := range threads {
		if bID.Equals(v.BulletinBoardID.Get()) {
			ts = append(ts, v)
		}
	}

	if len(ts) == 0 {
		return nil, errorobjects.NewNotFoundError(bID.String() + " associated with thread")
	}

	return ts, nil
}

// GetThreadByID は指定されたvalueobjects.ThreadIDを持つentities.Threadを取得します。
func (i *InMemoryRepository) GetThreadByID(ID valueobjects.ThreadID) (*entities.Thread, error) {
	_, exist := threads[ID.Get()]
	if !exist {
		return nil, errorobjects.NewNotFoundError(ID.String())
	}
	return threads[ID.Get()], nil
}

// AddThread はentities.Threadを追加します。
func (i *InMemoryRepository) AddThread(t entities.Thread) error {
	threads[t.ID.Get()] = &t
	return nil
}

// ListComment はentities.Commentの一覧を取得します。
func (i *InMemoryRepository) ListComment() ([]*entities.Comment, error) {
	var cs []*entities.Comment
	if len(comments) == 0 {
		return nil, errorobjects.NewNotFoundError("comment not registered,")
	}
	for _, v := range comments {
		cs = append(cs, v)
	}
	return cs, nil
}

// ListCommentByThreadID は指定されたvalueobjects.ThreadIDを元にentities.Commentの一覧を取得します。
func (i *InMemoryRepository) ListCommentByThreadID(tID valueobjects.ThreadID) ([]*entities.Comment, error) {
	var cs []*entities.Comment
	if len(comments) == 0 {
		return nil, errorobjects.NewNotFoundError("comment not registered,")
	}

	for _, v := range comments {
		if tID.Equals(v.ThreadID.Get()) {
			cs = append(cs, v)
		}
	}

	if len(cs) == 0 {
		return nil, errorobjects.NewNotFoundError(tID.String() + " associated with thread")
	}

	return cs, nil
}

// AddComment はentities.Comment を追加します。
func (i *InMemoryRepository) AddComment(c entities.Comment) error {
	comments[c.ID.Get()] = &c
	return nil
}
