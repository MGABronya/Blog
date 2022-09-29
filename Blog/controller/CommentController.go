// @Title  CommentController
// @Description  该文件用于提供操作前端文件评论的各种方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"Blog/vo"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ICommentController			定义了评论类接口
type ICommentController interface {
	Interface.RestInterface          // 实现增删查改功能
	PageList(ctx *gin.Context)       // 实现返回评论列表
	PageListMine(ctx *gin.Context)   // 实现返回用户个人的评论列表
	PageListOthers(ctx *gin.Context) // 实现返回某一用户的评论列表
}

// CommentController			定义了评论工具类
type CommentController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建一篇评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) Create(ctx *gin.Context) {
	var requestComment vo.CreateCommentRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestComment); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if c.DB.Where("id = ?", ctx.Params.ByName("id")).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 获取登录用户
	user, _ := ctx.Get("user")

	if !util.ZipfileComment(strconv.Itoa(int(user.(gmodel.User).ID)), file.ID.String(), strconv.Itoa(int(file.UserId))) {
		response.Fail(ctx, nil, "权限不足，不可评论")
		return
	}

	// TODO 创建Thread
	comment := model.Comment{
		UserId:   user.(gmodel.User).ID,
		FileId:   ctx.Params.ByName("id"),
		Content:  requestComment.Content,
		ResLong:  requestComment.ResLong,
		ResShort: requestComment.ResShort,
	}

	// TODO 插入数据
	if err := c.DB.Create(&comment).Error; err != nil {
		panic(err)
	}

	// TODO 成功
	response.Success(ctx, nil, "创建成功")
}

// @title    Update
// @description   更新评论内容
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) Update(ctx *gin.Context) {
	var requestComment vo.CreateCommentRequest
	// TODO 数据验证
	if err := ctx.ShouldBind(&requestComment); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(gmodel.User).ID

	// TODO 获取path中的id
	fileId := ctx.Params.ByName("id")

	var file model.ZipFile
	if c.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 判断当前用户是否为跟帖的作者
	if userId != file.UserId {
		response.Fail(ctx, nil, "前端文件不属于您，请勿非法操作")
		return
	}

	var comment model.Comment
	// TODO 更新帖子
	if err := c.DB.Model(&comment).Update(requestComment).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, gin.H{"comment": comment}, "更新成功")
}

// @title    Show
// @description   查看评论内容
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) Show(ctx *gin.Context) {
	// TODO 获取path中的id
	commentId := ctx.Params.ByName("id")

	var comment model.Comment
	// TODO 查看评论是否存在
	if c.DB.Where("id = ?", commentId).First(&comment).RecordNotFound() {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	var user gmodel.User
	c.DB.Where("id = ?", comment.UserId).First(&user)

	// TODO 返回数据
	response.Success(ctx, gin.H{"comment": comment}, "成功")
}

// @title    Delete
// @description   删除评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) Delete(ctx *gin.Context) {

	// TODO 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(gmodel.User).ID

	// TODO 获取path中的id
	commentId := ctx.Params.ByName("id")

	var comment model.Comment
	if c.DB.Where("id = ?", commentId).First(&comment).RecordNotFound() {
		response.Fail(ctx, nil, "评论不存在")
		return
	}

	if util.GetH(0, "permission", strconv.Itoa(int(userId)))[0] < '4' && userId != comment.UserId {
		response.Fail(ctx, nil, "评论不属于您，请勿非法操作")
		return
	}

	// TODO 删除评论
	c.DB.Delete(&comment)

	// TODO 移除点赞
	util.Del(2, "ciL"+comment.ID.String())

	response.Success(ctx, nil, "删除成功")
}

// @title    PageList
// @description   返回评论的列表
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) PageList(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// TODO 获取path中的id
	fileId := ctx.Params.ByName("id")

	var file model.ZipFile

	// TODO 查看前端文件是否存在
	if c.DB.Where("id = ?", fileId).First(&file).RecordNotFound() {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 获取登录用户
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 查看是否有权限
	if file.UserId != user.ID && (file.Visible == 3 || (file.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(file.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	// TODO 分页
	var comments []model.Comment
	c.DB.Where("file_id = ?", fileId).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&comments)

	// TODO 记录的总条数
	var total int
	c.DB.Model(model.Comment{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"comments": comments, "total": total}, "成功")
}

// @title    PageListMine
// @description   返回用户个人评论的列表
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) PageListMine(ctx *gin.Context) {
	// TODO 获取登录用户
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// TODO 分页
	var comments []model.Comment
	c.DB.Where("user_id = ?", user.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&comments)

	// TODO 记录的总条数
	var total int
	c.DB.Where("user_id = ?", user.ID).Model(model.Comment{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"comments": comments, "total": total}, "成功")
}

// @title    PageListOthers
// @description   返回某一用户评论的列表
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (c CommentController) PageListOthers(ctx *gin.Context) {
	// TODO 获取登录用户
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	userId := ctx.Params.ByName("id")

	// TODO 查看用户是否存在
	if c.DB.Where("id = ?", ctx.Params.ByName("id")).First(&gmodel.User{}).RecordNotFound() {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 分页
	var comments []model.Comment
	var level int8
	if strconv.Itoa(int(user.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	// TODO 查找所有分页中可见的条目
	c.DB.Table("comments").Joins("join zip_files on comments.file_id = zip_files.id").Where("zip_files.user_id = ? and zip_files.visible < ?", user.ID, level).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&comments)

	// TODO 记录的总条数
	var total int
	c.DB.Table("comments").Joins("join zip_files on comments.file_id = zip_files.id").Where("zip_files.user_id = ? and zip_files.visible < ?", user.ID, level).Model(model.Comment{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"comments": comments, "total": total}, "成功")
}

// @title    NewCommentController
// @description   新建一个评论的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewCommentController() ICommentController {
	db := common.GetDB()
	db.AutoMigrate(model.Comment{})
	return CommentController{DB: db}
}
