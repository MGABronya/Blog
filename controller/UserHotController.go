// @Title  UserHotController
// @Description  该文件用于提供关于用户热度的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IUserHotController			定义了用户热度类接口
type IUserHotController interface {
	Interface.HotInterface         // 热度功能
	Level(ctx *gin.Context)        // 查看用户的热度等级
	MyLevel(ctx *gin.Context)      // 用户查看自己的热度等级
	PowerPoint(ctx *gin.Context)   // 查看用户的简报
	MyPowerPoint(ctx *gin.Context) // 用户查看自己的简报
}

// UserHotController			定义了热度工具类
type UserHotController struct {
	DB *gorm.DB //包含一个数据库指针
}

// 用户等级
var HotLevel []float64 = []float64{
	200, 400, 800, 1600, 3200, 6400, 12800, 25600,
}

// @title    Ranking
// @description   查看前端文件的热度排行
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) Ranking(ctx *gin.Context) {
	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.RangeZ(4, "H", int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)
	total := util.CardZ(4, "H")

	response.Success(ctx, gin.H{"users": users, "total": total}, "查看用户排行成功")
}

// @title    PowerPoint
// @description   查看用户简报
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) PowerPoint(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取其实时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	// TODO 获取path中的id
	userId := ctx.Params.ByName("id")

	// TODO 判断当前用户是否为所查用户
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '3' && userId != strconv.Itoa(int(user.ID)) {
		response.Fail(ctx, nil, "权限不足，无法查看")
		return
	}

	var powerpoints []model.PowerPoint

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", userId, start, end).Find(&powerpoints)

	response.Success(ctx, gin.H{"powerpoints": powerpoints}, "查看用户简报成功")
}

// @title    MyPowerPoint
// @description   用户查看自身简报
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) MyPowerPoint(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取起始时间和终止时间
	start := ctx.Query("start")
	end := ctx.Query("end")

	var powerpoints []model.PowerPoint

	h.DB.Where("user_id = ? and created_at >= ? and created_at <= ?", user.ID, start, end).Find(&powerpoints)

	response.Success(ctx, gin.H{"powerpoints": powerpoints}, "查看用户简报成功")
}

// @title    Recomment
// @description   查看用户推荐
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) Recomment(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	users := util.RangeZ(4, "R"+strconv.Itoa(int(user.ID)), int64(pageNum)*int64(pageSize), int64(pageNum)*int64(pageSize)+int64(pageSize)-1)

	total := util.CardZ(4, "R"+strconv.Itoa(int(user.ID)))

	response.Success(ctx, gin.H{"users": users, "total": total}, "查看用户推荐成功")
}

// @title    Level
// @description   查看用户等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) Level(ctx *gin.Context) {

	// TODO 获取path中的id
	userId := ctx.Params.ByName("id")

	Level := 0

	score := util.ScoreZ(4, "H", userId)

	for i := 7; i >= 0; i-- {
		if score >= HotLevel[i] {
			Level = i + 1
			break
		}
	}

	response.Success(ctx, gin.H{"Level": Level}, "查看用户热度等级成功")
}

// @title    MyLevel
// @description   用户查看自身等级
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (h UserHotController) MyLevel(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	Level := 0

	score := util.ScoreZ(4, "H", strconv.Itoa(int(user.ID)))

	for i := 7; i >= 0; i-- {
		if score >= HotLevel[i] {
			Level = i + 1
			break
		}
	}

	response.Success(ctx, gin.H{"Level": Level}, "查看用户热度等级成功")
}

// @title    NewUserHotController
// @description   新建一个热度的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewUserHotController() IUserHotController {
	db := common.GetDB()
	db.AutoMigrate(model.PowerPoint{})
	return UserHotController{DB: db}
}
