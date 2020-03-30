package presenters

import (
	"net/http"
	"vspro/entities"
)

type BulletinBoardPresenter struct{}

func NewBulletinBoardPresenter() *BulletinBoardPresenter {
	return &BulletinBoardPresenter{}
}

type BulletinBoards []*BulletinBoard

type BulletinBoard struct {
	ID      string
	Title   string
	Threads []*Thread
}

func (bbp *BulletinBoardPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}

func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardListResponse(bbl []*entities.BulletinBoard) *HTTPResponse {
	res := make([]*BulletinBoard, 0)
	for _, bb := range bbl {
		// レスポンス時のデータ転送量を減らす為にListResponseではThreadのデータは返さない
		bb.Threads = nil
		res = append(res, convertEntitiesBulletinBoardToBulletinBoard(bb))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardResponse(bb *entities.BulletinBoard) *HTTPResponse {
	res := make([]*BulletinBoard, 0)
	pbb := convertEntitiesBulletinBoardToBulletinBoard(bb)
	res = append(res, pbb)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

func convertEntitiesBulletinBoardToBulletinBoard(bb *entities.BulletinBoard) *BulletinBoard {
	pbb := BulletinBoard{
		ID:    bb.ID.String(),
		Title: bb.Title,
	}

	// see https://golang.org/doc/faq#nil_error
	if bb.Threads == nil {
		pbb.Threads = make([]*Thread, 0)
		return &pbb
	}

	tl := make([]*Thread, 0)
	cl := make([]*Comment, 0)
	for _, t := range bb.Threads {
		tl = append(tl, &Thread{
			ID:              t.ID.String(),
			BulletinBoardID: t.BulletinBoardID.String(),
			Title:           t.Title,
			Comments:        cl,
		})
	}
	pbb.Threads = tl
	return &pbb
}
