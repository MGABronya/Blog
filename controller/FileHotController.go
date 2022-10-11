// @Title  FileHotController
// @Description  该文件用于提供前端文件关于热度的各种函数
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

// IFileHotController			定义了前端文件热度类接口
type IFileHotController interface {
	Interface.HotInterface     // 热度相关功能
	Visit(ctx *gin.Context)    // 前端文件的游览次数获取
	Download(ctx *gin.Context) // 前端文件的下载次数获取
	Use(ctx *gin.Context)      // 前端文件的使用人数获取
	Comment(ctx *gin.Context)  // 前端文件的评论人数
}

// FileHotController			定义了热度工具类
type FileHotController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Visit
// @description   用户获取前端文件的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Visit(ctx *gin.Context) {
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

// @title    Download
// @description   用户获取前端文件的下载次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Download(ctx *gin.Context) {
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

// @title    Use
// @description   用户获取前端文件的使用次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Use(ctx *gin.Context) {
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

// @title    Comment
// @description   用户获取前端文件的评论人数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Comment(ctx *gin.Context) {
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

// @title    Ranking
// @description   查看前端文件的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Ranking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	zipfiles := util.RangeZ(2, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(2, "H")

	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "查看前端文件排行成功")
}

// @title    Recomment
// @description   查看前端文件推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHotController) Recomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	zipfiles := util.RangeZ(2, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(2, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "查看前端文件推荐成功")
}

// @title    NewFileHotController
// @description   新建一个热度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileHotController() IFileHotController {
	db := common.GetDB()
	return FileHotController{DB: db}
}
