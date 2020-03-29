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
		}
	}

	router.Run(":8080")
}

// listQuestion はQuestionの一覧をjsonで出力します。
func listQuestion(c *gin.Context) {
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	ql, err := qc.ListQuestion()
	if err != nil {
		c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, err))
		logger.Debug(c, err.Error())
		return
	}

	res := qp.ConvertToHttpQuestionListResponse(ql)
	c.JSON(http.StatusOK, res)
	return
}

// getQuestionByID 指定したIDのQuestionをjsonで出力します。
func getQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.GetQuestionByID(qid)
	if err != nil {
		switch t := err.(type) {
		case *valueobjects.NotFoundError:
			c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, t))
			logger.Debug(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, qp.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.Error(c, t.Error())
		}
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusOK, res)
	return
}

// postQuestion はPostされてきたjsonを保存します。
func postQuestion(c *gin.Context) {
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.AddQuestion(c)
	if err != nil {
		switch t := err.(type) {
		case *valueobjects.NotFoundError:
			c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, t))
			logger.Debug(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, qp.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.Error(c, t.Error())
		}
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusCreated, res)
	return
}

// deleteQuestionByID 指定したIDのQuestionを削除します。
func deleteQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	err := qc.DeleteQuestionByID(qid)
	if err != nil {
		switch t := err.(type) {
		case *valueobjects.NotFoundError:
			c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, t))
			logger.Debug(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, qp.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.Error(c, t.Error())
		}
		return
	}

	res := qp.ConvertToHttpDeleteQuestionResponse(http.StatusOK, qid)
	c.JSON(http.StatusOK, res)
	return
}
