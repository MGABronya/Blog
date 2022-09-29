// @Title  comment
// @Description  定义评论
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Comment			定义评论
type Comment struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"` // 跟帖的id
	UserId    uint       `json:"user_id" gorm:"not null"`             // 作者的id
	FileId    string     `json:"file_id" gorm:"type:char(36)"`        // 前端文件的id
	Content   string     `json:"content" gorm:"type:text;not null"`   // 内容
	ResLong   string     `json:"res_long" gorm:"type:text;not null"`  // 备用长文本
	ResShort  string     `json:"res_short" gorm:"type:text;not null"` // 备用短文本
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`    // 创建时间
	UpdatedAt model.Time `json:"updated_at" gorm:"type:timestamp"`    // 更新时间
}

// @title    BeforeCreate
// @description   计算出一个uint
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (comment *Comment) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
