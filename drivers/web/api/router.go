package api

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/adapters/gateways"
	"vspro/adapters/middlewares/logger"
	"vspro/adapters/presenters"
	"vspro/drivers/configs"
	"vspro/entities/valueobjects"

	"github.com/gin-gonic/gin"
)

// Listen はAPIがリクエストを受け取れる様に待機状態にします。
func Listen() {
	gin.DisableConsoleColor()

	c := configs.GetOsConfigInstance()
	// c := configs.GetYamlConfigInstance()
	switch c.Get().Environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger.DefaultLogger)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			questions := v1.Group("/questions")
			{
				questions.GET("", listQuestion)
				questions.GET("/:id", getQuestionByID)
				questions.POST("", postQuestion)
				questions.DELETE("/:id", deleteQuestionByID)
			}

			bulletinBoards := v1.Group("/bulletinBoards")
			{
				bulletinBoards.GET("", listBulletinBoard)
				bulletinBoards.GET("/:id", getBulletinBoardByID)
				bulletinBoards.POST("", postBulletinBoard)
			}

			threads := v1.Group("/threads")
			{
				threads.GET("", listThread)
				threads.GET("/:id", getThreadByID)
				threads.POST("", postThread)
			}
		}
	}

	router.Run(":8080")
}

// postThread はPostされてきたjsonを保存します。
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

// postBulletinBoard はPostされてきたjsonを保存します。
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

// listQuestion はQuestionの一覧をjsonで出力します。
func listQuestion(c *gin.Context) {
	qr := gateways.GetInMemoryRepositoryInstance()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	ql, err := qc.ListQuestion()
	if err != nil {
		responseByError(c, err)
		return
	}

	res := qp.ConvertToHttpQuestionListResponse(ql)
	c.JSON(http.StatusOK, res)
	return
}

// getQuestionByID 指定したIDのQuestionをjsonで出力します。
func getQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.GetInMemoryRepositoryInstance()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.GetQuestionByID(qid)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusOK, res)
	return
}

// postQuestion はPostされてきたjsonを保存します。
func postQuestion(c *gin.Context) {
	qr := gateways.GetInMemoryRepositoryInstance()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.AddQuestion(c)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusCreated, res)
	return
}

// deleteQuestionByID 指定したIDのQuestionを削除します。
func deleteQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.GetInMemoryRepositoryInstance()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	err := qc.DeleteQuestionByID(qid)
	if err != nil {
		responseByError(c, err)
		return
	}

	res := qp.ConvertToHttpDeleteQuestionResponse(http.StatusOK, qid)
	c.JSON(http.StatusOK, res)
	return
}

func responseByError(c *gin.Context, err error) {
	ep := presenters.NewErrorPresenter()
	if err != nil {
		switch t := err.(type) {
		case *valueobjects.NotFoundError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.Debug(c, t.Error())
		case *valueobjects.MissingRequiredFieldsError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.Warn(c, t.Error())
		case *valueobjects.ParameterBindingError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.Warn(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, ep.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.Error(c, t.Error())
		}
	}
}
