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
	// Repository はBulletinBoardRepositorerを満たしたリポジトリです。
	// このコントローラーで利用するメインのリポジトリです。
	Repository usecases.BulletinBoardRepositorer
}

// BulletinBoard はリクエストされてきたPost値を受け取る為のStructです。
type BulletinBoard struct {
	ID    string // ID はBulletinBoardのIDです。
	Title string `validate:"required"` // Title はBulletinBoardのTitleです。
}

// NewBulletinBoardController はBulletinBoardControllerを初期化します。
func NewBulletinBoardController(r usecases.BulletinBoardRepositorer) *BulletinBoardController {
	return &BulletinBoardController{Repository: r}
}

// GetBulletinBoardByID はBulletinBoardIDからBulletinBoardを取得します。
func (bbc *BulletinBoardController) GetBulletinBoardByID(ID string) (*entities.BulletinBoard, error) {
	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)

	bbid, err := convertIDToBulletinBoardID(ID)
	if err != nil {
		return nil, err
	}

	tr := gateways.GetInMemoryRepositoryInstance()
	return bbu.GetBulletinBoardByID(bbid, tr)
}

// AddBulletinBoard はPostされてきたデータを元にBulletinBoardを追加します。
// コマンド・クエリの原則からは外れますがAPIのレスポンスに登録したデータを返却するためにエンティティを返します。
func (bbc *BulletinBoardController) AddBulletinBoard(c *gin.Context) (*entities.BulletinBoard, error) {
	pb := BulletinBoard{}
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

	bb, err := entities.NewBulletinBoard(bbid, pb.Title)
	if err != nil {
		return nil, err
	}

	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)
	return &bb, bbu.AddBulletinBoard(bb)
}

// ListBulletinBoard はBulletinBoardの一覧を取得します。
func (bbc *BulletinBoardController) ListBulletinBoard() ([]*entities.BulletinBoard, error) {
	bbu := usecases.NewBulletinBoardUsecase(bbc.Repository)
	return bbu.ListBulletinBoard()
}

// convertIDToBulletinBoardID は文字列のBulletinBoardIDをentities.BulletinBoardIDに変換します。
func convertIDToBulletinBoardID(ID string) (valueobjects.BulletinBoardID, error) {
	return valueobjects.NewBulletinBoardID(ID)
}
