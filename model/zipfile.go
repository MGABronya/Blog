// @Title  zipfile
// @Description  定义压缩文件
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ZipFile			定义压缩文件
type ZipFile struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`                                                             // 压缩文件的id
	UserId    uint       `json:"user_id" gorm:"index:idx_userId;not null"`                                                        // 作者的id
	Title     string     `json:"title" gorm:"type:varchar(50);not null;index:idx_search,class:FULLTEXT,option:WITH PARSER ngram"` // 文章的标题
	Content   string     `json:"content" gorm:"type:text;not null;index:idx_search,class:FULLTEXT,option:WITH PARSER ngram"`      // 文章的内容
	ResLong   string     `json:"res_long" gorm:"type:text;index:idx_search,class:FULLTEXT,option:WITH PARSER ngram"`              // 备用长文本
	ResShort  string     `json:"res_short" gorm:"type:text;index:idx_search,class:FULLTEXT,option:WITH PARSER ngram"`             // 备用短文本
	Visible   int8       `json:"visible" gorm:"type:tinyint;default:1"`                                                           // 可见等级
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`                                                                // 创建时间
	UpdatedAt model.Time `json:"updated_at" gorm:"type:timestamp"`                                                                // 更新时间
}

// @title    BeforeCreate
// @description   计算出一个uint
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (zipFile *ZipFile) BeforeCreate(scope *gorm.DB) error {
	zipFile.ID = uuid.NewV4()
	return nil
}
