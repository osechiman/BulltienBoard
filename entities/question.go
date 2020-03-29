package entities

import (
	"vspro/entities/valueobjects"
)

type Question struct {
	ID         QuestionID
	Difficulty int
	Answer     string
	Text       string
}

func NewQuestion(ID QuestionID, difficulty int, answer string, text string) Question {
	return Question{ID: ID, Difficulty: difficulty, Answer: answer, Text: text}
}

type QuestionID interface {
	Get() valueobjects.QuestionID
	String() string
}
