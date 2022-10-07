// @Title  GuestbookController
// @Description  该文件用于提供留言板操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/vo"
	"ginEssential/common"
	"ginEssential/response"
	"log"
	"strconv"

	gmodel "ginEssential/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IGuestbookController			定义了留言板类接口
type IGuestbookController interface {
	Interface.RestInterface      //增删查改
	ShowOthers(ctx *gin.Context) // 查看他人的留言板
}

// GuestbookController			定义了留言板工具类
type GuestbookController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   创建留言
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (g GuestbookController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	userId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	// TODO 查看用户是否存在
	if g.DB.Where("id = ?", userId).First(&gmodel.User{}).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	var requestGuestbook vo.CreateGuestbookRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestGuestbook); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	guestbook := model.GuestBook{
		UserId:  uint(userId),
		Content: requestGuestbook.Content,
		Author:  user.ID,
	}

	g.DB.Create(&guestbook)

	response.Success(ctx, nil, "创建成功")
}

// @title    Show
// @description   查看自己的留言板
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (g GuestbookController) Show(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var guestbooks []model.GuestBook

	// TODO 查看对应数据
	g.DB.Where("user_id = ?", user.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&guestbooks)

	// TODO 查看总数
	var total int64
	g.DB.Where("user_id = ?", user.ID).Model(model.GuestBook{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"guestbooks": guestbooks, "total": total}, "查找成功")
}

// @title    ShowOthers
// @description   查看指定用户的留言板
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (g GuestbookController) ShowOthers(ctx *gin.Context) {

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	userId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var guestbooks []model.GuestBook

	// TODO 查看对应数据
	g.DB.Where("user_id = ?", userId).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&guestbooks)

	// TODO 查看总数
	var total int64
	g.DB.Where("user_id = ?", userId).Model(model.GuestBook{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"guestbooks": guestbooks, "total": total}, "查找成功")
}

// @title    Delete
// @description   删除指定的留言
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (g GuestbookController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	guestbookId := ctx.Params.ByName("id")

	var guestbook model.GuestBook

	// TODO 查看留言是否存在
	if g.DB.Where("id = ?", guestbookId).First(&guestbook).Error != nil {
		response.Fail(ctx, nil, "留言不存在")
		return
	}

	if guestbook.Author != user.ID && guestbook.UserId != user.ID {
		response.Fail(ctx, nil, "权限不足，不可删除")
		return
	}

	g.DB.Delete(&guestbook)

	response.Success(ctx, nil, "删除成功")
}

// @title    Update
// @description   更新指定的留言
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (g GuestbookController) Update(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	guestbookId := ctx.Params.ByName("id")

	var guestbook model.GuestBook

	// TODO 查看留言是否存在
	if g.DB.Where("id = ?", guestbookId).First(&guestbook).Error != nil {
		response.Fail(ctx, nil, "留言不存在")
		return
	}

	if guestbook.Author != user.ID {
		response.Fail(ctx, nil, "权限不足，不可更改")
		return
	}

	var requestGuestbook vo.CreateCommentRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestGuestbook); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 更新文章
	if err := g.DB.Model(&guestbook).Updates(requestGuestbook).Error; err != nil {
		response.Fail(ctx, nil, "更新失败")
		return
	}

	response.Success(ctx, nil, "更新成功")
}

// @title    NewGuestbookController
// @description   新建一个文件标签的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewGuestbookController() IGuestbookController {
	db := common.GetDB()
	db.AutoMigrate(model.GuestBook{})
	return GuestbookController{DB: db}
}
