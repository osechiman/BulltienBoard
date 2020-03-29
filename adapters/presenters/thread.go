package presenters

import (
	"net/http"
	"vspro/entities"
)

type ThreadPresenter struct{}

func NewThreadPresenter() *ThreadPresenter {
	return &ThreadPresenter{}
}

type Threads []*Thread

type Thread struct {
	ID              string
	BulletinBoardID string
	Title           string
}

func (tp *ThreadPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (tp *ThreadPresenter) ConvertToHttpThreadListResponse(tl []*entities.Thread) *HTTPResponse {
	res := Threads{}
	for _, t := range tl {
		res = append(res, convertEntitiesThreadToThread(t))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (tp *ThreadPresenter) ConvertToHttpThreadResponse(t *entities.Thread) *HTTPResponse {
	res := Threads{}
	pt := convertEntitiesThreadToThread(t)
	res = append(res, pt)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesThreadToThread(t *entities.Thread) *Thread {
	pt := Thread{
		ID:              t.ID.String(),
		BulletinBoardID: t.BulletinBoardID.String(),
		Title:           t.Title,
	}
	return &pt
}
