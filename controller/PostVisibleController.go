// @Title PostVisibleController
// @Description  该文件用于提供帖子可见度/可跟帖的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IPostVisibleController			定义了帖子可见度类接口
type IPostVisibleController interface {
	Interface.VisibleInterface   // 可见度方法
	Thread(ctx *gin.Context)     // 设置帖子是否可以跟帖
	ThreadShow(ctx *gin.Context) // 查看帖子跟帖等级
	ThreadCan(ctx *gin.Context)  // 查看帖子是否可以跟帖
}

// PostVisibleController			定义了可见度工具类
type PostVisibleController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   设置帖子可见度
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v PostVisibleController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	// TODO 获取可见度等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var post gmodel.Post

	// TODO 查看文章是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，请勿非法操作")
		return
	}

	// TODO 查看值是否合法
	if Level > 4 {
		response.Fail(ctx, nil, "等级不合法")
		return
	}

	post.Visible = int8(Level)

	v.DB.Save(post)

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    Show
// @description   查看帖子可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v PostVisibleController) Show(ctx *gin.Context) {

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"Level": post.Visible}, "查看成功")
	return
}

// @title    Thread
// @description   设置帖子是否可以跟帖
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v PostVisibleController) Thread(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	// TODO 获取可评论等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，无法设置")
		return
	}

	util.SetH(3, "W", postId, strconv.Itoa(Level))

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    ThreadShow
// @description   查看帖子的跟帖等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v PostVisibleController) ThreadShow(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，无法查看")
		return
	}

	response.Success(ctx, gin.H{"level": util.GetH(3, "W", postId)}, "查看成功")
	return
}

// @title    ThreadCan
// @description   查看帖子是否可以跟帖
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v PostVisibleController) ThreadCan(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	flag := util.PostThread(strconv.Itoa(int(user.ID)), postId, strconv.Itoa(int(post.UserId)))

	response.Success(ctx, gin.H{"flag": flag}, "查看成功")
	return
}

// @title    NewPostVisibleController
// @description   新建一个设置可见度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostVisibleController() IPostVisibleController {
	db := common.GetDB()
	return PostVisibleController{DB: db}
}
