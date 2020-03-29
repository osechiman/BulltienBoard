package presenters

import (
	"net/http"
	"vspro/adapters/controllers"
	"vspro/entities"
)

type BulletinBoardPresenter struct{}

func NewBulletinBoardPresenter() *BulletinBoardPresenter {
	return &BulletinBoardPresenter{}
}

type BulletinBoards []*BulletinBoard

type BulletinBoard struct {
	ID     string
	Title  string
	Thread interface{} `json:",omitempty"`
}

func (bbp *BulletinBoardPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardListResponse(bbl []*entities.BulletinBoard) *HTTPResponse {
	res := BulletinBoards{}
	for _, bb := range bbl {
		// レスポンス時のデータ転送量を減らす為にListResponseではThreadのデータは返さない
		bb.Threads = nil
		res = append(res, convertEntitiesBulletinBoardToBulletinBoard(bb))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardResponse(bb *entities.BulletinBoard) *HTTPResponse {
	res := BulletinBoards{}
	pbb := convertEntitiesBulletinBoardToBulletinBoard(bb)
	res = append(res, pbb)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesBulletinBoardToBulletinBoard(bb *entities.BulletinBoard) *BulletinBoard {
	pbb := BulletinBoard{
		ID:     bb.ID.String(),
		Title:  bb.Title,
		Thread: bb.Threads,
	}

	// see https://golang.org/doc/faq#nil_error
	if bb.Threads == nil {
		pbb.Thread = nil
		return &pbb
	}

	tl := make([]controllers.PostThread, 0)
	for _, t := range bb.Threads {
		tl = append(tl, controllers.PostThread{
			ID:              t.ID.String(),
			BulletinBoardID: t.BulletinBoardID.String(),
			Title:           t.Title,
		})
	}
	pbb.Thread = tl
	return &pbb
}
