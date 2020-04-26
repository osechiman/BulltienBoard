package api

import (
	"encoding/json"
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/gateways"
	"vspro/adapters/middlewares/logger"
	"vspro/adapters/presenters"

	"github.com/gin-gonic/gin"
)

// postComment はPostされてきたComment(json)を保存します。
func postComment(c *gin.Context) {
	cr := gateways.GetInMemoryRepositoryInstance()
	cp := presenters.NewCommentPresenter()
	cc := controllers.NewCommentController(cr)
	cm, err := cc.AddComment(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := cp.ConvertToHttpCommentResponse(cm)
	j, err := json.Marshal(res)
	if err != nil {
		logger.GetLoggerColumns(c).Warn(c, err.Error())
	}

	err = m.Broadcast(j)
	if err != nil {
		logger.GetLoggerColumns(c).Warn(c, err.Error())
	}
	c.JSON(http.StatusCreated, res)
	return
}

// listComment はCommentの一覧をjsonで出力します。
func listComment(c *gin.Context) {
	r := gateways.GetInMemoryRepositoryInstance()
	cp := presenters.NewCommentPresenter()
	cc := controllers.NewCommentController(r)
	cl, err := cc.ListComment()
	if err != nil {
		responseByError(c, err)
		return
	}

	res := cp.ConvertToHttpCommentListResponse(cl)
	c.JSON(http.StatusOK, res)
	return
}
