// @Title  UserFindController
// @Description  该文件用于提供搜索用户的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"ginEssential/common"
	"ginEssential/response"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IUserFindController			定义了搜索用户类接口
type IUserFindController interface {
	NameFind(ctx *gin.Context)  // 通过名字查找用户
	EmailFind(ctx *gin.Context) // 通过邮箱查找用户
}

// UserFindController			定义了搜索用户工具类
type UserFindController struct {
	DB *gorm.DB // 包含一个数据库指针
}

// @title    NameFind
// @description   通过名字查找用户
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserFindController) NameFind(ctx *gin.Context) {

	name := ctx.Params.ByName("id")

	var user model.User

	// TODO 查看用户是否存在
	if u.DB.Where("name = ?", name).First(&user).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	response.Success(ctx, gin.H{"user": user}, "查找成功")
}

// @title    EmailFind
// @description   通过邮箱查找用户
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserFindController) EmailFind(ctx *gin.Context) {

	email := ctx.Params.ByName("id")

	var user model.User

	// TODO 查看用户是否存在
	if u.DB.Where("email = ?", email).First(&user).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	response.Success(ctx, gin.H{"user": user}, "查找成功")
}

// @title    NewUserFindController
// @description   新建一个搜寻用户的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewUserFindController() IUserFindController {
	db := common.GetDB()
	return UserFindController{DB: db}
}
