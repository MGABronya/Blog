// @Title FileVisibleController
// @Description  该文件用于提供前端文件可见度的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IFileVisibleController			定义了可见度类接口
type IFileVisibleController interface {
	Interface.VisibleInterface     // 可见度的相关方法
	Comment(ctx *gin.Context)      // 设置前端文件是否可以评论
	CommentShow(ctx *gin.Context)  // 查看前端文件的评论等级
	Download(ctx *gin.Context)     // 设置前端文件是否可以被下载
	DownloadShow(ctx *gin.Context) // 查看前端文件下载等级
	Can(ctx *gin.Context)          // 查看前端文件是否可以下载
	CommentCan(ctx *gin.Context)   // 查看前端文件是否可以评论
}

// FileVisibleController			定义了可见度工具类
type FileVisibleController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   设置当前文件可见度
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) Create(ctx *gin.Context) {
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

// @title    Show
// @description   查看前端文件可见度
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) Show(ctx *gin.Context) {

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

// @title    Comment
// @description   设置前端文件是否可以评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) Comment(ctx *gin.Context) {
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

// @title    CommentShow
// @description   查看评论的等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) CommentShow(ctx *gin.Context) {

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

// @title    Download
// @description   设置前端文件是否可以下载
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) Download(ctx *gin.Context) {
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

// @title    DownloadShow
// @description   查看前端文件下载等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) DownloadShow(ctx *gin.Context) {
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

// @title    CommentCan
// @description   查看前端文件是否可以评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) CommentCan(ctx *gin.Context) {
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

// @title    Can
// @description   查看前端文件是否可以下载
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (v FileVisibleController) Can(ctx *gin.Context) {
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

// @title    NewFileVisibleController
// @description   新建一个设置可见度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileVisibleController() IFileVisibleController {
	db := common.GetDB()
	return FileVisibleController{DB: db}
}
