package api

import (
	"net/http"
	"vspro/adapters/middlewares/logger"
	"vspro/adapters/presenters"
	"vspro/drivers/configs"
	"vspro/entities/errorobjects"

	"github.com/gin-gonic/gin"
)

// Listen はAPIがリクエストを受け取れる様に待機状態にします。
func Listen() {
	gin.DisableConsoleColor()

	c := configs.GetOsConfigInstance()
	// configerインターフェースを満たして実装すれば以下の様に置き換え可能になります。
	// c := configs.GetYamlConfigInstance()
	switch c.Get().Environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logger.DefaultLogger)

	// パス毎にGroupを分けるポリシーです。
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
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

			comments := v1.Group("/comments")
			{
				comments.GET("", listComment)
				comments.POST("", postComment)
			}
		}
	}

	router.Run(":8080")
}

// responseByError はerrorobjectsのType毎にjsonを出力します。
func responseByError(c *gin.Context, err error) {
	ep := presenters.NewErrorPresenter()
	if err != nil {
		switch t := err.(type) {
		case *errorobjects.NotFoundError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.GetLoggerColumns(c).Debug(c, t.Error())
		case *errorobjects.MissingRequiredFieldsError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.ParameterBindingError:
			c.JSON(t.HTTPStatusCode, ep.ConvertToHttpErrorResponse(t.HTTPStatusCode, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, ep.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.GetLoggerColumns(c).Error(c, t.Error())
		}
	}
}
