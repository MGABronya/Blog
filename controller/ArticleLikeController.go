// @Title  ArticleLikeController
// @Description  该文件用于提供文章点赞操作的各种函数
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

// IArticleLikeController			定义了文章点赞类接口
type IArticleLikeController interface {
	Interface.LikeInterface // 点赞相关方法
}

// ArticleLikeController			定义了文章点赞工具类
type ArticleLikeController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建文章点赞
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLikeController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	util.SetS(1, "iL"+articleId, strconv.Itoa(int(user.ID)))

	response.Success(ctx, nil, "点赞成功")
}

// @title    Show
// @description   查看文章是否点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLikeController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(1, "iL"+articleId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消文章的点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLikeController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	util.RemS(1, "iL"+articleId, strconv.Itoa(int(user.ID)))
	response.Success(ctx, nil, "删除成功")
}

// @title    LikeList
// @description   显示文章的点赞列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLikeController) LikeList(ctx *gin.Context) {

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).RecordNotFound() {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	es := util.MembersS(1, "iL"+articleId)
	total := len(es)
	response.Success(ctx, gin.H{"Liks": es, "total": total}, "查看成功")
}

// @title    NewArticleLikeController
// @description   新建一个文章点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleLikeController() IArticleLikeController {
	db := common.GetDB()
	return ArticleLikeController{DB: db}
}
