package api

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/gateways"
	"vspro/adapters/presenters"

	"github.com/gin-gonic/gin"
)

// postThread はPostされてきたThread(json)を保存します。
func postThread(c *gin.Context) {
	tr := gateways.GetInMemoryRepositoryInstance()
	tp := presenters.NewThreadPresenter()
	tc := controllers.NewThreadController(tr)
	t, err := tc.AddThread(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := tp.ConvertToHttpThreadResponse(t)
	c.JSON(http.StatusCreated, res)
	return
}

// getThreadByID 指定したIDのThreadをjsonで出力します。
func getThreadByID(c *gin.Context) {
	tid := c.Param("id")
	tr := gateways.GetInMemoryRepositoryInstance()
	tp := presenters.NewThreadPresenter()
	tc := controllers.NewThreadController(tr)
	t, err := tc.GetThreadByID(tid)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := tp.ConvertToHttpThreadResponse(t)
	c.JSON(http.StatusOK, res)
	return
}

// listThread はThreadの一覧をjsonで出力します。
func listThread(c *gin.Context) {
	r := gateways.GetInMemoryRepositoryInstance()
	tp := presenters.NewThreadPresenter()
	tc := controllers.NewThreadController(r)
	tl, err := tc.ListThread()
	if err != nil {
		responseByError(c, err)
		return
	}

	res := tp.ConvertToHttpThreadListResponse(tl)
	c.JSON(http.StatusOK, res)
	return
}
