package gateways

import (
	"errors"
	"vspro/entities"
)

var questions = make(map[entities.QuestionID]*entities.Question)

type InMemoryRepository struct{}

func (i *InMemoryRepository) GetQuestionByID(ID entities.QuestionID) (*entities.Question, error) {
	_, exist := questions[ID.Get()]
	if !exist {
		return nil, errors.New(ID.String() + " is not found.")
	}
	return questions[ID.Get()], nil
}

func (i *InMemoryRepository) ListQuestion() ([]*entities.Question, error) {
	var qs []*entities.Question
	if len(questions) == 0 {
		return nil, errors.New("question not registered, please register a question")
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
		return errors.New(ID.String() + " is not found.")
	}
	delete(questions, ID.Get())
	return nil
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}
