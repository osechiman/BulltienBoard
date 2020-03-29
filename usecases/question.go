package usecases

import (
	"vspro/entities"
)

type QuestionUsecase struct {
	Repository QuestionRepositorer
}

func NewQuestionUsecase(r QuestionRepositorer) *QuestionUsecase {
	return &QuestionUsecase{Repository: r}
}

func (qu *QuestionUsecase) GetQuestionByID(ID entities.QuestionID) (*entities.Question, error) {
	return qu.Repository.GetQuestionByID(ID.Get())
}

func (qu *QuestionUsecase) AddQuestion(q entities.Question) error {
	return qu.Repository.AddQuestion(q)
}

func (qu *QuestionUsecase) ListQuestion() ([]*entities.Question, error) {
	return qu.Repository.ListQuestion()
}

func (qu *QuestionUsecase) DeleteQuestionByID(ID entities.QuestionID) error {
	return qu.Repository.DeleteQuestionByID(ID.Get())
}

func AnswerQuestions(qr QuestionRepositorer) (bool, error) {
	return qr.AnswerQuestion()
}
