// @Title  guestbook
// @Description  定义用户留言板
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// GuestBook			定义留言板
type GuestBook struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`      // 留言id
	UserId    uint       `json:"user_id" gorm:"index:idx_member;not null"` // 所属者的id
	Author    uint       `json:"author" gorm:"not null"`                   // 作者的id
	Content   string     `json:"content" gorm:"type:text;not null"`        // 留言的内容
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`         // 创建时间
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (guestBook *GuestBook) BeforeCreate(scope *gorm.DB) error {
	guestBook.ID = uuid.NewV4()
	return nil
}
