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

// ThreadController はThreadRepositorerのコントローラーです。
// 初期化時に渡すリポジトリ以外を利用したい場合はそれぞれメソッドの引数で受け取ってください。
type ThreadController struct {
	// Repository はこのコントローラーで利用するメインのリポジトリです。
	// このコントローラーで利用するメインのリポジトリです。
	Repository usecases.ThreadRepositorer
}

// Thread はリクエストされてきたPost値を受け取る為のStructです。
type Thread struct {
	ID              string // ID はThreadのIDです。
	BulletinBoardID string `validate:"required"`              // BulletinBoardIDer はBulletinBoardのIDです。
	Title           string `validate:"required,min=1,max=50"` // Title はユーザーが入力した文字列です。
}

// NewThreadController はThreadControllerを初期化します。
func NewThreadController(r usecases.ThreadRepositorer) *ThreadController {
	return &ThreadController{Repository: r}
}

// GetThreadByID はThreadIDからThreadを取得します。
func (tc *ThreadController) GetThreadByID(ID string) (*entities.Thread, error) {
	tu := usecases.NewThreadUsecase(tc.Repository)

	tid, err := convertIDToThreadID(ID)
	if err != nil {
		return nil, err
	}

	return tu.GetThreadByID(tid, gateways.GetInMemoryRepositoryInstance())
}

// AddThread はPostされてきたデータを元にThreadを追加します。
// コマンド・クエリの原則からは外れますがAPIのレスポンスに登録したデータを返却するためにエンティティを返します。
func (tc *ThreadController) AddThread(c *gin.Context) (*entities.Thread, error) {
	pt := Thread{}
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

	t, err := entities.NewThread(tid.Get(), bid, pt.Title)
	if err != nil {
		return nil, err
	}

	tu := usecases.NewThreadUsecase(tc.Repository)
	return &t, tu.AddThread(t, gateways.GetInMemoryRepositoryInstance())
}

// ListThread はThreadの一覧を取得します。
func (tc *ThreadController) ListThread() ([]*entities.Thread, error) {
	tu := usecases.NewThreadUsecase(tc.Repository)
	return tu.ListThread()
}

// ListThreadByBulletinBoardID は指定されたBulletinBoardIDを持つThread一覧を取得します。
func (tc *ThreadController) ListThreadByBulletinBoardID(bID string) ([]*entities.Thread, error) {
	bid, err := convertIDToBulletinBoardID(bID)
	if err != nil {
		return nil, err
	}
	tu := usecases.NewThreadUsecase(tc.Repository)
	return tu.ListThreadByBulletinBoardID(bid)
}

// convertIDToThreadID は文字列のThreadIDをvalueobjects.ThreadIDに変換します。
func convertIDToThreadID(ID string) (valueobjects.ThreadID, error) {
	return valueobjects.NewThreadID(ID)
}
