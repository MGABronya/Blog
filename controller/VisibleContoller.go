// @Title VisibleController
// @Description  该文件用于提供可见度/可评论的各种函数
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

// IVisibleController			定义了可见度类接口
type IVisibleController interface {
	ArticleCreate(ctx *gin.Context)       // 设置文章可见度
	PostCreate(ctx *gin.Context)          // 设置帖子可见度
	ZipFileCreate(ctx *gin.Context)       // 设置前端文件可见度
	ArticleShow(ctx *gin.Context)         // 查看文章可见度
	PostShow(ctx *gin.Context)            // 查看帖子可见度
	ZipFileShow(ctx *gin.Context)         // 查看前端文件可见度
	PostThread(ctx *gin.Context)          // 设置帖子是否可以跟帖
	ZipfileComment(ctx *gin.Context)      // 设置前端文件是否可以评论
	PostThreadShow(ctx *gin.Context)      // 查看帖子是否可以跟帖
	ZipfileCommentShow(ctx *gin.Context)  // 查看前端文件是否可以评论
	ZipfileDownload(ctx *gin.Context)     // 设置前端文件是否可以被下载
	ZipfileDownloadShow(ctx *gin.Context) // 查看前端文件下载等级
	PostThreadCan(ctx *gin.Context)       // 查看帖子是否可以跟帖
	ZipfileCan(ctx *gin.Context)          // 查看前端文件是否可以下载
	ZipfileCommentCan(ctx *gin.Context)   // 查看前端文件是否可以评论
}

// VisibleController			定义了可见度工具类
type VisibleController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    ArticleCreate
// @description   设置文章可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ArticleCreate(ctx *gin.Context) {
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

// @title    PostCreate
// @description   设置帖子可见度
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) PostCreate(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	// TODO 获取可见度等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var post gmodel.Post

	// TODO 查看文章是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，请勿非法操作")
		return
	}

	// TODO 查看值是否合法
	if Level > 4 {
		response.Fail(ctx, nil, "等级不合法")
		return
	}

	post.Visible = int8(Level)

	v.DB.Save(post)

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    ZipFileCreate
// @description   设置当前文件可见度
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipFileCreate(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	fileId := ctx.Params.ByName("id")

	// TODO 获取可见度等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", fileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否为前端文件作者
	if zipfile.UserId != user.ID {
		response.Fail(ctx, nil, "前端文件不属于你，请勿非法操作")
		return
	}

	// TODO 查看值是否合法
	if Level > 4 {
		response.Fail(ctx, nil, "等级不合法")
		return
	}

	zipfile.Visible = int8(Level)

	v.DB.Save(zipfile)

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    ArticleShow
// @description   查看文章可见等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ArticleShow(ctx *gin.Context) {

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

// @title    PostShow
// @description   查看帖子可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) PostShow(ctx *gin.Context) {

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"Level": post.Visible}, "查看成功")
	return
}

// @title    ZipFileShow
// @description   查看前端文件可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipFileShow(ctx *gin.Context) {

	// 获取path中的id
	fileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", fileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	response.Success(ctx, gin.H{"Level": zipfile.Visible}, "查看成功")
	return
}

// @title    PostThread
// @description   设置帖子是否可以跟帖
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) PostThread(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	// TODO 获取可评论等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，无法设置")
		return
	}

	util.SetH(3, "W", postId, strconv.Itoa(Level))

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    ZipfileComment
// @description   设置前端文件是否可以评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileComment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	filezipId := ctx.Params.ByName("id")

	// TODO 获取可评论等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", filezipId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否为前端文件作者
	if zipfile.UserId != user.ID {
		response.Fail(ctx, nil, "前端文件不属于你，无法设置")
		return
	}

	util.SetH(2, "W", filezipId, strconv.Itoa(Level))

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    PostThreadShow
// @description   查看帖子的跟帖等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) PostThreadShow(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否为帖子作者
	if post.UserId != user.ID {
		response.Fail(ctx, nil, "帖子不属于你，无法查看")
		return
	}

	response.Success(ctx, gin.H{"level": util.GetH(3, "W", postId)}, "查看成功")
	return
}

// @title    ZipfileCommentShow
// @description   查看评论的等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileCommentShow(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", zipfile).First(&zipfileId).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否为前端文件作者
	if zipfile.UserId != user.ID {
		response.Fail(ctx, nil, "前端文件不属于你，无法查看")
		return
	}

	response.Success(ctx, gin.H{"level": util.GetH(2, "W", zipfileId)}, "查看成功")
	return
}

// @title    ZipfileDownload
// @description   设置前端文件是否可以下载
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileDownload(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	filezipId := ctx.Params.ByName("id")

	// TODO 获取可评论等级
	Level, _ := strconv.Atoi(ctx.DefaultQuery("Level", "1"))

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", filezipId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否为前端文件作者
	if zipfile.UserId != user.ID {
		response.Fail(ctx, nil, "前端文件不属于你，无法设置")
		return
	}

	util.SetH(2, "D", filezipId, strconv.Itoa(Level))

	response.Success(ctx, nil, "设置成功")
	return
}

// @title    ZipfileDownloadShow
// @description   查看前端文件下载等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileDownloadShow(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	filezipId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", filezipId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否为前端文件作者
	if zipfile.UserId != user.ID {
		response.Fail(ctx, nil, "前端文件不属于你，无法查看")
		return
	}

	response.Success(ctx, gin.H{"Level": util.GetH(2, "W", filezipId)}, "查看成功")
	return
}

// @title    PostThreadCan
// @description   查看帖子是否可以跟帖
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) PostThreadCan(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否存在
	if v.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	flag := util.PostThread(strconv.Itoa(int(user.ID)), postId, strconv.Itoa(int(post.UserId)))

	response.Success(ctx, gin.H{"flag": flag}, "查看成功")
	return
}

// @title    ZipfileCommentCan
// @description   查看前端文件是否可以评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileCommentCan(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	flag := util.ZipfileComment(strconv.Itoa(int(user.ID)), zipfileId, strconv.Itoa(int(zipfile.UserId)))

	response.Success(ctx, gin.H{"flag": flag}, "查看成功")
	return
}

// @title    ZipfileCan
// @description   查看前端文件是否可以下载
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v VisibleController) ZipfileCan(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// 获取path中的id
	zipfileId := ctx.Params.ByName("id")

	var zipfile model.ZipFile

	// TODO 查看前端文件是否存在
	if v.DB.Where("id = ?", zipfileId).First(&zipfile).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	flag := util.Zipfile(strconv.Itoa(int(user.ID)), zipfileId, strconv.Itoa(int(zipfile.UserId)))

	response.Success(ctx, gin.H{"flag": flag}, "查看成功")
	return
}

// @title    NewVisibleController
// @description   新建一个设置可见度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewVisibleController() IVisibleController {
	db := common.GetDB()
	return VisibleController{DB: db}
}
