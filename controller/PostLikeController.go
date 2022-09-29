// @Title PostLikeController
// @Description  该文件用于提供文帖子点赞操作的各种函数
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

// IPostLikeController			定义了帖子点赞类接口
type IPostLikeController interface {
	Interface.LikeInterface //点赞相关方法
}

// PostLikeController			定义了帖子点赞工具类
type PostLikeController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建帖子点赞
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLikeController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	util.SetS(3, "piL"+postId, strconv.Itoa(int(user.ID)))

	response.Success(ctx, nil, "点赞成功")
}

// @title    Show
// @description   查看帖子是否点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLikeController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(3, "piL"+postId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消帖子的点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLikeController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	util.RemS(3, "piL"+postId, strconv.Itoa(int(user.ID)))
	response.Success(ctx, nil, "删除成功")
}

// @title    LikeList
// @description   显示帖子的点赞列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLikeController) LikeList(ctx *gin.Context) {

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	es := util.MembersS(3, "piL"+postId)
	total := len(es)
	response.Success(ctx, gin.H{"Liks": es, "total": total}, "查看成功")
}

// @title    NewPostLikeController
// @description   新建一个帖子点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostLikeController() IPostLikeController {
	db := common.GetDB()
	return PostLikeController{DB: db}
}
