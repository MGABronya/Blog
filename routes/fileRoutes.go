// @Title  fileRoutes
// @Description  程序的前端文件相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    FileRoute
// @description   给gin引擎挂上前端文件相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func FileRoute(r *gin.Engine) *gin.Engine {

	// TODO 上传个人前端的路由分组
	fileRoutes := r.Group("/zipfile")

	// TODO 添加中间件
	fileRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建个人前端controller
	fileController := controller.NewFileController()

	// TODO 用户个人上传前端文件
	fileRoutes.POST("/upload", fileController.Create)

	// TODO 用户下载前端文件
	fileRoutes.GET("/download/:id", fileController.Download)

	// TODO 用户个人删除前端文件
	fileRoutes.DELETE("/delete/:id", fileController.Delete)

	// TODO 前端页面列表
	fileRoutes.GET("/showlist", fileController.FileList)

	// TODO 获取用户自己上传的前端页面列表
	fileRoutes.GET("/show/mine", fileController.FileListMine)

	// TODO 获取某一用户上传的前端页面列表
	fileRoutes.GET("/show/others/:id", fileController.FileListOthers)

	// TODO 前端页面信息查询
	fileRoutes.GET("/show/:id", fileController.Show)

	// TODO 用户更新前端信息
	fileRoutes.PUT("/update/:id", fileController.Update)

	// TODO 用户上传前端描述图片
	fileRoutes.POST("/img/create/:id", fileController.CreateImg)

	// TODO 用户删除前端描述图片
	fileRoutes.DELETE("/img/delete/:id", fileController.DeleteImg)

	// TODO 用户查看前端描述图片
	fileRoutes.GET("/img/show/:id", fileController.ShowImg)

	// TODO 用户选择博客主页
	fileRoutes.GET("/choose/:id", fileController.Choose)

	// TODO 评论的路由分组
	commentRoutes := r.Group("/comment")

	// TODO 添加中间件
	commentRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建评论controller
	commentController := controller.NewCommentController()

	// TODO 评论的创建路由
	commentRoutes.POST("/:id", commentController.Create)

	// TODO 评论的更新路由
	commentRoutes.PUT("/:id", commentController.Update)

	// TODO 评论的查看路由
	commentRoutes.GET("/:id", commentController.Show)

	// TODO 评论的删除路由
	commentRoutes.DELETE("/:id", commentController.Delete)

	//TODO 评论的列表路由
	commentRoutes.GET("/pagelist/:id", commentController.PageList)

	// TODO 用户查看自己的评论列表
	commentRoutes.GET("/pagelist/mine", commentController.PageListMine)

	// TODO 查看某一用户的评论列表
	commentRoutes.GET("/pagelist/others/:id", commentController.PageListOthers)

	return r
}
