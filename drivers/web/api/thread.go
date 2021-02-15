package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// postThread はPostされてきたThread(json)を保存します。
func (r *Router) postThread(c *gin.Context) {
	t, err := r.ThreadController.AddThread(c)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.ThreadPresenter.ConvertToHttpThreadResponse(t)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusCreated, res)
	return
}

// getThreadByID 指定したIDのThreadをjsonで出力します。
func (r *Router) getThreadByID(c *gin.Context) {
	tid := c.Param("id")
	t, err := r.ThreadController.GetThreadByID(tid)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.ThreadPresenter.ConvertToHttpThreadResponse(t)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// listThread はThreadの一覧をjsonで出力します。
func (r *Router) listThread(c *gin.Context) {
	tl, err := r.ThreadController.ListThread()
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.ThreadPresenter.ConvertToHttpThreadListResponse(tl)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
