package api

import (
	"bulltienboard/adapters/controllers"
	"bulltienboard/adapters/middlewares/logger"
	"bulltienboard/adapters/presenters"
	"bulltienboard/drivers/configs"
	"bulltienboard/entities/errorobjects"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router はWebAPIのリクエストを処理するために必要な構造体をまとめた構造体です。
type Router struct {
	BulletinBoardController *controllers.BulletinBoardController
	BulletinBoardPresenter  *presenters.BulletinBoardPresenter
	ThreadController        *controllers.ThreadController
	ThreadPresenter         *presenters.ThreadPresenter
	CommentController       *controllers.CommentController
	CommentPresenter        *presenters.CommentPresenter
	ErrorPresenter          *presenters.ErrorPresenter
}

// NewRouter はRouterのコンストラクタです。
func NewRouter(bulletinBoardController *controllers.BulletinBoardController,
	bulletinBoardPresenter *presenters.BulletinBoardPresenter,
	threadController *controllers.ThreadController,
	threadPresenter *presenters.ThreadPresenter,
	commentController *controllers.CommentController,
	commentPresenter *presenters.CommentPresenter,
	errorPresenter *presenters.ErrorPresenter) *Router {
	return &Router{BulletinBoardController: bulletinBoardController,
		BulletinBoardPresenter: bulletinBoardPresenter,
		ThreadController:       threadController,
		ThreadPresenter:        threadPresenter,
		CommentController:      commentController,
		CommentPresenter:       commentPresenter,
		ErrorPresenter:         errorPresenter,
	}
}

// Listen はAPIがリクエストを受け取れる様に待機状態にします。
func Listen(r *Router) {
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
				bulletinBoards.GET("", r.listBulletinBoard)
				bulletinBoards.GET("/:id", r.getBulletinBoardByID)
				bulletinBoards.POST("", r.postBulletinBoard)
			}

			threads := v1.Group("/threads")
			{
				threads.GET("", r.listThread)
				threads.GET("/:id", r.getThreadByID)
				threads.POST("", r.postThread)
			}

			comments := v1.Group("/comments")
			{
				comments.GET("", r.listComment)
				comments.POST("", r.postComment)
			}
		}
	}

	router.Run(":8080")
}

// responseByError はerrorobjectsのType毎にjsonを出力します。
func (r *Router) responseByError(c *gin.Context, err error) {
	if err != nil {
		res := r.ErrorPresenter.ConvertToHTTPErrorResponse(err)
		switch t := err.(type) {
		case *errorobjects.NotFoundError:
			c.JSON(http.StatusNotFound, res)
			logger.GetLoggerColumns(c).Debug(c, t.Error())
		case *errorobjects.MissingRequiredFieldsError:
			c.JSON(http.StatusBadRequest, res)
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.ParameterBindingError:
			c.JSON(http.StatusBadRequest, res)
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.CharacterSizeValidationError:
			c.JSON(http.StatusBadRequest, res)
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		case *errorobjects.ResourceLimitedError:
			c.JSON(http.StatusInsufficientStorage, res)
			logger.GetLoggerColumns(c).Warn(c, t.Error())
		default:
			c.JSON(http.StatusInternalServerError, res)
			logger.GetLoggerColumns(c).Error(c, t.Error())
		}
	}
}
