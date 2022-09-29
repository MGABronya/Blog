// @Title FileLikeController
// @Description  该文件用于提供文前端文件点赞操作的各种函数
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

	"Blog/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IFileLikeController			定义了前端文件点赞类接口
type IFileLikeController interface {
	Interface.LikeInterface //点赞相关方法
}

// FileLikeController			定义了前端文件点赞工具类
type FileLikeController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建前端文件点赞
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLikeController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	util.SetS(2, "fiL"+fileId, strconv.Itoa(int(user.ID)))

	response.Success(ctx, nil, "点赞成功")
}

// @title    Show
// @description   查看前端文件是否点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLikeController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(2, "fiL"+fileId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消前端文件的点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLikeController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	util.RemS(2, "fiL"+fileId, strconv.Itoa(int(user.ID)))
	response.Success(ctx, nil, "删除成功")
}

// @title    LikeList
// @description   显示前端文件的点赞列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLikeController) LikeList(ctx *gin.Context) {

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看跟帖是否存在
	if f.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	es := util.MembersS(2, "fiL"+fileId)
	total := len(es)
	response.Success(ctx, gin.H{"Liks": es, "total": total}, "查看成功")
}

// @title    NewFileLikeController
// @description   新建一个前端文件点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileLikeController() IFileLikeController {
	db := common.GetDB()
	return FileLikeController{DB: db}
}
