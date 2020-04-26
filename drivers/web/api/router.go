package api

import (
	"net/http"
	"vspro/adapters/middlewares/logger"
	"vspro/adapters/presenters"
	"vspro/drivers/configs"
	"vspro/entities/errorobjects"

	"gopkg.in/olahol/melody.v1"

	"github.com/gin-gonic/gin"
)

// TODO:: 別のパッケージでシングルトンかつセッション管理できるものを作る
var m = melody.New()

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
				comments.GET("/new", newComment)
				comments.POST("", postComment)
			}
		}
	}

	router.Run(":8080")
}

// TODO:: implements
func newComment(c *gin.Context) {
	err := m.HandleRequest(c.Writer, c.Request)
	// TODO:: cookie を取得して、存在しない場合はセットする @ gin middleware
	// https://github.com/gin-gonic/gin#set-and-get-a-cookie
	if err != nil {
		logger.GetLoggerColumns(c).Error(c, err.Error())
	}

	// m.HandleConnect(func(s *melody.Session) {
	// 	// TODO:: すでに接続済みのユーザーかどうか判定
	// })

	// m.HandleMessage(func(s *melody.Session, msg []byte) {
	// 	// TODO:: コメントの更新があれば送信する
	// })

	// m.HandleError(func(s *melody.Session, err error) {
	// 	// TODO:: 調べてエラーハンドリング

	// })

	// 	mrouter.HandleMessage(func(s *melody.Session, msg []byte) {
	// 		mrouter.Broadcast(msg)
	// 		d, _ := strconv.Atoi(string(msg))
	// 		battle.Participant[s] = User{battle.Participant[s].ID, battle.Participant[s].LifePoint - d}
	// 		l := strconv.Itoa(battle.Participant[s].LifePoint)
	// 		dm := []byte("user:" + battle.Participant[s].ID + " collect answer. remaining is " + l)
	// 		mrouter.Broadcast(dm)
	// 	})
	// 	mrouter.HandleConnect(func(s *melody.Session) {
	// 		u := User{uuid.New().String(), life}
	// 		battle.Participant[s] = u
	// 	})
	// 	router.Run(":8080")
}

// responseByError はerrorobjectsのType毎にjsonを出力します。
func responseByError(c *gin.Context, err error) {
	ep := presenters.NewErrorPresenter()
	if err != nil {
		switch t := err.(type) {
		case *errorobjects.NotFoundError:
			c.JSON(http.StatusNotFound, ep.ConvertToHttpErrorResponse(http.StatusNotFound, t))
			logger.GetLoggerColumns(c).Debug(c, t.Error())
		case *errorobjects.MissingRequiredFieldsError:
			c.JSON(http.StatusBadRequest, ep.ConvertToHttpErrorResponse(http.StatusBadRequest, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.ParameterBindingError:
			c.JSON(http.StatusBadRequest, ep.ConvertToHttpErrorResponse(http.StatusBadRequest, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.CharacterSizeValidationError:
			c.JSON(http.StatusBadRequest, ep.ConvertToHttpErrorResponse(http.StatusBadRequest, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.ResourceLimitedError:
			c.JSON(http.StatusInsufficientStorage, ep.ConvertToHttpErrorResponse(http.StatusInsufficientStorage, t))
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, ep.ConvertToHttpErrorResponse(http.StatusInternalServerError, t))
			logger.GetLoggerColumns(c).Error(c, t.Error())
		}
	}
}
