package controllers

import (
	"vspro/entities"
	"vspro/entities/valueobjects"
	"vspro/usecases"

	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	Repository usecases.QuestionRepositorer
}

type PostQuestion struct {
	ID         entities.QuestionID
	Difficulty int
	Answer     string
	Text       string
}

func NewQuestionController(r usecases.QuestionRepositorer) *QuestionController {
	return &QuestionController{Repository: r}
}

func (qc *QuestionController) GetQuestionByID(ID string) (*entities.Question, error) {
	qu := usecases.NewQuestionUsecase(qc.Repository)

	qid, err := convertIDToQuestionID(ID)
	if err != nil {
		return nil, err
	}

	return qu.GetQuestionByID(qid)
}

// コマンド・クエリの原則からは外れるがAPIのレスポンスに登録したデータを返却するためにエンティティを返す
func (qc *QuestionController) AddQuestion(c *gin.Context) (*entities.Question, error) {
	pq := PostQuestion{}
	err := c.BindJSON(&pq)
	if err != nil {
		return nil, err
	}
	qid, err := valueobjects.NewQuestionID("")
	if err != nil {
		return nil, err
	}
	q := entities.NewQuestion(qid, pq.Difficulty, pq.Answer, pq.Text)

	qu := usecases.NewQuestionUsecase(qc.Repository)
	return &q, qu.AddQuestion(q)
}

func (qc *QuestionController) ListQuestion() ([]*entities.Question, error) {
	qu := usecases.NewQuestionUsecase(qc.Repository)
	return qu.ListQuestion()
}

func (qc *QuestionController) DeleteQuestionByID(ID string) error {
	qu := usecases.NewQuestionUsecase(qc.Repository)

	qid, err := convertIDToQuestionID(ID)
	if err != nil {
		return err
	}

	return qu.DeleteQuestionByID(qid)
}

func convertIDToQuestionID(ID string) (entities.QuestionID, error) {
	return valueobjects.NewQuestionID(ID)
}

func AnswerQuestion() {}
