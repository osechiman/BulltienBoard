package presenters

import (
	"bulltienboard/entities"
)

// CommentPresenter はentities.Commentを外部へ渡す為にデータの変換を行います。
type CommentPresenter struct{}

// NewCommentPresenter はCommentPresenterを初期化します。
func NewCommentPresenter() *CommentPresenter {
	return &CommentPresenter{}
}

// Comment はentitiesを外部へ渡す際に利用するStructです。
type Comment struct {
	ID       string // ID はCommentのIDです。
	ThreadID string // ThreadIDer はThreadのIDです。
	Text     string // Text はCommentの内容です。
	CreatAt  int64  // CreatAt はCommentが登録された時間をunixTimeです。
}

// ConvertToHTTPCommentListResponse はComment一覧のレスポンスを返却します。
func (cp *CommentPresenter) ConvertToHTTPCommentListResponse(cl []entities.Comment) *HTTPResponse {
	res := make([]Comment, 0)
	for _, c := range cl {
		res = append(res, convertEntitiesCommentToComment(c))
	}
	return newHTTPSuccessResponse(res)
}

// ConvertToHTTPCommentResponse はCommentのレスポンスを返却します。
func (cp *CommentPresenter) ConvertToHTTPCommentResponse(c entities.Comment) *HTTPResponse {
	res := make([]Comment, 0)
	pc := convertEntitiesCommentToComment(c)
	res = append(res, pc)
	return newHTTPSuccessResponse(res)
}

// convertEntitiesCommentToComment はentities.CommentからHTTPレスポンス用のStructを返却します。
func convertEntitiesCommentToComment(c entities.Comment) Comment {
	pc := Comment{
		ID:       c.ID.String(),
		ThreadID: c.ThreadID.String(),
		Text:     c.Text,
		CreatAt:  c.CreateAt.ToUnixTime(),
	}
	return pc
}
