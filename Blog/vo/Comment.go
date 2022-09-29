// @Title  Comment
// @Description  接收前端请求时的评论信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package vo

// CreateCommentRequest			接收前端请求时的评论信息
type CreateCommentRequest struct {
	Content  string `json:"content" binding:"required"` // 内容
	ResLong  string `json:"res_long"`                   // 备用长文本
	ResShort string `json:"res_short"`                  // 备用短文本
}
