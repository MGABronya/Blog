// @Title  PostFavoriteController
// @Description  该文件用于提供帖子收藏操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IPostFavoriteController			定义了帖子收藏类接口
type IPostFavoriteController interface {
	Interface.FavoriteInterface // 收藏相关方法
}

// PostFavoriteController			定义了帖子收藏工具类
type PostFavoriteController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建帖子收藏
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostFavoriteController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看文章是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	util.SetS(3, "paF"+postId, strconv.Itoa(int(user.ID)))
	util.SetS(3, "pFa"+strconv.Itoa(int(user.ID)), postId)

	response.Success(ctx, nil, "收藏成功")
}

// @title    Show
// @description   查看帖子是否收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostFavoriteController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(3, "paF"+postId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消帖子的收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostFavoriteController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	util.RemS(3, "paF"+postId, strconv.Itoa(int(user.ID)))
	util.RemS(3, "pFa"+strconv.Itoa(int(user.ID)), postId)
	response.Success(ctx, nil, "删除成功")
}

// @title    FavoriteList
// @description   显示帖子的收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostFavoriteController) FavoriteList(ctx *gin.Context) {

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	es := util.MembersS(3, "paF"+postId)
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    UserFavoriteList
// @description   显示用户的帖子收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostFavoriteController) UserFavoriteList(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	es := util.MembersS(3, "pFa"+strconv.Itoa(int(user.ID)))
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    NewPostFavoriteController
// @description   新建一个帖子收藏的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostFavoriteController() IPostFavoriteController {
	db := common.GetDB()
	return PostFavoriteController{DB: db}
}
