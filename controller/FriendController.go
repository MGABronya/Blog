// @Title FriendController
// @Description  该文件用于提供好友操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"Blog/model"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IFriendController			定义了好友类接口
type IFriendController interface {
	Show(ctx *gin.Context)         // 查看好友列表
	ShowApplied(ctx *gin.Context)  // 查看好友申请列表
	ShowApplying(ctx *gin.Context) // 查看申请列表
	Applying(ctx *gin.Context)     // 发送好友申请
	Applied(ctx *gin.Context)      // 接受好友申请
	Delete(ctx *gin.Context)       // 删除好友
	ShowArticles(ctx *gin.Context) // 查看好友圈内文章
	ShowPosts(ctx *gin.Context)    // 查看好友圈内贴子
	ShowZipfiles(ctx *gin.Context) // 查看好友圈内前端文件
}

// FriendController			定义了好友工具类
type FriendController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Show
// @description   查看好友列表
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	friends := util.MembersS(4, "Fr"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"friends": friends}, "查看成功")
}

// @title    ShowApplied
// @description   查看查看好友申请列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) ShowApplied(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	friends := util.MembersS(4, "Frad"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"friends": friends}, "查看成功")
}

// @title    ShowApplying
// @description   查看申请列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) ShowApplying(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	friends := util.MembersS(4, "Frag"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"friends": friends}, "查看成功")
}

// @title    Applying
// @description   发送好友申请
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) Applying(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	userId := ctx.Params.ByName("id")

	var userb gmodel.User

	// TODO 查看用户是否存在
	if f.DB.Where("id = ?", ctx.Params.ByName("id")).First(&userb).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 查看用户是否已经为好友
	if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		response.Fail(ctx, nil, "已经是好友了")
	}

	// TODO 加入申请列表
	util.SetS(4, "Frag"+strconv.Itoa(int(usera.ID)), userId)

	// TODO 加入被申请列表
	util.SetS(4, "Frad"+userId, strconv.Itoa(int(usera.ID)))

	response.Success(ctx, nil, "申请成功")
}

// @title    Applied
// @description   接受好友申请
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) Applied(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	userId := ctx.Params.ByName("id")

	var userb gmodel.User

	// TODO 查看用户是否存在
	if f.DB.Where("id = ?", ctx.Params.ByName("id")).First(&userb).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 查看用户是否已经为好友
	if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		response.Fail(ctx, nil, "已经是好友了")
	}

	// TODO 删除原先在被申请列表中的位置
	util.RemS(4, "Frad"+strconv.Itoa(int(usera.ID)), userId)

	// TODO 删除原先在申请列表里的位置
	util.RemS(4, "Frag"+userId, strconv.Itoa(int(usera.ID)))

	// TODO 加入a的好友列表
	util.SetS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId)

	// TODO 加入b的好友列表
	util.SetS(4, "Fr"+userId, strconv.Itoa(int(usera.ID)))

	response.Success(ctx, nil, "接受成功")
}

// @title    Delete
// @description   删除好友
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) Delete(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	userId := ctx.Params.ByName("id")

	var userb gmodel.User

	// TODO 查看用户是否存在
	if f.DB.Where("id = ?", ctx.Params.ByName("id")).First(&userb).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 查看用户是否已经为好友
	if !util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		response.Fail(ctx, nil, "已经不是好友了")
	}

	// TODO 删除a好友列表中的位置
	util.RemS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId)

	// TODO 删除b好友列表里的位置
	util.RemS(4, "Fr"+userId, strconv.Itoa(int(usera.ID)))

	response.Success(ctx, nil, "删除成功")
}

// @title    ShowArticles
// @description   展示朋友圈文章
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) ShowArticles(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	var articles []gmodel.Article

	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 记录的总条数
	var total int64
	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Model(gmodel.Article{}).Count(&total)

	response.Success(ctx, gin.H{"articles": articles, "total": total}, "查找成功")
}

// @title    ShowPosts
// @description   展示朋友圈帖子
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) ShowPosts(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	var posts []gmodel.Post

	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 记录的总条数
	var total int64
	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Model(gmodel.Post{}).Count(&total)

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "查找成功")
}

// @title    ShowZipfiles
// @description   展示朋友圈前端文件
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FriendController) ShowZipfiles(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	var zipfiles []model.ZipFile

	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 记录的总条数
	var total int64
	f.DB.Where("user_id in ? and visible < 3", users).Or("user_id = ?", usera.ID).Model(model.ZipFile{}).Count(&total)

	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "查找成功")
}

// @title    NewFriendController
// @description   新建一个前端文件点赞的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFriendController() IFriendController {
	db := common.GetDB()
	return FriendController{DB: db}
}
