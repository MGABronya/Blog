// @Title  Guestbook
// @Description  接收前端请求时的留言信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package vo

// CreateGuestbookRequest			接收前端请求时的留言信息
type CreateGuestbookRequest struct {
	Content  string `json:"content" binding:"required"` // 内容
}
