// @Title  zipfile
// @Description  定义压缩文件
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:46
package model

import (
	"ginEssential/model"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// ZipFile			定义压缩文件
type ZipFile struct {
	ID        uuid.UUID  `json:"id" gorm:"type:char(36);primary_key"`    // 压缩文件的id
	UserId    uint       `json:"user_id" gorm:"not null"`                // 作者的id
	Title     string     `json:"title" gorm:"type:varchar(50);not null"` // 标题
	Remark    string     `json:"remark" gorm:"type:text;not null"`       // 备注
	ResLong   string     `json:"res_long" gorm:"type:text;"`             // 备用长文本
	ResShort  string     `json:"res_short" gorm:"type:text;"`            // 备用短文本
	Visible   int8       `json:"visible" gorm:"type:tinyint;default:1"`  // 可见等级
	CreatedAt model.Time `json:"created_at" gorm:"type:timestamp"`       // 创建时间
	UpdatedAt model.Time `json:"updated_at" gorm:"type:timestamp"`       // 更新时间
}

// @title    BeforeCreate
// @description   计算出一个uint
// @auth      MGAronya（张健）             2022-9-16 10:19
// @param     scope *gorm.Scope
// @return    error
func (zipfile *ZipFile) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
