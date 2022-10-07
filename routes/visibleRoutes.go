// @Title  visibleRoutes
// @Description  程序的可见相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    VisibleRoutes
// @description   给gin引擎挂上可见相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func VisibleRoutes(r *gin.Engine) *gin.Engine {

	// TODO 后台管理的路由分组
	visibleRoutes := r.Group("/visible")

	// TODO 添加中间件
	visibleRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建帖子controller
	visibleController := controller.NewVisibleController()

	// TODO 设置文章可见等级
	visibleRoutes.PUT("/article/:id", visibleController.ArticleCreate)

	// TODO 设置帖子可见等级
	visibleRoutes.PUT("/post/:id", visibleController.PostCreate)

	// TODO 设置前端文件可见等级
	visibleRoutes.PUT("/zipfile/:id", visibleController.ZipFileCreate)

	// TODO 查看文章可见等级
	visibleRoutes.GET("/article/:id", visibleController.ArticleShow)

	// TODO 查看帖子可见等级
	visibleRoutes.GET("/post/:id", visibleController.PostShow)

	// TODO 查看前端文件可见等级
	visibleRoutes.GET("/zipfile/:id", visibleController.ZipFileShow)

	// TODO 设置帖子是否可以跟帖
	visibleRoutes.PUT("/post/thread/:id", visibleController.PostThread)

	// TODO 设置前端文件是否可以被评论
	visibleRoutes.PUT("/zipfile/comment/:id", visibleController.ZipfileComment)

	// TODO 查看帖子跟帖等级
	visibleRoutes.GET("/post/thread/:id", visibleController.PostThreadShow)

	// TODO 查看前端文件评论等级
	visibleRoutes.GET("/zipfile/comment/:id", visibleController.ZipfileCommentShow)

	// TODO 设置前端文件是否可以被下载
	visibleRoutes.PUT("/zipfile/download/:id", visibleController.ZipfileDownload)

	// TODO 查看前端文件下载等级
	visibleRoutes.GET("/zipfile/download/:id", visibleController.ZipfileDownloadShow)

	// TODO 查看帖子是否可以跟帖
	visibleRoutes.GET("/post/thread/can/:id", visibleController.PostThreadCan)

	// TODO 查看前端文件是否可以下载
	visibleRoutes.GET("/zipfile/can/:id", visibleController.ZipfileCan)

	// TODO 查看前端文件是否可以评论
	visibleRoutes.GET("/zipfile/comment/can/:id", visibleController.ZipfileCommentCan)

	return r
}
