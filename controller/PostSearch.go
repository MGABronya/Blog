// @Title  PostSearchController
// @Description  该文件用于提供帖子搜索操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/util"
	"Blog/vo"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IPostSearchController			定义了帖子搜索类接口
type IPostSearchController interface {
	Interface.SeartchInterface // 定义了搜索类方法
}

// PostSearchController			定义了帖子搜索工具类
type PostSearchController struct {
	DB *gorm.DB
}

// @title    Show
// @description   帖子文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var posts []gmodel.Post
	var total int64

	// TODO 模糊匹配
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    ShowUser
// @description   用户指定的帖子文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) ShowUser(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")
	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var posts []gmodel.Post

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    ShowWithLabelInter
// @description   帖子文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) ShowWithLabelInter(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var requestLabels vo.LabelsRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestLabels); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 增加前缀
	for i := range requestLabels.Labels {
		requestLabels.Labels[i] = "La" + requestLabels.Labels[i]
	}

	// TODO 求出标签们的交集
	postIds := util.InterS(3, requestLabels.Labels...)

	var posts []gmodel.Post

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", postIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", postIds, users, usera.ID, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    ShowWithLabelInterUser
// @description   指定用户的帖子文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) ShowWithLabelInterUser(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var requestLabels vo.LabelsRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestLabels); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 增加前缀
	for i := range requestLabels.Labels {
		requestLabels.Labels[i] = "La" + requestLabels.Labels[i]
	}

	// TODO 求出标签们的交集
	postIds := util.InterS(3, requestLabels.Labels...)

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	var posts []gmodel.Post

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, postIds, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, postIds, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    ShowWithLabelUnion
// @description   帖子文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) ShowWithLabelUnion(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var requestLabels vo.LabelsRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestLabels); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 增加前缀
	for i := range requestLabels.Labels {
		requestLabels.Labels[i] = "La" + requestLabels.Labels[i]
	}

	// TODO 求出标签们的并集
	postIds := util.UnionS(3, requestLabels.Labels...)

	var posts []gmodel.Post

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", postIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", postIds, users, usera.ID, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    ShowWithLabelUnionUser
// @description   指定用户的帖子文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s PostSearchController) ShowWithLabelUnionUser(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var requestLabels vo.LabelsRequest

	// TODO 数据验证
	if err := ctx.ShouldBind(&requestLabels); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// TODO 增加前缀
	for i := range requestLabels.Labels {
		requestLabels.Labels[i] = "La" + requestLabels.Labels[i]
	}

	// TODO 求出标签们的并集
	postIds := util.UnionS(3, requestLabels.Labels...)

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	var posts []gmodel.Post

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, postIds, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, postIds, text).Model(gmodel.Post{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"posts": posts, "total": total}, "成功")
}

// @title    NewPostSearchController
// @description   新建一个搜索帖子文本的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewPostSearchController() IPostSearchController {
	db := common.GetDB()
	return PostSearchController{db}
}
