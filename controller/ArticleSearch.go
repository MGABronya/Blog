// @Title  ArticleSearchController
// @Description  该文件用于提供搜索操作的各种函数
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

// IArticleSearchController			定义了文章搜索类接口
type IArticleSearchController interface {
	Interface.SeartchInterface // 定义了搜索类方法
}

// ArticleSearchController			定义了搜索工具类
type ArticleSearchController struct {
	DB *gorm.DB
}

// @title    Show
// @description   文章文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) Show(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", users, usera.ID, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", users, usera.ID, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    ShowUser
// @description   指定用户的文章文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) ShowUser(ctx *gin.Context) {
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

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    ShowWithLabelInter
// @description   文章文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) ShowWithLabelInter(ctx *gin.Context) {
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
	articleIds := util.InterS(1, requestLabels.Labels...)

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", articleIds, users, usera.ID, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", articleIds, users, usera.ID, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    ShowWithLabelInterUser
// @description   文章文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) ShowWithLabelInterUser(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

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
	articleIds := util.InterS(1, requestLabels.Labels...)

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, articleIds, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, articleIds, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    ShowWithLabelUnion
// @description   文章文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) ShowWithLabelUnion(ctx *gin.Context) {
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
	articleIds := util.UnionS(1, requestLabels.Labels...)

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", articleIds, users, usera.ID, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against(? in boolean mode)", articleIds, users, usera.ID, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    ShowWithLabelUnionUser
// @description   指定用户的文章文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s ArticleSearchController) ShowWithLabelUnionUser(ctx *gin.Context) {
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
	articleIds := util.UnionS(1, requestLabels.Labels...)

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, articleIds, text+"*").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against(? in boolean mode)", level, userId, articleIds, text+"*").Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    NewArticleSearchController
// @description   新建一个搜索文本的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewArticleSearchController() IArticleSearchController {
	db := common.GetDB()
	return ArticleSearchController{db}
}
