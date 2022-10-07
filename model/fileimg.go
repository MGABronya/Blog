// @Title  fileimg
// @Description  前端文件的展示图片
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// FileImg			定义前端文件展示图片
type FileImg struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`      // 展示图片的id
	UserId    uint       `json:"user_id" gorm:"index:idx_userId;not null"` // 作者的id
	FileId    string     `json:"file_id" gorm:"index:idx_fileId;not null"` // 所属前端文件的id
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`         // 创建时间
	UpdatedAt model.Time `json:"updated_at" gorm:"type:timestamp"`         // 更新时间
}

// @title    BeforeCreate
// @description   计算出一个uint
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (fileImg *FileImg) BeforeCreate(scope *gorm.DB) error {
	fileImg.ID = uuid.NewV4()
	return nil
}
