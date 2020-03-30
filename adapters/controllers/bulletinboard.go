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

// BulletinBoardController はBulletinBoardRepositorerのコントローラーです。
// 初期化時に渡すリポジトリ以外を利用したい場合はそれぞれメソッドの引数で受け取ってください。
type BulletinBoardController struct {
	Repository usecases.BulletinBoardRepositorer
}

type PostBulletinBoard struct {
	ID    string
	Title string `validate:"required,min=1,max=50"`
}

func NewBulletinBoardController(r usecases.BulletinBoardRepositorer) *BulletinBoardController {
	return &BulletinBoardController{Repository: r}
}

func (bbc *BulletinBoardController) GetBulletinBoardByID(ID string) (*entities.BulletinBoard, error) {
	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)

	bbid, err := convertIDToBulletinBoardID(ID)
	if err != nil {
		return nil, err
	}

	tr := gateways.GetInMemoryRepositoryInstance()
	return bbu.GetBulletinBoardByID(bbid, tr)
}

// コマンド・クエリの原則からは外れるがAPIのレスポンスに登録したデータを返却するためにエンティティを返す
func (bbc *BulletinBoardController) AddBulletinBoard(c *gin.Context) (*entities.BulletinBoard, error) {
	pb := PostBulletinBoard{}
	err := c.BindJSON(&pb)
	if err != nil {
		return nil, errorobjects.NewParameterBindingError(err)
	}
	validate := validator.New()
	err = validate.Struct(pb)
	if err != nil {
		return nil, errorobjects.NewMissingRequiredFieldsError(err)
	}

	bbid, err := valueobjects.NewBulletinBoardID("")
	if err != nil {
		return nil, err
	}
	bb := entities.NewBulletinBoard(bbid, pb.Title)

	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)
	return &bb, bbu.AddBulletinBoard(bb)
}

func (bbc *BulletinBoardController) ListBulletinBoard() ([]*entities.BulletinBoard, error) {
	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)
	return bbu.ListBulletinBoard()
}

func convertIDToBulletinBoardID(ID string) (entities.BulletinBoardID, error) {
	return valueobjects.NewBulletinBoardID(ID)
}
