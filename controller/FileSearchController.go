// @Title  FileSearchController
// @Description  该文件用于提供前端文件搜索操作的各种函数
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
	"gorm.io/gorm"
)

// IFileSearchController			定义了搜索类接口
type IFileSearchController interface {
	Interface.SeartchInterface // 定义了搜索类相关方法
}

// FileSearchController			定义了搜索工具类
type FileSearchController struct {
	DB *gorm.DB
}

// @title    Show
// @description   前端文件文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) Show(ctx *gin.Context) {
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

// @title    ShowUser
// @description   指定用户的前端文件文本搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) ShowUser(ctx *gin.Context) {
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

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    ShowWithLabelInter
// @description   前端文件带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) ShowWithLabelInter(ctx *gin.Context) {
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
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("id in (?) and (visible = 2 and user_id in (?) or visible = 1 or visible = 3 and user_id = ?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", zipfileIds, users, usera.ID, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    ShowWithLabelInterUser
// @description   指定用户的前端文件带标签交集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) ShowWithLabelInterUser(ctx *gin.Context) {
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
	zipfileIds := util.InterS(2, requestLabels.Labels...)

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, zipfileIds, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, zipfileIds, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    ShowWithLabelUnion
// @description   前端文件带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) ShowWithLabelUnion(ctx *gin.Context) {
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

// @title    ShowWithLabelUnionUser
// @description   指定用户的前端文件带标签并集搜索
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (s FileSearchController) ShowWithLabelUnionUser(ctx *gin.Context) {
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
	zipfileIds := util.UnionS(2, requestLabels.Labels...)

	userId := ctx.Params.ByName("id")

	var level int8
	if strconv.Itoa(int(usera.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(usera.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	var zipfiles []model.ZipFile

	// TODO 模糊匹配
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, zipfileIds, text).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 查看查询总数
	var total int64
	s.DB.Where("visible < ? and user_id = ? and id in (?) and match(title,content,res_long,res_short) against('?*' in boolean mode)", level, userId, zipfileIds, text).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"zipfiles": zipfiles, "total": total}, "成功")
}

// @title    NewFileSearchController
// @description   新建一个搜索文本的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileSearchController() IFileSearchController {
	db := common.GetDB()
	return FileSearchController{db}
}
