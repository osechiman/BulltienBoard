package gateways

import (
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
)

var inMemoryRepository = NewInMemoryRepository()
var bulletinBoards = make(map[valueobjects.BulletinBoardID]*entities.BulletinBoard)
var threads = make(map[valueobjects.ThreadID]*entities.Thread)
var comments = make(map[valueobjects.CommentID]*entities.Comment)

type InMemoryRepository struct{}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func GetInMemoryRepositoryInstance() *InMemoryRepository {
	return inMemoryRepository
}

func (i *InMemoryRepository) GetBulletinBoardByID(ID entities.BulletinBoardID) (*entities.BulletinBoard, error) {
	_, exist := bulletinBoards[ID.Get()]
	if !exist {
		return nil, errorobjects.NewNotFoundError(ID.String())
	}
	return bulletinBoards[ID.Get()], nil
}

func (i *InMemoryRepository) AddBulletinBoard(b entities.BulletinBoard) error {
	bulletinBoards[b.ID.Get()] = &b
	return nil
}

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

func (i *InMemoryRepository) ListThreadByBulletinBoardID(bID entities.BulletinBoardID) ([]*entities.Thread, error) {
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

func (i *InMemoryRepository) GetThreadByID(ID entities.ThreadID) (*entities.Thread, error) {
	_, exist := threads[ID.Get()]
	if !exist {
		return nil, errorobjects.NewNotFoundError(ID.String())
	}
	return threads[ID.Get()], nil
}

func (i *InMemoryRepository) AddThread(t entities.Thread) error {
	threads[t.ID.Get()] = &t
	return nil
}

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

func (i *InMemoryRepository) ListCommentByThreadID(tID entities.ThreadID) ([]*entities.Comment, error) {
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

func (i *InMemoryRepository) AddComment(c entities.Comment) error {
	comments[c.ID.Get()] = &c
	return nil
}
