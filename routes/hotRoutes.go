// @Title  hotRoutes
// @Description  程序的热点相关路由均在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    HotRoutes
// @description   给gin引擎挂上热点相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func HotRoutes(r *gin.Engine) *gin.Engine {
	// TODO 好友的路由分组
	hotRoutes := r.Group("/hot")

	// TODO 添加中间件
	hotRoutes.Use(middleware.AuthMiddleware())

	// TODO 文章热度控制器
	articleHotController := controller.NewArticleHotController()

	// TODO 查看文章的游览次数
	hotRoutes.GET("/article/visit/:id", articleHotController.Visit)

	// TODO 看文章推荐
	hotRoutes.GET("/article/recomment", articleHotController.Recomment)

	// TODO 查看文章的热度排行
	hotRoutes.GET("/article", articleHotController.Ranking)

	// TODO 帖子热度控制器
	postHotController := controller.NewPostHotController()

	// TODO 查看帖子的跟帖人数
	hotRoutes.GET("/post/thread/:id", postHotController.Thread)

	// TODO 查看帖子的游览次数
	hotRoutes.GET("/post/visit/:id", postHotController.Visit)

	// TODO 查看帖子的热度排行
	hotRoutes.GET("/post", postHotController.Ranking)

	// TODO 查看帖子推荐
	hotRoutes.GET("/post/recomment", postHotController.Recomment)

	// TODO 前端文件热度控制器
	zipfileHotController := controller.NewFileHotController()

	// TODO 查看前端文件的游览次数
	hotRoutes.GET("/zipfile/visit/:id", zipfileHotController.Visit)

	// TODO 查看前端文件的下载次数
	hotRoutes.GET("/zipfile/download/:id", zipfileHotController.Download)

	// TODO 查看前端文件的使用人数
	hotRoutes.GET("/zipfile/use/:id", zipfileHotController.Use)

	// TODO 查看前端文件的评论人数
	hotRoutes.GET("/zipfile/comment/:id", zipfileHotController.Comment)

	// TODO 查看前端文件的热度排行
	hotRoutes.GET("/zipfile", zipfileHotController.Ranking)

	// TODO 查看前端文件推荐
	hotRoutes.GET("/zipfile/recomment", zipfileHotController.Recomment)

	// TODO 用户热度控制器
	userHotController := controller.NewUserHotController()

	// TODO 查看用户的热度排行
	hotRoutes.GET("/user", userHotController.Ranking)

	// TODO 查看用户的热度等级
	hotRoutes.GET("/user/level/:id", userHotController.Level)

	// TODO 用户查看自己热度等级
	hotRoutes.GET("/user/level", userHotController.MyLevel)

	// TODO 用户查看自己的简报
	hotRoutes.GET("/user/powerpoint", userHotController.MyPowerPoint)

	// TODO 查看用户的简报
	hotRoutes.GET("/user/powerpoint/:id", userHotController.PowerPoint)

	// TODO 查看用户推荐
	hotRoutes.GET("/user/recomment", userHotController.Recomment)

	return r
}
