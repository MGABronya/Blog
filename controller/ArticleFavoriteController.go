// @Title  ArticleFavoriteController
// @Description  该文件用于提供文章收藏操作的各种函数
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
	"gorm.io/gorm"
)

// IArticleFavoriteController			定义了文章收藏类接口
type IArticleFavoriteController interface {
	Interface.FavoriteInterface // 收藏相关方法
}

// ArticleFavoriteController			定义了文章收藏工具类
type ArticleFavoriteController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建文章收藏
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleFavoriteController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 查看文章是否已经收藏
	if util.IsS(1, "aF"+articleId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "文章已收藏")
		return
	}

	util.SetS(1, "aF"+articleId, strconv.Itoa(int(user.ID)))
	util.SetS(1, "Fa"+strconv.Itoa(int(user.ID)), articleId)

	// TODO 热度提升
	util.IncrByZ(1, "H", articleId, 50)
	util.IncrByZ(4, "H", strconv.Itoa(int(article.UserId)), 50)

	// TODO 用户标签分数提升
	labels := util.MembersS(1, "aL"+articleId)
	for _, label := range labels {
		util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), label, 50)
	}

	response.Success(ctx, nil, "收藏成功")
}

// @title    Show
// @description   查看文章是否收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleFavoriteController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(1, "aF"+articleId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消文章的收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleFavoriteController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 查看文章是否已经收藏
	if !util.IsS(1, "aF"+articleId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "文章未收藏")
		return
	}

	util.RemS(1, "aF"+articleId, strconv.Itoa(int(user.ID)))
	util.RemS(1, "Fa"+strconv.Itoa(int(user.ID)), articleId)

	// TODO 热度下降
	util.IncrByZ(1, "H", articleId, -50)
	util.IncrByZ(4, "H", strconv.Itoa(int(article.UserId)), -50)

	// TODO 用户标签分数下降
	labels := util.MembersS(1, "aL"+articleId)
	for _, label := range labels {
		util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), label, -50)
		if util.ScoreZ(4, "L"+strconv.Itoa(int(user.ID)), label) <= 0 {
			util.RemZ(4, "L"+strconv.Itoa(int(user.ID)), label)
		}
	}

	response.Success(ctx, nil, "删除成功")
}

// @title    FavoriteList
// @description   显示文章的收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleFavoriteController) FavoriteList(ctx *gin.Context) {

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	es := util.MembersS(1, "aF"+articleId)
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    UserFavoriteList
// @description   显示用户的文章收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleFavoriteController) UserFavoriteList(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	es := util.MembersS(1, "Fa"+strconv.Itoa(int(user.ID)))
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    NewArticleFavoriteController
// @description   新建一个文章收藏的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleFavoriteController() IArticleFavoriteController {
	db := common.GetDB()
	return ArticleFavoriteController{DB: db}
}
