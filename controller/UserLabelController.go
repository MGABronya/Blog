// @Title  UserLableController
// @Description  该文件用于提供用户标签操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"Blog/vo"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IUserLabelController			定义了用户标签类接口
type IUserLabelController interface {
	Interface.LabelInterface      //创查删
	DeleteLabel(ctx *gin.Context) //删除指定用户标签
	ShowLabel(ctx *gin.Context)   // 查看指定用户标签
}

// UserLabelController			定义了用户标签工具类
type UserLabelController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建用户标签
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserLabelController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	util.SetS(4, "aL"+strconv.Itoa(int(user.ID)), requestLabel.Label)
	util.SetS(4, "La"+requestLabel.Label, strconv.Itoa(int(user.ID)))

	response.Success(ctx, nil, "设置成功")
}

// @title    Show
// @description   查看用户的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserLabelController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	response.Success(ctx, gin.H{"labels": util.GetS(4, "aL"+strconv.Itoa(int(user.ID)))}, "查找成功")
}

// @title    ShowLabel
// @description   查看指定用户的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserLabelController) ShowLabel(ctx *gin.Context) {

	userId := ctx.Params.ByName("id")

	response.Success(ctx, gin.H{"labels": util.GetS(4, "aL"+userId)}, "查找成功")
}

// @title    Delete
// @description   删除用户的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserLabelController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	util.RemS(4, "aL"+strconv.Itoa(int(user.ID)), requestLabel.Label)
	util.RemS(4, "La"+requestLabel.Label, strconv.Itoa(int(user.ID)))

	if util.CardS(4, "La"+requestLabel.Label) == 0 {
		util.Del(4, "La"+requestLabel.Label)
	}

	response.Success(ctx, nil, "删除成功")
}

// @title    DeleteLabel
// @description   删除用户的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (u UserLabelController) DeleteLabel(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")

	var user model.User

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	// TODO 查看用户是否存在
	if u.DB.Where("id = ?", userId).First(&user).RecordNotFound() {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 判断当前用户是有权限修改
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && strconv.Itoa(int(user.ID)) != userId {
		response.Fail(ctx, nil, "权限不足")
		return
	}

	util.RemS(4, "aL"+userId, requestLabel.Label)
	util.RemS(4, "La"+requestLabel.Label, userId)
	response.Success(ctx, nil, "删除成功")
}

// @title    NewUserLabelController
// @description   新建一个用户标签的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewUserLabelController() IUserLabelController {
	db := common.GetDB()
	return UserLabelController{db}
}
