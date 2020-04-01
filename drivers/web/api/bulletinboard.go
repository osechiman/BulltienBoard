package api

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/gateways"
	"vspro/adapters/presenters"

	"github.com/gin-gonic/gin"
)

// listBulletinBoard はBulletinBoardの一覧をjsonで出力します。
func listBulletinBoard(c *gin.Context) {
	r := gateways.GetInMemoryRepositoryInstance()
	bbp := presenters.NewBulletinBoardPresenter()
	bbc := controllers.NewBulletinBoardController(r)
	bbl, err := bbc.ListBulletinBoard()
	if err != nil {
		responseByError(c, err)
		return
	}

	res := bbp.ConvertToHttpBulletinBoardListResponse(bbl)
	c.JSON(http.StatusOK, res)
	return
}

// getBulletinBoardByID 指定したIDのBulletinBoardをjsonで出力します。
func getBulletinBoardByID(c *gin.Context) {
	bbid := c.Param("id")
	bbr := gateways.GetInMemoryRepositoryInstance()
	bbp := presenters.NewBulletinBoardPresenter()
	bbc := controllers.NewBulletinBoardController(bbr)
	bb, err := bbc.GetBulletinBoardByID(bbid)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := bbp.ConvertToHttpBulletinBoardResponse(bb)
	c.JSON(http.StatusOK, res)
	return
}

// postBulletinBoard はPostされてきたBulletinBoard(json)を保存します。
func postBulletinBoard(c *gin.Context) {
	bbr := gateways.GetInMemoryRepositoryInstance()
	bbp := presenters.NewBulletinBoardPresenter()
	bbc := controllers.NewBulletinBoardController(bbr)
	bb, err := bbc.AddBulletinBoard(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := bbp.ConvertToHttpBulletinBoardResponse(bb)
	c.JSON(http.StatusCreated, res)
	return
}
