package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// listBulletinBoard はBulletinBoardの一覧をjsonで出力します。
func (r *Router) listBulletinBoard(c *gin.Context) {
	bbl, err := r.BulletinBoardController.ListBulletinBoard()
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.BulletinBoardPresenter.ConvertToHttpBulletinBoardListResponse(bbl)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// getBulletinBoardByID 指定したIDのBulletinBoardをjsonで出力します。
func (r *Router) getBulletinBoardByID(c *gin.Context) {
	bbid := c.Param("id")
	bb, err := r.BulletinBoardController.GetBulletinBoardByID(bbid)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.BulletinBoardPresenter.ConvertToHttpBulletinBoardResponse(bb)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// postBulletinBoard はPostされてきたBulletinBoard(json)を保存します。
func (r *Router) postBulletinBoard(c *gin.Context) {
	bb, err := r.BulletinBoardController.AddBulletinBoard(c)
	if err != nil {
		r.responseByError(c, err)
		return
	}

	res := r.BulletinBoardPresenter.ConvertToHttpBulletinBoardResponse(bb)
	if err != nil {
		r.responseByError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
	return
}
