// @Title  SystemController
// @Description  该文件用于提供后台操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"Blog/util"
	"ginEssential/response"
	"strconv"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
)

// ISystemController			定义了后台类接口
type ISystemController interface {
	Permission(ctx *gin.Context)     // 实现定义用户权限
	ShowPermission(ctx *gin.Context) // 实现查看当前用户权限
}

// SystemController			定义了后台工具类
type SystemController struct {
}

// @title    Permission
// @description   修改一个用户的用户等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SystemController) Permission(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)
	id := ctx.Params.ByName("id")
	level := ctx.Params.ByName("level")
	fromLevel := util.GetH(0, "permission", strconv.Itoa(int(user.ID)))
	toLevel := util.GetH(0, "permission", id)
	if toLevel >= fromLevel || level >= fromLevel {
		response.Fail(ctx, nil, "权限不足")
		return
	}
	util.SetH(0, "permission", strconv.Itoa(int(user.ID)), level)
	response.Success(ctx, nil, "设置成功")
}

// @title    ShowPermission
// @description   用户查看自己的权限等级
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SystemController) ShowPermission(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)
	level := util.GetH(0, "permission", strconv.Itoa(int(user.ID)))
	response.Success(ctx, gin.H{"level": level}, "设置成功")
}

// @title    NewFileController
// @description   新建一个前端文件的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewSystemController() ISystemController {
	return SystemController{}
}
