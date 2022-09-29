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
	"github.com/jinzhu/gorm"
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
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
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

	util.SetS(3, "paL"+postId, requestLabel.Label)
	util.SetS(3, "pLa"+requestLabel.Label, postId)

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
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	response.Success(ctx, gin.H{"labels": util.GetS(3, "paL"+postId)}, "查找成功")
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

	var requestLabel = vo.LabelRequest{}
	ctx.Bind(&requestLabel)

	// TODO 查看帖子是否存在
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 判断当前用户是否为文章的作者
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != post.UserId {
		response.Fail(ctx, nil, "帖子不属于您，请勿非法操作")
		return
	}

	util.RemS(3, "paL"+postId, requestLabel.Label)
	util.RemS(3, "pLa"+requestLabel.Label, postId)

	if util.CardS(3, "pLa"+requestLabel.Label) == 0 {
		util.Del(3, "pLa"+requestLabel.Label)
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
