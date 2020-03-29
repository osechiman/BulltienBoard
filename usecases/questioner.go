package usecases

import (
	"vspro/entities"
)

type QuestionRepositorer interface {
	GetQuestionByID(ID entities.QuestionID) (*entities.Question, error)
	ListQuestion() ([]*entities.Question, error)
	AddQuestion(q entities.Question) error
	DeleteQuestionByID(ID entities.QuestionID) error
	AnswerQuestion() (bool, error)
}
