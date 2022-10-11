// @Title ArticleVisibleController
// @Description  该文件用于提供文章可见度的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IArticleVisibleController			定义了文章可见度类接口
type IArticleVisibleController interface {
	Interface.VisibleInterface // 可见度的相关方法
}

// ArticleVisibleController			定义了可见度工具类
type ArticleVisibleController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   设置文章可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v ArticleVisibleController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	articleId := ctx.Params.ByName("id")

	// TODO 获取可见度等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var article gmodel.Article

	// TODO 查看文章是否存在
	if v.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 查看是否为文章作者
	if article.UserId != user.ID {
		response.Fail(ctx, nil, "文章不属于你，请勿非法操作")
		return
	}

	// TODO 查看值是否合法
	if Level > 4 {
		response.Fail(ctx, nil, "等级不合法")
		return
	}

	article.Visible = int8(Level)

	v.DB.Save(article)

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    Show
// @description   查看文章可见等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v ArticleVisibleController) Show(ctx *gin.Context) {

	// 获取path中的id
	articleId := ctx.Params.ByName("id")

	var article gmodel.Article

	// TODO 查看文章是否存在
	if v.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"Level": article.Visible}, "查看成功")
	return
}

// @title    NewArticleVisibleController
// @description   新建一个设置可见度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleVisibleController() IArticleVisibleController {
	db := common.GetDB()
	return ArticleVisibleController{DB: db}
}
