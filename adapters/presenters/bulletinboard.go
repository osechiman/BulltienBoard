package presenters

import (
	"vspro/entities"
)

// BulletinBoardPresenter はentities.BulletinBoardを外部へ渡す為にデータの変換を行います。
type BulletinBoardPresenter struct{}

// NewBulletinBoardPresenter はBulletinBoardPresenterを初期化します。
func NewBulletinBoardPresenter() *BulletinBoardPresenter {
	return &BulletinBoardPresenter{}
}

// BulletinBoard はentitiesを外部へ渡す際に利用するStructです。
type BulletinBoard struct {
	ID      string   // ID はBulletinBoardのIDです。
	Title   string   // Title はBulletinBoardのTitleです。
	Threads []Thread // Threads はThreadの一覧です。
}

// ConvertToHttpBulletinBoardListResponse はBulletinBoard一覧のレスポンスを返却します。
func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardListResponse(bbl []entities.BulletinBoard) *HTTPResponse {
	res := make([]BulletinBoard, 0)
	for _, bb := range bbl {
		// レスポンス時のデータ転送量を減らす為にListResponseではThreadのデータは返さない
		bb.Threads = nil
		res = append(res, convertEntitiesBulletinBoardToBulletinBoard(bb))
	}
	return newHTTPSuccessResponse(res)
}

// ConvertToHttpBulletinBoardResponse はThreadを含むBulletinBoardのレスポンスを返却します。
func (bbp *BulletinBoardPresenter) ConvertToHttpBulletinBoardResponse(bb entities.BulletinBoard) *HTTPResponse {
	res := make([]BulletinBoard, 0)
	pbb := convertEntitiesBulletinBoardToBulletinBoard(bb)
	res = append(res, pbb)
	return newHTTPSuccessResponse(res)
}

// convertEntitiesBulletinBoardToBulletinBoard はentities.BulletinBoardからHTTPレスポンス用のStructを返却します。
func convertEntitiesBulletinBoardToBulletinBoard(bb entities.BulletinBoard) BulletinBoard {
	pbb := BulletinBoard{
		ID:    bb.ID.String(),
		Title: bb.Title,
	}

	// see https://golang.org/doc/faq#nil_error
	if bb.Threads == nil {
		pbb.Threads = make([]Thread, 0)
		return pbb
	}

	tl := make([]Thread, 0)
	cl := make([]Comment, 0)
	for _, t := range bb.Threads {
		tl = append(tl, Thread{
			ID:              t.ID.String(),
			BulletinBoardID: t.BulletinBoardID.String(),
			Title:           t.Title,
			Comments:        cl,
		})
	}
	pbb.Threads = tl
	return pbb
}
