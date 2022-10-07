// @Title  history
// @Description  定义历史记录
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// ArticleHistory			定义文章游览记录
type ArticleHistory struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`            // 历史记录id
	UserId    uint       `json:"user_id" gorm:"index:idx_userId;not null"`       // 作者的id
	ArticleId string     `json:"article_id" gorm:"index:idx_articleId;not null"` // 文章的id
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`               // 创建时间
}

// ZipfileHistory			定义前端文件游览记录
type ZipfileHistory struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`            // 历史记录id
	UserId    uint       `json:"user_id" gorm:"index:idx_userId;not null"`       // 作者的id
	ZipfileId string     `json:"zipfile_id" gorm:"index:idx_zipfileId;not null"` // 前端文件的id
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`               // 创建时间
}

// PostHistory			定义帖子游览记录
type PostHistory struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`      // 历史记录id
	UserId    uint       `json:"user_id" gorm:"index:idx_userId;not null"` // 作者的id
	PostId    string     `json:"post_id" gorm:"index:idx_postId;not null"` // 帖子的id
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`         // 创建时间
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (articleHistory *ArticleHistory) BeforeCreate(scope *gorm.DB) error {
	articleHistory.ID = uuid.NewV4()
	return nil
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (zipfileHistory *ZipfileHistory) BeforeCreate(scope *gorm.DB) error {
	zipfileHistory.ID = uuid.NewV4()
	return nil
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (postHistory *PostHistory) BeforeCreate(scope *gorm.DB) error {
	postHistory.ID = uuid.NewV4()
	return nil
}
