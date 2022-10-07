// @Title  ArticleLableController
// @Description  该文件用于提供文章标签操作的各种函数
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
	"gorm.io/gorm"
)

// IArticleLabelController			定义了文章标签类接口
type IArticleLabelController interface {
	Interface.LabelInterface //标签相关方法
}

// ArticleLabelController			定义了文章标签工具类
type ArticleLabelController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建文章标签
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLabelController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 判断当前用户是否为文章的作者
	if user.ID != article.UserId {
		response.Fail(ctx, nil, "文章不属于您，请勿非法操作")
		return
	}

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	if util.IsS(1, "aL"+articleId, requestLabel.Label) {
		response.Fail(ctx, nil, "标签已设置")
		return
	}

	// TODO 用户标签分数上升
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label, 30)

	util.SetS(1, "aL"+articleId, requestLabel.Label)
	util.SetS(1, "La"+requestLabel.Label, articleId)

	response.Success(ctx, nil, "设置成功")
}

// @title    Show
// @description   查看文章的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLabelController) Show(ctx *gin.Context) {

	articleId := ctx.Params.ByName("id")

	var article model.Article

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"labels": util.MembersS(1, "aL"+articleId)}, "查找成功")
}

// @title    Delete
// @description   删除文章的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (a ArticleLabelController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	articleId := ctx.Params.ByName("id")

	var article model.Article

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	// TODO 查看文章是否存在
	if a.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 判断当前用户是否为文章的作者
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != article.UserId {
		response.Fail(ctx, nil, "文章不属于您，请勿非法操作")
		return
	}

	if !util.IsS(1, "aL"+articleId, requestLabel.Label) {
		response.Fail(ctx, nil, "标签未设置")
		return
	}

	// TODO 用户标签分数下降
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label, -30)
	if util.ScoreZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label) <= 0 {
		util.RemZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label)
	}

	util.RemS(1, "aL"+articleId, requestLabel.Label)
	util.RemS(1, "La"+requestLabel.Label, articleId)

	if util.CardS(1, "La"+requestLabel.Label) == 0 {
		util.Del(1, "La"+requestLabel.Label)
	}

	response.Success(ctx, nil, "删除成功")
}

// @title    NewArticleLabelController
// @description   新建一个文章标签的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleLabelController() IArticleLabelController {
	db := common.GetDB()
	return ArticleLabelController{DB: db}
}
