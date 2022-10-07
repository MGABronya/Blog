// @Title CommentLikeController
// @Description  该文件用于提供评论点赞操作的各种函数
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
	"gorm.io/gorm"
)

// ICommentLikeController			定义了评论点赞类接口
type ICommentLikeController interface {
	Interface.LikeInterface //点赞相关方法
}

// CommentLikeController			定义了评论点赞工具类
type CommentLikeController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建评论点赞
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentLikeController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	commentId := ctx.Params.ByName("id")

	var comment model.Comment

	// TODO 查看评论是否存在
	if c.DB.Where("id = ?", commentId).First(&comment).Error != nil {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	// TODO 查看前端文件是否存在
	var file model.ZipFile
	if c.DB.Where("id = ?", comment.FileId).First(&file).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看评论是否已经点赞
	if util.IsS(2, "ciL"+commentId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "评论已点赞")
		return
	}

	util.SetS(2, "ciL"+commentId, strconv.Itoa(int(user.ID)))

	util.IncrByZ(2, "H", comment.FileId, 5)
	util.IncrByZ(2, "CH", commentId, 10)
	util.IncrByZ(4, "H", strconv.Itoa(int(comment.UserId)), 10)
	util.IncrByZ(4, "H", strconv.Itoa(int(file.UserId)), 5)

	response.Success(ctx, nil, "点赞成功")
}

// @title    Show
// @description   查看评论是否点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentLikeController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	commentId := ctx.Params.ByName("id")

	var comment model.Comment

	// TODO 查看评论是否存在
	if c.DB.Where("id = ?", commentId).First(&comment).Error != nil {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	response.Success(ctx, gin.H{"flag": util.IsS(2, "ciL"+commentId, strconv.Itoa(int(user.ID)))}, "查看成功")
}

// @title    Delete
// @description   取消评论的点赞
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentLikeController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	commentId := ctx.Params.ByName("id")

	var comment model.Comment

	// TODO 查看评论是否存在
	if c.DB.Where("id = ?", commentId).First(&comment).Error != nil {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	// TODO 查看前端文件是否存在
	var file model.ZipFile
	if c.DB.Where("id = ?", comment.FileId).First(&file).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看评论是否已经点赞
	if !util.IsS(2, "ciL"+commentId, strconv.Itoa(int(user.ID))) {
		response.Fail(ctx, nil, "评论未点赞")
		return
	}

	util.RemS(2, "ciL"+commentId, strconv.Itoa(int(user.ID)))

	util.IncrByZ(2, "H", comment.FileId, -5)
	util.IncrByZ(2, "CH", comment.ID.String(), -10)
	util.IncrByZ(4, "H", strconv.Itoa(int(comment.UserId)), -10)
	util.IncrByZ(4, "H", strconv.Itoa(int(file.UserId)), -5)

	response.Success(ctx, nil, "删除成功")
}

// @title    LikeList
// @description   显示评论的点赞列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentLikeController) LikeList(ctx *gin.Context) {

	commentId := ctx.Params.ByName("id")

	var comment model.Comment

	// TODO 查看评论是否存在
	if c.DB.Where("id = ?", commentId).First(&comment).Error != nil {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	es := util.MembersS(2, "ciL"+commentId)
	total := len(es)
	response.Success(ctx, gin.H{"Liks": es, "total": total}, "查看成功")
}

// @title    NewCommentLikeController
// @description   新建一个评论点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewCommentLikeController() ICommentLikeController {
	db := common.GetDB()
	return CommentLikeController{DB: db}
}
