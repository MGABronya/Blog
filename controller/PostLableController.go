// @Title  PostLableController
// @Description  该文件用于提供帖子标签操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"Blog/vo"
	"ginEssential/common"
	"ginEssential/response"
	"strconv"

	"ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IPostLabelController			定义了帖子标签类接口
type IPostLabelController interface {
	Interface.LabelInterface //创查删
}

// PostLabelController			定义了帖子标签工具类
type PostLabelController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建帖子标签
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLabelController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看文章是否存在
	if p.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 判断当前用户是否为帖子的作者
	if user.ID != post.UserId {
		response.Fail(ctx, nil, "帖子不属于您，请勿非法操作")
		return
	}

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	if util.IsS(3, "aL"+postId, requestLabel.Label) {
		response.Fail(ctx, nil, "标签已设置")
		return
	}

	// TODO 用户标签分数上升
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), requestLabel.Label, 30)

	util.SetS(3, "aL"+postId, requestLabel.Label)
	util.SetS(3, "La"+requestLabel.Label, postId)

	response.Success(ctx, nil, "设置成功")
}

// @title    Show
// @description   查看帖子的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLabelController) Show(ctx *gin.Context) {

	postId := ctx.Params.ByName("id")

	var post model.Post

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"labels": util.MembersS(3, "aL"+postId)}, "查找成功")
}

// @title    Delete
// @description   删除帖子的标签
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (p PostLabelController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(model.User)

	postId := ctx.Params.ByName("id")

	var post model.Post

	label := ctx.Query("label")

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 判断当前用户是否为文章的作者
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != post.UserId {
		response.Fail(ctx, nil, "帖子不属于您，请勿非法操作")
		return
	}

	if !util.IsS(3, "aL"+postId, label) {
		response.Fail(ctx, nil, "标签未设置")
		return
	}

	// TODO 用户标签分数下降
	util.IncrByZ(4, "L"+strconv.Itoa(int(user.ID)), label, -30)
	if util.ScoreZ(4, "L"+strconv.Itoa(int(user.ID)), label) <= 0 {
		util.RemZ(4, "L"+strconv.Itoa(int(user.ID)), label)
	}

	util.RemS(3, "aL"+postId, label)
	util.RemS(3, "La"+label, postId)

	if util.CardS(3, "La"+label) == 0 {
		util.Del(3, "La"+label)
	}

	response.Success(ctx, nil, "删除成功")
}

// @title    NewPostLabelController
// @description   新建一个文章标签的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostLabelController() IPostLabelController {
	db := common.GetDB()
	return PostLabelController{DB: db}
}
