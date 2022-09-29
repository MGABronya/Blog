// @Title  likeRoutes
// @Description  程序的点赞相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    LikeRoutes
// @description   给gin引擎挂上点赞相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func LikeRoutes(r *gin.Engine) *gin.Engine {

	// TODO 文章点赞的路由分组
	likeArticleRoutes := r.Group("/articlelike")

	// TODO 添加中间件
	likeArticleRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建文章点赞controller
	articleLikeController := controller.NewArticleLikeController()

	// TODO 查看文章是否点赞
	likeArticleRoutes.GET("/show/:id", articleLikeController.Show)

	// TODO 创建文章点赞
	likeArticleRoutes.PUT("/creat/:id", articleLikeController.Create)

	// TODO 删除文章点赞
	likeArticleRoutes.DELETE("/delete/:id", articleLikeController.Delete)

	// TODO 查看文章点赞列表
	likeArticleRoutes.GET("/list/:id", articleLikeController.LikeList)

	// TODO 帖子点赞的路由分组
	likePostRoutes := r.Group("/postlike")

	// TODO 添加中间件
	likePostRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建帖子点赞controller
	postLikeController := controller.NewPostLikeController()

	// TODO 查看帖子点赞
	likePostRoutes.GET("/show/:id", postLikeController.Show)

	// TODO 创建帖子点赞
	likePostRoutes.PUT("/creat/:id", postLikeController.Create)

	// TODO 删除帖子点赞
	likePostRoutes.DELETE("/delete/:id", postLikeController.Delete)

	// TODO 查看帖子点赞列表
	likePostRoutes.GET("/list/:id", postLikeController.LikeList)

	// TODO 跟帖点赞的路由分组
	likeThreadRoutes := r.Group("/threadlike")

	// TODO 添加中间件
	likeThreadRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建跟帖点赞controller
	threadLikeController := controller.NewThreadLikeController()

	// TODO 查看跟帖点赞
	likeThreadRoutes.GET("/show/:id", threadLikeController.Show)

	// TODO 创建跟帖点赞
	likeThreadRoutes.PUT("/creat/:id", threadLikeController.Create)

	// TODO 删除跟帖点赞
	likeThreadRoutes.DELETE("/delete/:id", threadLikeController.Delete)

	// TODO 查看跟帖点赞列表
	likeThreadRoutes.GET("/list/:id", threadLikeController.LikeList)

	// TODO 前端文件点赞的路由分组
	likeFileRoutes := r.Group("/filelike")

	// TODO 添加中间件
	likeFileRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建前端文件点赞controller
	fileLikeController := controller.NewFileLikeController()

	// TODO 查看前端文件点赞
	likeFileRoutes.GET("/show/:id", fileLikeController.Show)

	// TODO 创建前端文件点赞
	likeFileRoutes.PUT("/creat/:id", fileLikeController.Create)

	// TODO 删除前端文件点赞
	likeFileRoutes.DELETE("/delete/:id", fileLikeController.Delete)

	// TODO 查看前端文件点赞列表
	likeFileRoutes.GET("/list/:id", fileLikeController.LikeList)

	// TODO 评论点赞的路由分组
	likeCommentRoutes := r.Group("/commentlike")

	// TODO 添加中间件
	likeCommentRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建评论点赞controller
	commentLikeController := controller.NewCommentLikeController()

	// TODO 查看评论点赞
	likeCommentRoutes.GET("/show/:id", commentLikeController.Show)

	// TODO 创建评论点赞
	likeCommentRoutes.PUT("/creat/:id", commentLikeController.Create)

	// TODO 删除评论点赞
	likeCommentRoutes.DELETE("/delete/:id", commentLikeController.Delete)

	// TODO 查看评论点赞列表
	likeCommentRoutes.GET("/list/:id", commentLikeController.LikeList)

	return r
}
