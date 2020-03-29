package presenters

import (
	"net/http"
	"vspro/entities"
)

type QuestionPresenter struct{}

func NewQuestionPresenter() *QuestionPresenter {
	return &QuestionPresenter{}
}

type Questions []*Question

type Question struct {
	ID         string
	Difficulty int
	Answer     string
	Text       string
}

func (qp *QuestionPresenter) ConvertToHttpDeleteQuestionResponse(httpStatusCode int, qid string) *HTTPResponse {
	message := qid + " has been deleted."
	return newHTTPSuccessResponse(httpStatusCode, http.StatusText(httpStatusCode), message)
}

func (qp *QuestionPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (qp *QuestionPresenter) ConvertToHttpQuestionListResponse(ql []*entities.Question) *HTTPResponse {
	res := Questions{}
	for _, q := range ql {
		res = append(res, convertEntitiesQuestionToQuestion(q))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (qp *QuestionPresenter) ConvertToHttpQuestionResponse(q *entities.Question) *HTTPResponse {
	res := Questions{}
	pq := convertEntitiesQuestionToQuestion(q)
	res = append(res, pq)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesQuestionToQuestion(q *entities.Question) *Question {
	pq := Question{
		ID:         q.ID.String(),
		Difficulty: q.Difficulty,
		Answer:     q.Answer,
		Text:       q.Text,
	}
	return &pq
}
