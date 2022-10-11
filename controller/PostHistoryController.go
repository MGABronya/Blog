// @Title  PostHistoryController
// @Description  该文件用于提供帖子历史记录操作的各种函数
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

// IPostHistoryController			定义了历史记录类接口
type IPostHistoryController interface {
	Interface.HistoryInterface // 关于历史记录的相关功能
}

// PostHistoryController			定义了历史记录工具类
type PostHistoryController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @description   用户创建帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHistoryController) Create(ctx *gin.Context) {
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

// @title    Show
// @description   用户查看帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHistoryController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var postHistorys []model.PostHistory

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&postHistorys)

	response.Success(ctx, gin.H{"postHistorys": postHistorys}, "查看帖子历史记录成功")
}

// @title    Delete
// @description   用户删除帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHistoryController) Delete(ctx *gin.Context) {
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

// @title    DeleteAll
// @description   用户清空帖子历史记录
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHistoryController) DeleteAll(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	h.DB.Where("user_id = ?", user.ID).Delete(model.PostHistory{})

	response.Success(ctx, nil, "清空帖子历史记录成功")
}

// @title    NewHistoryController
// @description   新建一个历史记录的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostHistoryController() IPostHistoryController {
	db := common.GetDB()
	db.AutoMigrate(model.PostHistory{})
	return PostHistoryController{db}
}
