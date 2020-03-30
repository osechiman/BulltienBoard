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

type ThreadController struct {
	Repository usecases.ThreadRepositorer
}

type PostThread struct {
	ID              string
	BulletinBoardID string `validate:"required"`
	Title           string `validate:"required,min=1,max=50"`
}

func NewThreadController(r usecases.ThreadRepositorer) *ThreadController {
	return &ThreadController{Repository: r}
}

func (tc *ThreadController) GetThreadByID(ID string) (*entities.Thread, error) {
	tu := usecases.NewThreadUsecase(tc.Repository)

	tid, err := convertIDToThreadID(ID)
	if err != nil {
		return nil, err
	}

	return tu.GetThreadByID(tid, gateways.GetInMemoryRepositoryInstance())
}

// コマンド・クエリの原則からは外れるがAPIのレスポンスに登録したデータを返却するためにエンティティを返す
func (tc *ThreadController) AddThread(c *gin.Context) (*entities.Thread, error) {
	pt := PostThread{}
	err := c.BindJSON(&pt)
	if err != nil {
		return nil, errorobjects.NewParameterBindingError(err)
	}

	validate := validator.New()
	err = validate.Struct(pt)
	if err != nil {
		return nil, errorobjects.NewMissingRequiredFieldsError(err)
	}

	tid, err := valueobjects.NewThreadID("")
	if err != nil {
		return nil, err
	}

	convertIDToBulletinBoardID(pt.BulletinBoardID)
	bid, err := convertIDToBulletinBoardID(pt.BulletinBoardID)
	if err != nil {
		return nil, err
	}
	t := entities.NewThread(tid.Get(), bid, pt.Title)

	tu := usecases.NewThreadUsecase(tc.Repository)
	return &t, tu.AddThread(t, gateways.GetInMemoryRepositoryInstance())
}

func (tc *ThreadController) ListThread() ([]*entities.Thread, error) {
	tu := usecases.NewThreadUsecase(tc.Repository)
	return tu.ListThread()
}

func (tc *ThreadController) ListThreadByByBulletinBoard(bID string) ([]*entities.Thread, error) {
	bid, err := convertIDToBulletinBoardID(bID)
	if err != nil {
		return nil, err
	}
	tu := usecases.NewThreadUsecase(tc.Repository)
	return tu.ListThreadByBulletinBoardID(bid)
}

func convertIDToThreadID(ID string) (valueobjects.ThreadID, error) {
	return valueobjects.NewThreadID(ID)
}
