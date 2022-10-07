// @Title  FileLableController
// @Description  该文件用于提供前端文件标签操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"Blog/vo"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	gmodel "ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IFileLabelController			定义了文件标签类接口
type IFileLabelController interface {
	Interface.LabelInterface //创查删
}

// FileLabelController			定义了文件标签工具类
type FileLabelController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建文件标签
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLabelController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	// TODO 判断当前用户是否为文件的作者
	if user.ID != file.UserId {
		response.Fail(ctx, nil, "文件不属于您，请勿非法操作")
		return
	}

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	if util.IsS(2, "aL"+fileId, requestLabel.Label) {
		response.Fail(ctx, nil, "标签已设置")
		return
	}

	// TODO 用户标签分数上升
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label, 30)

	util.SetS(2, "aL"+fileId, requestLabel.Label)
	util.SetS(2, "La"+requestLabel.Label, fileId)

	response.Success(ctx, nil, "设置成功")
}

// @title    Show
// @description   查看文件的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLabelController) Show(ctx *gin.Context) {

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	response.Success(ctx, gin.H{"labels": util.MembersS(2, "aL"+fileId)}, "查找成功")
}

// @title    Delete
// @description   删除文件的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileLabelController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	// TODO 查看文件是否存在
	if f.DB.Where("id = ?", fileId).First(&file).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	// TODO 判断当前用户是否为文件的作者
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != file.UserId {
		response.Fail(ctx, nil, "文件不属于您，请勿非法操作")
		return
	}

	if !util.IsS(2, "aL"+fileId, requestLabel.Label) {
		response.Fail(ctx, nil, "标签未设置")
		return
	}

	// TODO 用户标签分数下降
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label, -30)
	if util.ScoreZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label) <= 0 {
		util.RemZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label)
	}

	util.RemS(2, "aL"+fileId, requestLabel.Label)
	util.RemS(2, "La"+requestLabel.Label, fileId)

	if util.CardS(2, "La"+requestLabel.Label) == 0 {
		util.Del(2, "La"+requestLabel.Label)
	}

	response.Success(ctx, nil, "删除成功")
}

// @title    NewFileLabelController
// @description   新建一个文件标签的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileLabelController() IFileLabelController {
	db := common.GetDB()
	return FileLabelController{DB: db}
}
