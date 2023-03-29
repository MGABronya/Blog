// @Title  ArticleHotController
// @Description  该文件用于提供关于文章热度的各种函数
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

// IArticleHotController			定义了热度类接口
type IArticleHotController interface {
	Interface.HotInterface  // 热度相关方法
	Visit(ctx *gin.Context) // 文章的游览次数获取
}

// ArticleHotController			定义了热度工具类
type ArticleHotController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Visit
// @description   用户获取文章的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHotController) Visit(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	articleId := ctx.Params.ByName("id")

	var article gmodel.Article

	// TODO 查看文章是否在数据库中存在
	if h.DB.Where("id = ?", articleId).First(&article).Error != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// TODO 查看是否有权限
	if article.UserId != user.ID && (article.Visible == 3 || (article.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(article.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"views": util.CardS(1, "C"+articleId)}, "查看游览次数成功")
}

// @title    Ranking
// @description   查看文章的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHotController) Ranking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	articles := util.RangeZ(1, "H", int64(pageNum-1)*int64(pageSize), int64(pageNum-1)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(1, "H")

	response.Success(ctx, gin.H{"articles": articles, "total": total}, "查看文章排行成功")
}

// @title    Recomment
// @description   查看文章推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h ArticleHotController) Recomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	articles := util.RangeZ(1, "R"+strconv.Itoa(int(user.ID)), int64(pageNum-1)*int64(pageSize), int64(pageNum-1)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(1, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"articles": articles, "total": total}, "查看文章推荐成功")
}

// @title    NewArticleHotController
// @description   新建一个文章热度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleHotController() IArticleHotController {
	db := common.GetDB()
	return ArticleHotController{DB: db}
}
