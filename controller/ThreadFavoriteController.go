// @Title  ThreadFavoriteController
// @Description  该文件用于提供跟帖收藏操作的各种函数
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

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IThreadFavoriteController			定义了跟帖收藏类接口
type IThreadFavoriteController interface {
	Interface.FavoriteInterface // 收藏相关方法
}

// ThreadFavoriteController			定义了跟帖收藏工具类
type ThreadFavoriteController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建跟帖收藏
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadFavoriteController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).Error != nil {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	// TODO 查看帖子是否存在
	var post gmodel.Post
	if t.DB.Where("id = ?", thread.PostId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看跟帖是否已经收藏
	if util.IsS(3, "taF"+threadId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "跟帖已收藏")
		return
	}

	util.IncrByZ(3, "H", thread.PostId, 25)
	util.IncrByZ(3, "TH", threadId, 50)
	util.IncrByZ(4, "H", strconv.Itoa(int(thread.UserId)), 25)
	util.IncrByZ(4, "H", strconv.Itoa(int(post.UserId)), 50)

	util.SetS(3, "taF"+threadId, strconv.Itoa(int(user.ID)))
	util.SetS(3, "tFa"+strconv.Itoa(int(user.ID)), threadId)

	response.Success(ctx, nil, "收藏成功")
}

// @title    Show
// @description   查看跟帖是否收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p ThreadFavoriteController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if p.DB.Where("id = ?", threadId).First(&thread).Error != nil {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(3, "taF"+threadId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消跟帖的收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadFavoriteController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).Error != nil {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	// TODO 查看帖子是否存在
	var post gmodel.Post
	if t.DB.Where("id = ?", thread.PostId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看跟帖是否已经收藏
	if !util.IsS(3, "taF"+threadId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "跟帖未收藏")
		return
	}

	util.IncrByZ(3, "H", thread.PostId, -25)
	util.IncrByZ(3, "TH", threadId, -50)
	util.IncrByZ(4, "H", strconv.Itoa(int(thread.UserId)), -25)
	util.IncrByZ(4, "H", strconv.Itoa(int(post.UserId)), -50)

	util.RemS(3, "taF"+threadId, strconv.Itoa(int(user.ID)))
	util.RemS(3, "tFa"+strconv.Itoa(int(user.ID)), threadId)
	response.Success(ctx, nil, "删除成功")
}

// @title    FavoriteList
// @description   显示跟帖的收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadFavoriteController) FavoriteList(ctx *gin.Context) {

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).Error != nil {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	es := util.MembersS(3, "taF"+threadId)
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    UserFavoriteList
// @description   显示用户的跟帖收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadFavoriteController) UserFavoriteList(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	es := util.MembersS(3, "tFa"+strconv.Itoa(int(user.ID)))
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    NewThreadFavoriteController
// @description   新建一个帖子收藏的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewThreadFavoriteController() IThreadFavoriteController {
	db := common.GetDB()
	return ThreadFavoriteController{DB: db}
}
