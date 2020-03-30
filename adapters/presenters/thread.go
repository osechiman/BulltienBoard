package presenters

import (
	"net/http"
	"vspro/entities"
)

type ThreadPresenter struct{}

func NewThreadPresenter() *ThreadPresenter {
	return &ThreadPresenter{}
}

type Thread struct {
	ID              string
	BulletinBoardID string
	Title           string
	Comments        []*Comment
}

func (tp *ThreadPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (tp *ThreadPresenter) ConvertToHttpThreadListResponse(tl []*entities.Thread) *HTTPResponse {
	res := make([]*Thread, 0)
	for _, t := range tl {
		pt := convertEntitiesThreadToThread(t)
		pt.Comments = make([]*Comment, 0)
		res = append(res, pt)
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (tp *ThreadPresenter) ConvertToHttpThreadResponse(t *entities.Thread) *HTTPResponse {
	res := make([]*Thread, 0)
	pt := convertEntitiesThreadToThread(t)
	if pt.Comments == nil {
		pt.Comments = make([]*Comment, 0)
	}
	res = append(res, pt)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesThreadToThread(t *entities.Thread) *Thread {
	pt := Thread{
		ID:              t.ID.String(),
		BulletinBoardID: t.BulletinBoardID.String(),
		Title:           t.Title,
	}

	// see https://golang.org/doc/faq#nil_error
	if t.Comments == nil {
		pt.Comments = make([]*Comment, 0)
		return &pt
	}

	cl := make([]*Comment, 0)
	for _, c := range t.Comments {
		cl = append(cl, &Comment{
			ID:       c.ID.String(),
			ThreadID: c.ThreadID.String(),
			Text:     c.Text,
			CreatAt:  c.CreateAt.ToUnixTime(),
		})
	}
	pt.Comments = cl
	return &pt
}
