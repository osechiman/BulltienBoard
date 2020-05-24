package api

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/presenters"

	"github.com/gin-gonic/gin"
)

// postComment はPostされてきたComment(json)を保存します。
func postComment(c *gin.Context) {
	cc := controllers.NewCommentController()
	cm, err := cc.AddComment(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	cp := presenters.NewCommentPresenter()
	res := cp.ConvertToHttpCommentResponse(cm)
	c.JSON(http.StatusCreated, res)
	return
}

// listComment はCommentの一覧をjsonで出力します。
func listComment(c *gin.Context) {
	cc := controllers.NewCommentController()
	cl, err := cc.ListComment()
	if err != nil {
		responseByError(c, err)
		return
	}

	cp := presenters.NewCommentPresenter()
	res := cp.ConvertToHttpCommentListResponse(cl)
	c.JSON(http.StatusOK, res)
	return
}
