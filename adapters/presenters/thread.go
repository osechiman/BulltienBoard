package presenters

import (
	"bulltienboard/entities"
)

// ThreadPresenter はentities.Threadを外部へ渡す為にデータの変換を行います。
type ThreadPresenter struct{}

// NewThreadPresenter はThreadPresenterを初期化します。
func NewThreadPresenter() *ThreadPresenter {
	return &ThreadPresenter{}
}

// Thread はentitiesを外部へ渡す際に利用するStructです。
type Thread struct {
	ID              string    // ID はThreadのIDです。
	BulletinBoardID string    // BulletinBoardIDer はBulletinBoardのIDです。
	Title           string    // Title はThreadのTitleです。
	Comments        []Comment // Comments はCommentの一覧です。
}

// ConvertToHTTPThreadListResponse はThread一覧のレスポンスを返却します。
func (tp *ThreadPresenter) ConvertToHTTPThreadListResponse(tl []entities.Thread) *HTTPResponse {
	res := make([]Thread, 0)
	for _, t := range tl {
		pt := convertEntitiesThreadToThread(t)
		pt.Comments = make([]Comment, 0)
		res = append(res, pt)
	}
	return newHTTPSuccessResponse(res)
}

// ConvertToHTTPThreadResponse はCommentを含むThreadのレスポンスを返却します。
func (tp *ThreadPresenter) ConvertToHTTPThreadResponse(t entities.Thread) *HTTPResponse {
	res := make([]Thread, 0)
	pt := convertEntitiesThreadToThread(t)
	if pt.Comments == nil {
		pt.Comments = make([]Comment, 0)
	}
	res = append(res, pt)
	return newHTTPSuccessResponse(res)
}

// convertEntitiesThreadToThread はentities.ThreadからHTTPレスポンス用のStructを返却します。
func convertEntitiesThreadToThread(t entities.Thread) Thread {
	pt := Thread{
		ID:              t.ID.String(),
		BulletinBoardID: t.BulletinBoardID.String(),
		Title:           t.Title,
	}

	// see https://golang.org/doc/faq#nil_error
	if t.Comments == nil {
		pt.Comments = make([]Comment, 0)
		return pt
	}

	cl := make([]Comment, 0)
	for _, c := range t.Comments {
		cl = append(cl, Comment{
			ID:       c.ID.String(),
			ThreadID: c.ThreadID.String(),
			Text:     c.Text,
			CreatAt:  c.CreateAt.ToUnixTime(),
		})
	}
	pt.Comments = cl
	return pt
}
