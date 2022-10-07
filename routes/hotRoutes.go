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

	hotController := controller.NewHotController()

	// TODO 查看文章的游览次数
	hotRoutes.GET("/article/visit/:id", hotController.ArticleVisit)

	// TODO 查看帖子的游览次数
	hotRoutes.GET("/post/visit/:id", hotController.PostVisit)

	// TODO 查看前端文件的游览次数
	hotRoutes.GET("/zipfile/visit/:id", hotController.ZipfileVisit)

	// TODO 查看前端文件的下载次数
	hotRoutes.GET("/zipfile/download/:id", hotController.ZipfileDownload)

	// TODO 查看前端文件的使用人数
	hotRoutes.GET("/zipfile/use/:id", hotController.ZipfileUse)

	// TODO 查看前端文件的评论人数
	hotRoutes.GET("/zipfile/comment/:id", hotController.ZipfileComment)

	// TODO 查看帖子的跟帖人数
	hotRoutes.GET("/post/thread/:id", hotController.PostThread)

	// TODO 查看文章的热度排行
	hotRoutes.GET("/article", hotController.ArticleRanking)

	// TODO 查看帖子的热度排行
	hotRoutes.GET("/post", hotController.PostRanking)

	// TODO 查看前端文件的热度排行
	hotRoutes.GET("/zipfile", hotController.ZipfileRanking)

	// TODO 查看用户的热度排行
	hotRoutes.GET("/user", hotController.UserRanking)

	// TODO 查看用户的热度等级
	hotRoutes.GET("/user/level/:id", hotController.UserLevel)

	// TODO 用户查看自己热度等级
	hotRoutes.GET("/user/level", hotController.MyLevel)

	// TODO 用户查看自己的简报
	hotRoutes.GET("/user/powerpoint", hotController.MyPowerPoint)

	// TODO 查看用户的简报
	hotRoutes.GET("/user/powerpoint/:id", hotController.PowerPoint)

	// TODO 看文章推荐
	hotRoutes.GET("/article/recomment", hotController.ArticleRecomment)

	// TODO 查看帖子推荐
	hotRoutes.GET("/post/recomment", hotController.PostRecomment)

	// TODO 查看前端文件推荐
	hotRoutes.GET("/zipfile/recomment", hotController.ZipfileRecomment)

	// TODO 查看用户推荐
	hotRoutes.GET("/user/recomment", hotController.UserRecomment)

	return r
}
