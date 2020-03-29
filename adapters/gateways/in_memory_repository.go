package gateways

import (
	"vspro/entities"
	"vspro/entities/valueobjects"
)

var inMemoryRepository = NewInMemoryRepository()
var questions = make(map[entities.QuestionID]*entities.Question)
var bulletinBoards = make(map[entities.BulletinBoardID]*entities.BulletinBoard)
var threads = make(map[entities.ThreadID]*entities.Thread)

type InMemoryRepository struct{}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

func GetInMemoryRepositoryInstance() *InMemoryRepository {
	return inMemoryRepository
}

func (i *InMemoryRepository) GetQuestionByID(ID entities.QuestionID) (*entities.Question, error) {
	_, exist := questions[ID.Get()]
	if !exist {
		return nil, valueobjects.NewNotFoundError(ID.String())
	}
	return questions[ID.Get()], nil
}

func (i *InMemoryRepository) ListQuestion() ([]*entities.Question, error) {
	var qs []*entities.Question
	if len(questions) == 0 {
		return nil, valueobjects.NewNotFoundError("question not registered,")
	}
	for _, v := range questions {
		qs = append(qs, v)
	}
	return qs, nil
}

func (i *InMemoryRepository) AddQuestion(q entities.Question) error {
	questions[q.ID.Get()] = &q
	return nil
}

func (i *InMemoryRepository) AnswerQuestion() (bool, error) {
	return true, nil
}

func (i *InMemoryRepository) DeleteQuestionByID(ID entities.QuestionID) error {
	_, exist := questions[ID.Get()]
	if !exist {
		return valueobjects.NewNotFoundError(ID.String())
	}
	delete(questions, ID.Get())
	return nil
}

func (i *InMemoryRepository) GetBulletinBoardByID(ID entities.BulletinBoardID) (*entities.BulletinBoard, error) {
	_, exist := bulletinBoards[ID.Get()]
	if !exist {
		return nil, valueobjects.NewNotFoundError(ID.String())
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
		return nil, valueobjects.NewNotFoundError("bulletinBoard not registered,")
	}
	for _, v := range bulletinBoards {
		bs = append(bs, v)
	}
	return bs, nil
}

func (i *InMemoryRepository) ListThread() ([]*entities.Thread, error) {
	var ts []*entities.Thread
	if len(threads) == 0 {
		return nil, valueobjects.NewNotFoundError("thread not registered,")
	}
	for _, v := range threads {
		ts = append(ts, v)
	}
	return ts, nil
}

func (i *InMemoryRepository) ListThreadByBulletinBoard(bID entities.BulletinBoardID) ([]*entities.Thread, error) {
	var ts []*entities.Thread
	if len(threads) == 0 {
		return nil, valueobjects.NewNotFoundError("thread not registered,")
	}

	for _, v := range threads {
		if v.BulletinBoardID == bID.Get() {
			ts = append(ts, v)
		}
	}

	if len(ts) == 0 {
		return nil, valueobjects.NewNotFoundError(bID.String() + " associated with thread")
	}

	return ts, nil
}

func (i *InMemoryRepository) GetThreadByID(ID entities.ThreadID) (*entities.Thread, error) {
	_, exist := threads[ID.Get()]
	if !exist {
		return nil, valueobjects.NewNotFoundError(ID.String())
	}
	return threads[ID.Get()], nil
}

func (i *InMemoryRepository) AddThread(t entities.Thread) error {
	threads[t.ID.Get()] = &t
	return nil
}
