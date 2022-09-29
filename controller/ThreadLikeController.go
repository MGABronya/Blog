// @Title ThreadLikeController
// @Description  该文件用于提供文跟帖点赞操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// IThreadLikeController			定义了跟帖点赞类接口
type IThreadLikeController interface {
	Interface.LikeInterface //点赞相关方法
}

// ThreadLikeController			定义了跟帖点赞工具类
type ThreadLikeController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建跟帖点赞
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadLikeController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).RecordNotFound() {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	util.SetS(3, "tiL"+threadId, strconv.Itoa(int(user.ID)))

	response.Success(ctx, nil, "点赞成功")
}

// @title    Show
// @description   查看跟帖是否点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadLikeController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(3, "tiL"+threadId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消跟帖的点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadLikeController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).RecordNotFound() {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	util.RemS(3, "tiL"+threadId, strconv.Itoa(int(user.ID)))
	response.Success(ctx, nil, "删除成功")
}

// @title    LikeList
// @description   显示跟帖的点赞列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (t ThreadLikeController) LikeList(ctx *gin.Context) {

	threadId := ctx.Params.ByName("id")

	var thread model.Thread

	// TODO 查看跟帖是否存在
	if t.DB.Where("id = ?", threadId).First(&thread).RecordNotFound() {
		response.Fail(ctx, nil, "跟帖不存在")
		return
	}

	es := util.MembersS(3, "tiL"+threadId)
	total := len(es)
	response.Success(ctx, gin.H{"Liks": es, "total": total}, "查看成功")
}

// @title    NewThreadLikeController
// @description   新建一个跟帖点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewThreadLikeController() IThreadLikeController {
	db := common.GetDB()
	return ThreadLikeController{DB: db}
}
