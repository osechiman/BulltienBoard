package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postComment はPostされてきたComment(json)を保存します。
func (r *Router) postComment(c *gin.Context) {
	cm, err := r.CommentController.AddComment(c)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.CommentPresenter.ConvertToHTTPCommentResponse(cm)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
	return
}

// listComment はCommentの一覧をjsonで出力します。
func (r *Router) listComment(c *gin.Context) {
	cl, err := r.CommentController.ListComment()
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.CommentPresenter.ConvertToHTTPCommentListResponse(cl)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
