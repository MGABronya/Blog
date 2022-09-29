// @Title  FileFavoriteController
// @Description  该文件用于提供前端文件收藏操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	gmodel "ginEssential/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IFileFavoriteController			定义了前端文件收藏类接口
type IFileFavoriteController interface {
	Interface.FavoriteInterface // 收藏相关方法
}

// FileFavoriteController			定义了前端文件收藏工具类
type FileFavoriteController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建前端文件收藏
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileFavoriteController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	util.SetS(2, "aF"+fileId, strconv.Itoa(int(user.ID)))
	util.SetS(2, "Fa"+strconv.Itoa(int(user.ID)), fileId)

	response.Success(ctx, nil, "收藏成功")
}

// @title    Show
// @description   查看前端文件是否收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileFavoriteController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看文章是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(2, "aF"+fileId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消前端文件的收藏
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileFavoriteController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	util.RemS(2, "aF"+fileId, strconv.Itoa(int(user.ID)))
	util.RemS(2, "Fa"+strconv.Itoa(int(user.ID)), fileId)
	response.Success(ctx, nil, "删除成功")
}

// @title    FavoriteList
// @description   显示前端文件的收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileFavoriteController) FavoriteList(ctx *gin.Context) {

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	es := util.MembersS(2, "aF"+fileId)
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    UserFavoriteList
// @description   显示用户的前端文件收藏列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileFavoriteController) UserFavoriteList(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	es := util.MembersS(2, "Fa"+strconv.Itoa(int(user.ID)))
	total := len(es)
	response.Success(ctx, gin.H{"Favorites": es, "total": total}, "查看成功")
}

// @title    NewFileFavoriteController
// @description   新建一个前端文件收藏的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileFavoriteController() IFileFavoriteController {
	db := common.GetDB()
	return FileFavoriteController{DB: db}
}
