// @Title  powerpoint
// @Description  定义用户简报
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// powerpoint			定义日报
type PowerPoint struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`      // 留言id
	UserId    uint       `json:"user_id" gorm:"index:idx_member;not null"` // 作者的id
	Label     string     `json:"label" gorm:"type:char(36);not null"`      // 标签
	Score     float64    `json:"score" gorm:"type:double;not null"`        // 分数
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`         // 创建时间
}

// @title    BeforeCreate
// @description   计算出一个uuid
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (powerPoint *PowerPoint) BeforeCreate(scope *gorm.DB) error {
	powerPoint.ID = uuid.NewV4()
	return nil
}
