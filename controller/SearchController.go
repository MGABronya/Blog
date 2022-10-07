// @Title  SearchController
// @Description  该文件用于提供搜索操作的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	"Blog/model"
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

// ISearchController			定义了搜索类接口
type ISearchController interface {
	Article(ctx *gin.Context)               // 实现文章搜索
	Post(ctx *gin.Context)                  // 实现帖子搜索
	Zipfile(ctx *gin.Context)               // 实现前端文件搜索
	ArticleWithLabelInter(ctx *gin.Context) // 实现带标签交集文章搜索
	PostWithLabelInter(ctx *gin.Context)    // 实现带标签交集帖子搜索
	ZipfileWithLabelInter(ctx *gin.Context) // 实现带标签交集前端文件搜索
	ArticleWithLabelUnion(ctx *gin.Context) // 实现带标签并集文章搜索
	PostWithLabelUnion(ctx *gin.Context)    // 实现带标签并集帖子搜索
	ZipfileWithLabelUnion(ctx *gin.Context) // 实现带标签并集前端文件搜索
}

// SearchController			定义了搜索工具类
type SearchController struct {
	DB *gorm.DB
}

// @title    Article
// @description   文章文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) Article(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var articles []gmodel.Article

	// TODO 模糊匹配
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    Zipfile
// @description   前端文件文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) Zipfile(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)
	text := ctx.Params.ByName("text")

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("(visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", users, usera.ID, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    Post
// @description   帖子文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) Post(ctx *gin.Context) {
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

// @title    ArticleWithLabelInter
// @description   文章文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) ArticleWithLabelInter(ctx *gin.Context) {
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
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", articleIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", articleIds, users, usera.ID, text).Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    PostWithLabelInter
// @description   帖子文本带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) PostWithLabelInter(ctx *gin.Context) {
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

// @title    ZipfileWithLabelInter
// @description   前端文件带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) ZipfileWithLabelInter(ctx *gin.Context) {
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
	zipfileIds := util.InterS(2, requestLabels.Labels...)

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    ArticleWithLabelUnion
// @description   文章文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) ArticleWithLabelUnion(ctx *gin.Context) {
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
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", articleIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&articles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", articleIds, users, usera.ID, text).Model(gmodel.Article{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"articles": articles, "total": total}, "成功")
}

// @title    PostWithLabelUnion
// @description   帖子文本带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) PostWithLabelUnion(ctx *gin.Context) {
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

// @title    ZipfileWithLabelUnion
// @description   前端文件带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s SearchController) ZipfileWithLabelUnion(ctx *gin.Context) {
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
	zipfileIds := util.UnionS(2, requestLabels.Labels...)

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,comment,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    NewSearchController
// @description   新建一个搜索文本的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewSearchController() ISearchController {
	db := common.GetDB()
	return SearchController{db}
}
