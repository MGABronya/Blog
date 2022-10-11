// @Title  ArticleHistoryController
// @Description  该文件用于提供文章历史记录操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	gmodel "ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IArticleHistoryController			定义了文章历史记录类接口
type IArticleHistoryController interface {
	Interface.HistoryInterface // 历史记录功能
}

// ArticleHistoryController			定义了历史记录工具类
type ArticleHistoryController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   用户创建文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHistoryController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	articleId := ctx.Params.ByName("id")

	articleHistory := model.ArticleHistory{
		ArticleId: articleId,
		UserId:    user.ID,
	}

	var article gmodel.Article
	// TODO 查看文章是否在数据库中存在
	if h.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 更新热度
	if !util.IsS(1, "C"+articleId, strconv.Itoa(int(user.ID))) {
		util.SetS(1, "C"+articleId, strconv.Itoa(int(user.ID)))
		util.IncrByZ(1, "H", articleId, 1)
		util.IncrByZ(4, "H", strconv.Itoa(int(article.UserId)), 1)
		util.IncrByZ(1, "L"+strconv.Itoa(int(user.ID)), articleId, 1)
	}

	// TODO 更新历史记录
	h.DB.Create(&articleHistory)

	response.Success(ctx, nil, "创建文章历史记录成功")
}

// @title    Show
// @description   用户查看文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHistoryController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var articleHistorys []model.ArticleHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&articleHistorys)

	response.Success(ctx, gin.H{"articleHistorys": articleHistorys}, "查看文章历史记录成功")
}

// @title    Delete
// @description   用户删除文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHistoryController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	articleId := ctx.Params.ByName("id")

	var articleHistory model.ArticleHistory

	// TODO 查看历史记录是否存在
	if h.DB.Where("id = ?", articleId).First(&articleHistory).Error != nil {
		response.Fail(ctx, nil, "该记录不存在")
		return
	}

	// TODO 判断当前用户是否为历史记录的作者
	if user.ID != articleHistory.UserId {
		response.Fail(ctx, nil, "该条记录不属于您，请勿非法操作")
		return
	}

	h.DB.Delete(&articleHistory)

	response.Success(ctx, nil, "删除文章历史记录成功")
}

// @title    DeleteAll
// @description   用户清空文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHistoryController) DeleteAll(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	h.DB.Where("user_id = ?", user.ID).Delete(model.ArticleHistory{})

	response.Success(ctx, nil, "清空文章历史记录成功")
}

// @title    NewArticleHistoryController
// @description   新建一个文章历史记录的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleHistoryController() IArticleHistoryController {
	db := common.GetDB()
	db.AutoMigrate(model.ArticleHistory{})
	return ArticleHistoryController{db}
}
