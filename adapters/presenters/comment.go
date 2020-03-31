package presenters

import (
	"net/http"
	"vspro/entities"
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
	ThreadID string // ThreadID はThreadのIDです。
	Text     string // Text はCommentの内容です。
	CreatAt  int64  // CreatAt はCommentが登録された時間をunixTimeです。
}

// ConvertToHttpCommentListResponse はComment一覧のレスポンスを返却します。
func (cp *CommentPresenter) ConvertToHttpCommentListResponse(cl []*entities.Comment) *HTTPResponse {
	res := make([]*Comment, 0)
	for _, c := range cl {
		res = append(res, convertEntitiesCommentToComment(c))
	}
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

// ConvertToHttpCommentResponse はCommentのレスポンスを返却します。
func (cp *CommentPresenter) ConvertToHttpCommentResponse(c *entities.Comment) *HTTPResponse {
	res := make([]*Comment, 0)
	pc := convertEntitiesCommentToComment(c)
	res = append(res, pc)
	return newHTTPSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), res)
}

// convertEntitiesCommentToComment はentities.CommentからHTTPレスポンス用のStructを返却します。
func convertEntitiesCommentToComment(c *entities.Comment) *Comment {
	pc := Comment{
		ID:       c.ID.String(),
		ThreadID: c.ThreadID.String(),
		Text:     c.Text,
		CreatAt:  c.CreateAt.ToUnixTime(),
	}
	return &pc
}
