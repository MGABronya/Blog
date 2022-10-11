// @Title  FileHistoryController
// @Description  该文件用于提供前端文件历史记录操作的各种函数
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

// IFileHistoryController			定义了历史记录类接口
type IFileHistoryController interface {
	Interface.HistoryInterface // 历史记录相关操作
}

// FileHistoryController			定义了历史记录工具类
type FileHistoryController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   用户创建前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHistoryController) Create(ctx *gin.Context) {
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

// @title    Show
// @description   用户查看前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHistoryController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var zipfileHistorys []model.ZipfileHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&zipfileHistorys)

	response.Success(ctx, gin.H{"zipfileHistorys": zipfileHistorys}, "查看前端文件历史记录成功")
}

// @title    Delete
// @description   用户删除前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHistoryController) Delete(ctx *gin.Context) {
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

// @title    DeleteAll
// @description   用户清空前端文件历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h FileHistoryController) DeleteAll(ctx *gin.Context) {
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
func NewFileHistoryController() IFileHistoryController {
	db := common.GetDB()
	db.AutoMigrate(model.ZipfileHistory{})
	return FileHistoryController{db}
}
