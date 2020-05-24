package api

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/presenters"

	"github.com/gin-gonic/gin"
)

// listBulletinBoard はBulletinBoardの一覧をjsonで出力します。
func listBulletinBoard(c *gin.Context) {
	bbc := controllers.NewBulletinBoardController()
	bbl, err := bbc.ListBulletinBoard()
	if err != nil {
		responseByError(c, err)
		return
	}

	bbp := presenters.NewBulletinBoardPresenter()
	res := bbp.ConvertToHttpBulletinBoardListResponse(bbl)
	c.JSON(http.StatusOK, res)
	return
}

// getBulletinBoardByID 指定したIDのBulletinBoardをjsonで出力します。
func getBulletinBoardByID(c *gin.Context) {
	bbid := c.Param("id")
	bbc := controllers.NewBulletinBoardController()
	bb, err := bbc.GetBulletinBoardByID(bbid)
	if err != nil {
		responseByError(c, err)
		return
	}

	bbp := presenters.NewBulletinBoardPresenter()
	res := bbp.ConvertToHttpBulletinBoardResponse(bb)
	c.JSON(http.StatusOK, res)
	return
}

// postBulletinBoard はPostされてきたBulletinBoard(json)を保存します。
func postBulletinBoard(c *gin.Context) {
	bbc := controllers.NewBulletinBoardController()
	bb, err := bbc.AddBulletinBoard(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	bbp := presenters.NewBulletinBoardPresenter()
	res := bbp.ConvertToHttpBulletinBoardResponse(bb)
	c.JSON(http.StatusCreated, res)
	return
}
