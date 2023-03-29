// @Title  PostHotController
// @Description  该文件用于提供关于帖子热度的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IPostHotController			定义了帖子热度类接口
type IPostHotController interface {
	Interface.HotInterface    //热度相关方法
	Visit(ctx *gin.Context)   // 帖子的游览次数获取
	Thread(ctx *gin.Context)  // 帖子的跟帖人数
	Ranking(ctx *gin.Context) // 帖子的热度排行
}

// PostHotController			定义了帖子热度工具类
type PostHotController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Visit
// @description   用户获取帖子的游览次数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHotController) Visit(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否有权限
	if post.UserId != user.ID && (post.Visible == 3 || (post.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(post.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"views": util.CardS(3, "C"+postId)}, "查看游览次数成功")
}

// @title    Thread
// @description   用户获取帖子的跟帖人数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHotController) Thread(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	postId := ctx.Params.ByName("id")

	var post gmodel.Post

	// TODO 查看帖子是否在数据库中存在
	if h.DB.Where("id = ?", postId).First(&post).Error != nil {
		response.Fail(ctx, nil, "帖子不存在")
		return
	}

	// TODO 查看是否有权限
	if post.UserId != user.ID && (post.Visible == 3 || (post.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(post.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"threads": util.CardS(4, "M"+postId)}, "查看跟帖人数成功")
}

// @title    Ranking
// @description   查看帖子的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHotController) Ranking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	posts := util.RangeZ(3, "H", int64(pageNum-1)*int64(pageSize), int64(pageNum-1)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(3, "H")

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "查看帖子排行成功")
}

// @title    Recomment
// @description   查看帖子推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h PostHotController) Recomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	posts := util.RangeZ(3, "R"+strconv.Itoa(int(user.ID)), int64(pageNum-1)*int64(pageSize), int64(pageNum-1)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(3, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"posts": posts, "total": total}, "查看帖子推荐成功")
}

// @title    NewPostHotController
// @description   新建一个帖子热度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostHotController() IPostHotController {
	db := common.GetDB()
	return PostHotController{DB: db}
}
