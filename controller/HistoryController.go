// @Title  HistoryController
// @Description  该文件用于提供历史记录操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	"Blog/model"
	"Blog/util"
	gmodel "ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IHistoryController			定义了历史记录类接口
type IHistoryController interface {
	ArticleCreate(ctx *gin.Context)    // 用户设置文章历史记录
	PostCreate(ctx *gin.Context)       // 用户设置帖子历史记录
	ZipfileCreate(ctx *gin.Context)    // 用户设置前端文件历史记录
	AriticleShow(ctx *gin.Context)     // 用户查看文章历史记录
	PostShow(ctx *gin.Context)         // 用户查看帖子历史记录
	ZipfileShow(ctx *gin.Context)      // 用户查看前端文件历史记录
	ArticleDelete(ctx *gin.Context)    // 用户删除文章历史记录
	PostDelete(ctx *gin.Context)       // 用户删除帖子历史记录
	ZipfileDelete(ctx *gin.Context)    // 用户删除前端文件历史记录
	ArticleDeleteAll(ctx *gin.Context) // 用户清空文章历史记录
	PostDeleteAll(ctx *gin.Context)    // 用户清空帖子历史记录
	ZipfileDeleteAll(ctx *gin.Context) // 用户清空前端文件历史记录
}

// HistoryController			定义了历史记录工具类
type HistoryController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    ArticleCreate
// @description   用户创建文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ArticleCreate(ctx *gin.Context) {
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

// @title    PostCreate
// @description   用户创建帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) PostCreate(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	postId := ctx.Params.ByName("id")

	postHistory := model.PostHistory{
		PostId: postId,
		UserId: user.ID,
	}

	var post gmodel.Post
	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 更新热度
	if !util.IsS(3, "C"+postId, strconv.Itoa(int(user.ID))) {
		util.SetS(3, "C"+postId, strconv.Itoa(int(user.ID)))
		util.IncrByZ(3, "L"+strconv.Itoa(int(user.ID)), postId, 1)
	}

	// TODO 更新历史记录
	h.DB.Create(&postHistory)

	response.Success(ctx, nil, "创建文章历史记录成功")
}

// @title    ZipfileCreate
// @description   用户创建前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ZipfileCreate(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	zipfileId := ctx.Params.ByName("id")

	zipfileHistory := model.ZipfileHistory{
		ZipfileId: zipfileId,
		UserId:    user.ID,
	}

	var zipfile model.ZipFile
	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 更新热度
	if !util.IsS(2, "C"+zipfileId, strconv.Itoa(int(user.ID))) {
		util.SetS(2, "C"+zipfileId, strconv.Itoa(int(user.ID)))
		util.IncrByZ(2, "H", zipfileId, 1)
		util.IncrByZ(4, "H", strconv.Itoa(int(zipfile.UserId)), 1)
		util.IncrByZ(2, "L"+strconv.Itoa(int(user.ID)), zipfile.ID.String(), 1)
	}

	// TODO 更新历史记录
	h.DB.Create(&zipfileHistory)

	response.Success(ctx, nil, "创建文章历史记录成功")
}

// @title    AriticleShow
// @description   用户查看文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) AriticleShow(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var articleHistorys []model.ArticleHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&articleHistorys)

	response.Success(ctx, gin.H{"articleHistorys": articleHistorys}, "查看文章历史记录成功")
}

// @title    PostShow
// @description   用户查看帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) PostShow(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var postHistorys []model.PostHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&postHistorys)

	response.Success(ctx, gin.H{"postHistorys": postHistorys}, "查看帖子历史记录成功")
}

// @title    ZipfileShow
// @description   用户查看帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ZipfileShow(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var zipfileHistorys []model.ZipfileHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&zipfileHistorys)

	response.Success(ctx, gin.H{"zipfileHistorys": zipfileHistorys}, "查看前端文件历史记录成功")
}

// @title    ArticleDelete
// @description   用户删除文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ArticleDelete(ctx *gin.Context) {
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

// @title    PostDelete
// @description   用户删除帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) PostDelete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	postId := ctx.Params.ByName("id")

	var postHistory model.PostHistory

	// TODO 查看历史记录是否存在
	if h.DB.Where("id = ?", postId).First(&postHistory).Error != nil {
		response.Fail(ctx, nil, "该记录不存在")
		return
	}

	// TODO 判断当前用户是否为历史记录的作者
	if user.ID != postHistory.UserId {
		response.Fail(ctx, nil, "该条记录不属于您，请勿非法操作")
		return
	}

	h.DB.Delete(&postHistory)

	response.Success(ctx, nil, "删除帖子历史记录成功")
}

// @title    ZipfileDelete
// @description   用户删除前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ZipfileDelete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取id
	zipfileId := ctx.Params.ByName("id")

	var zipfileHistory model.ZipfileHistory

	// TODO 查看历史记录是否存在
	if h.DB.Where("id = ?", zipfileId).First(&zipfileHistory).Error != nil {
		response.Fail(ctx, nil, "该记录不存在")
		return
	}

	// TODO 判断当前用户是否为历史记录的作者
	if user.ID != zipfileHistory.UserId {
		response.Fail(ctx, nil, "该条记录不属于您，请勿非法操作")
		return
	}

	h.DB.Delete(&zipfileHistory)

	response.Success(ctx, nil, "删除帖子历史记录成功")
}

// @title    ArticleDeleteAll
// @description   用户清空文章历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ArticleDeleteAll(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	h.DB.Where("user_id = ?", user.ID).Delete(model.ArticleHistory{})

	response.Success(ctx, nil, "清空文章历史记录成功")
}

// @title    PostDeleteAll
// @description   用户清空帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) PostDeleteAll(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	h.DB.Where("user_id = ?", user.ID).Delete(model.PostHistory{})

	response.Success(ctx, nil, "清空帖子历史记录成功")
}

// @title    ZipfileDeleteAll
// @description   用户清空前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h HistoryController) ZipfileDeleteAll(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	h.DB.Where("user_id = ?", user.ID).Delete(model.ZipfileHistory{})

	response.Success(ctx, nil, "清空前端文件历史记录成功")
}

// @title    NewHistoryController
// @description   新建一个历史记录的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewHistoryController() IHistoryController {
	db := common.GetDB()
	return HistoryController{db}
}
