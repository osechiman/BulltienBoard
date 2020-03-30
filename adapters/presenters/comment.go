package presenters

import (
	"net/http"
	"vspro/entities"
)

type CommentPresenter struct{}

func NewCommentPresenter() *CommentPresenter {
	return &CommentPresenter{}
}

type Comment struct {
	ID       string
	ThreadID string
	Text     string
	CreatAt  int64
}

func (cp *CommentPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (cp *CommentPresenter) ConvertToHttpCommentListResponse(cl []*entities.Comment) *HTTPResponse {
	res := make([]*Comment, 0)
	for _, c := range cl {
		res = append(res, convertEntitiesCommentToComment(c))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (cp *CommentPresenter) ConvertToHttpCommentResponse(c *entities.Comment) *HTTPResponse {
	res := make([]*Comment, 0)
	pc := convertEntitiesCommentToComment(c)
	res = append(res, pc)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesCommentToComment(c *entities.Comment) *Comment {
	pc := Comment{
		ID:       c.ID.String(),
		ThreadID: c.ThreadID.String(),
		Text:     c.Text,
		CreatAt:  c.CreateAt.ToUnixTime(),
	}
	return &pc
}
