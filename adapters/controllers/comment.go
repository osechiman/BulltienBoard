package controllers

import (
	"vspro/adapters/gateways"
	"vspro/entities"
	"vspro/entities/errorobjects"
	"vspro/entities/valueobjects"
	"vspro/usecases"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type CommentController struct {
	Repository usecases.CommentRepositorer
}

type Comment struct {
	ID       string
	ThreadID string `validate:"required"`
	Text     string `validate:"required,min=1,max=2048"`
}

func NewCommentController(r usecases.CommentRepositorer) *CommentController {
	return &CommentController{Repository: r}
}

// コマンド・クエリの原則からは外れるがAPIのレスポンスに登録したデータを返却するためにエンティティを返す
func (cc *CommentController) AddComment(c *gin.Context) (*entities.Comment, error) {
	pc := Comment{}
	err := c.BindJSON(&pc)
	if err != nil {
		return nil, errorobjects.NewParameterBindingError(err)
	}

	validate := validator.New()
	err = validate.Struct(pc)
	if err != nil {
		return nil, errorobjects.NewMissingRequiredFieldsError(err)
	}

	cid, err := valueobjects.NewCommentID("")
	if err != nil {
		return nil, err
	}

	tid, err := convertIDToThreadID(pc.ThreadID)
	if err != nil {
		return nil, err
	}

	ct, err := valueobjects.NewCommentTime(0)
	if err != nil {
		return nil, err
	}

	cm := entities.NewComment(cid, tid, pc.Text, ct)

	cu := usecases.NewCommentUsecase(cc.Repository)
	return &cm, cu.AddComment(cm, gateways.GetInMemoryRepositoryInstance())
}

func (cc *CommentController) ListComment() ([]*entities.Comment, error) {
	tu := usecases.NewCommentUsecase(cc.Repository)
	return tu.ListComment()
}

func (cc *CommentController) ListCommentByThreadID(tID string) ([]*entities.Comment, error) {
	tid, err := convertIDToThreadID(tID)
	if err != nil {
		return nil, err
	}
	tu := usecases.NewCommentUsecase(cc.Repository)
	return tu.ListCommentByThreadID(tid)
}

func convertIDToCommentID(ID string) (entities.CommentID, error) {
	return valueobjects.NewCommentID(ID)
}

func convertCreatAtToCommentTime(unixTime int64) (entities.CommentTime, error) {
	return valueobjects.NewCommentTime(unixTime)
}
