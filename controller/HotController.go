// @Title  HotController
// @Description  该文件用于提供关于热度的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"Blog/model"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IHotController			定义了热度类接口
type IHotController interface {
	ArticleVisit(ctx *gin.Context)     // 文章的游览次数获取
	PostVisit(ctx *gin.Context)        // 帖子的游览次数获取
	ZipfileVisit(ctx *gin.Context)     // 前端文件的游览次数获取
	ZipfileDownload(ctx *gin.Context)  // 前端文件的下载次数获取
	ZipfileUse(ctx *gin.Context)       // 前端文件的使用人数获取
	ZipfileComment(ctx *gin.Context)   // 前端文件的评论人数
	PostThread(ctx *gin.Context)       // 帖子的跟帖人数
	ArticleRanking(ctx *gin.Context)   // 文章的热度排行
	PostRanking(ctx *gin.Context)      // 帖子的热度排行
	ZipfileRanking(ctx *gin.Context)   // 前端文件的热度排行
	UserRanking(ctx *gin.Context)      // 用户的热度排行
	UserLevel(ctx *gin.Context)        // 查看用户的热度等级
	MyLevel(ctx *gin.Context)          // 用户查看自己的热度等级
	PowerPoint(ctx *gin.Context)       // 查看用户的简报
	MyPowerPoint(ctx *gin.Context)     // 用户查看自己的简报
	ArticleRecomment(ctx *gin.Context) // 文章推荐
	PostRecomment(ctx *gin.Context)    // 帖子推荐
	ZipfileRecomment(ctx *gin.Context) // 前端文件推荐
	UserRecomment(ctx *gin.Context)    // 用户推荐
}

// HotController			定义了热度工具类
type HotController struct {
	DB *gorm.DB //包含一个数据库指针
}

//
var HotLevel []float64 = []float64{
	200, 400, 800, 1600, 3200, 6400, 12800, 25600,
}

// @title    ArticleVisit
// @description   用户获取文章的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ArticleVisit(ctx *gin.Context) {
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

// @title    PostVisit
// @description   用户获取帖子的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) PostVisit(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否有权限
	if post.UserId != user.ID && (post.Visible == 3 || (post.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(post.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"views": util.CardS(3, "C"+postId)}, "查看游览次数成功")
}

// @title    ZipfileVisit
// @description   用户获取前端文件的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileVisit(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否有权限
	if zipfile.UserId != user.ID && (zipfile.Visible == 3 || (zipfile.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(zipfile.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"views": util.CardS(2, "C"+zipfileId)}, "查看游览次数成功")
}

// @title    ZipfileDownload
// @description   用户获取前端文件的下载次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileDownload(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否有权限
	if zipfile.UserId != user.ID && (zipfile.Visible == 3 || (zipfile.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(zipfile.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"downloads": util.CardS(2, "T"+zipfileId)}, "查看下载次数成功")
}

// @title    ZipfileUse
// @description   用户获取前端文件的使用次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileUse(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否有权限
	if zipfile.UserId != user.ID && (zipfile.Visible == 3 || (zipfile.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(zipfile.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"uses": util.CardS(2, "U"+zipfileId)}, "查看使用人数成功")
}

// @title    ZipfileComment
// @description   用户获取前端文件的评论人数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileComment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否有权限
	if zipfile.UserId != user.ID && (zipfile.Visible == 3 || (zipfile.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(zipfile.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"comments": util.CardS(2, "M"+zipfileId)}, "查看评论人数成功")
}

// @title    PostThread
// @description   用户获取帖子的跟帖人数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) PostThread(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否有权限
	if post.UserId != user.ID && (post.Visible == 3 || (post.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(post.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"threads": util.CardS(4, "M"+postId)}, "查看跟帖人数成功")
}

// @title    ArticleRanking
// @description   查看文章的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ArticleRanking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	articles := util.RangeZ(1, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(1, "H")

	response.Success(ctx, gin.H{"articles": articles, "total": total}, "查看文章排行成功")
}

// @title    PostRanking
// @description   查看帖子的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) PostRanking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	posts := util.RangeZ(3, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(3, "H")

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "查看帖子排行成功")
}

// @title    UserRanking
// @description   查看前端文件的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) UserRanking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.RangeZ(4, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(4, "H")

	response.Success(ctx, gin.H{"users": users, "total": total}, "查看用户排行成功")
}

// @title    ZipfileRanking
// @description   查看前端文件的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileRanking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	zipfiles := util.RangeZ(2, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(2, "H")

	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "查看前端文件排行成功")
}

// @title    PowerPoint
// @description   查看用户简报
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) PowerPoint(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	// TODO 获取path中的id
	userId := ctx.Params.ByName("id")

	// TODO 判断当前用户是否为所查用户
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '3' && userId != strconv.Itoa(int(user.ID)) {
		response.Fail(ctx, nil, "权限不足，无法查看")
		return
	}

	var powerpoints []model.PowerPoint

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", userId, start, end).Find(&powerpoints)

	response.Success(ctx, gin.H{"powerpoints": powerpoints}, "查看用户简报成功")
}

// @title    MyPowerPoint
// @description   用户查看自身简报
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) MyPowerPoint(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取起始时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var powerpoints []model.PowerPoint

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&powerpoints)

	response.Success(ctx, gin.H{"powerpoints": powerpoints}, "查看用户简报成功")
}

// @title    ArticleRecomment
// @description   查看文章推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ArticleRecomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	articles := util.RangeZ(1, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(1, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"articles": articles, "total": total}, "查看文章推荐成功")
}

// @title    PostRecomment
// @description   查看帖子推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) PostRecomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	posts := util.RangeZ(3, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(3, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "查看帖子推荐成功")
}

// @title    ZipfileRecomment
// @description   查看前端文件推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) ZipfileRecomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	zipfiles := util.RangeZ(2, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(2, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "查看前端文件推荐成功")
}

// @title    UsersRecomment
// @description   查看用户推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) UserRecomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.RangeZ(4, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(4, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"users": users, "total": total}, "查看用户推荐成功")
}

// @title    UserLevel
// @description   查看用户等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) UserLevel(ctx *gin.Context) {

	// TODO 获取path中的id
	userId := ctx.Params.ByName("id")

	Level := 0

	score := util.ScoreZ(4, "H", userId)

	for i := 7; i >= 0; i-- {
		if score >= HotLevel[i] {
			Level = i + 1
			break
		}
	}

	response.Success(ctx, gin.H{"Level": Level}, "查看用户热度等级成功")
}

// @title    MyLevel
// @description   用户查看自身等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HotController) MyLevel(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	Level := 0

	score := util.ScoreZ(4, "H", strconv.Itoa(int(user.ID)))

	for i := 7; i >= 0; i-- {
		if score >= HotLevel[i] {
			Level = i + 1
			break
		}
	}

	response.Success(ctx, gin.H{"Level": Level}, "查看用户热度等级成功")
}

// @title    NewHotController
// @description   新建一个热度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewHotController() IHotController {
	db := common.GetDB()
	db.AutoMigrate(model.PowerPoint{})
	return HotController{DB: db}
}
