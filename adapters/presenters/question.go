package presenters

import (
	"net/http"
	"vspro/entities"
)

type QuestionPresenter struct {
	HTTPQuestionResponse HTTPQuestionResponse
}

func NewQuestionPresenter() *QuestionPresenter {
	return &QuestionPresenter{}
}

type HTTPQuestionResponse struct {
	Status  int
	Message string
	Data    interface{}
}

type Questions []*Question

type Question struct {
	ID         string
	Difficulty int
	Answer     string
	Text       string
}

func (qp *QuestionPresenter) ConvertToHttpDeleteQuestionResponse(httpStatusCode int, qid string) *HTTPQuestionResponse {
	message := qid + " has been deleted."
	return newHTTPSuccessResponse(httpStatusCode, http.StatusText(httpStatusCode), message)
}

func (qp *QuestionPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPQuestionResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (qp *QuestionPresenter) ConvertToHttpQuestionListResponse(ql []*entities.Question) *HTTPQuestionResponse {
	res := Questions{}
	for _, q := range ql {
		res = append(res, convertEntitiesQuestionToQuestion(q))
	}
	return newHTTPQuestionResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (qp *QuestionPresenter) ConvertToHttpQuestionResponse(q *entities.Question) *HTTPQuestionResponse {
	res := Questions{}
	pq := convertEntitiesQuestionToQuestion(q)
	res = append(res, pq)
	return newHTTPQuestionResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
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

func newHTTPQuestionResponse(status int, message string, data Questions) *HTTPQuestionResponse {
	return &HTTPQuestionResponse{Status: status, Message: message, Data: data}
}

func newHTTPSuccessResponse(status int, message string, i interface{}) *HTTPQuestionResponse {
	res := make([]interface{}, 0)
	res = append(res, i)
	return &HTTPQuestionResponse{Status: status, Message: message, Data: res}
}

func newHTTPErrorResponse(status int, message string, err error) *HTTPQuestionResponse {
	res := make([]string, 0)
	res = append(res, err.Error())
	return &HTTPQuestionResponse{Status: status, Message: message, Data: res}
}
