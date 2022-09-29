// @Title  labelRoutes
// @Description  程序的标签相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    LabelRoute
// @description   给gin引擎挂上标签相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func LabelRoutes(r *gin.Engine) *gin.Engine {

	// TODO 文章标签路由
	articleLabelRoutes := r.Group("/articlelabel")

	// TODO 添加中间件
	articleLabelRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建文章标签controller
	articleLabelController := controller.NewArticleLabelController()

	// TODO 查看文章标签
	articleLabelRoutes.GET("/show/:id", articleLabelController.Show)

	// TODO 创建文章标签
	articleLabelRoutes.POST("/creat/:id", articleLabelController.Create)

	// TODO 删除文章标签
	articleLabelRoutes.DELETE("/delete/:id", articleLabelController.Delete)

	// TODO 前端文件标签路由
	fileLabelRoutes := r.Group("/filelabel")

	// TODO 添加中间件
	fileLabelRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建前端文件标签controller
	fileLabelController := controller.NewFileLabelController()

	// TODO 查看前端文件标签
	fileLabelRoutes.GET("/show/:id", fileLabelController.Show)

	// TODO 创建前端文件标签
	fileLabelRoutes.POST("/creat/:id", fileLabelController.Create)

	// TODO 删除前端文件标签
	fileLabelRoutes.DELETE("/delete/:id", fileLabelController.Delete)

	// TODO 帖子标签路由
	postLabelRoutes := r.Group("/postlabel")

	// TODO 添加中间件
	postLabelRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建帖子标签controller
	postLabelController := controller.NewPostLabelController()

	// TODO 查看帖子标签
	postLabelRoutes.GET("/show/:id", postLabelController.Show)

	// TODO 创建帖子标签
	postLabelRoutes.POST("/creat/:id", postLabelController.Create)

	// TODO 删除帖子标签
	postLabelRoutes.DELETE("/delete/:id", postLabelController.Delete)

	// TODO 用户标签路由
	userLabelRoutes := r.Group("/userlabel")

	// TODO 添加中间件
	userLabelRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建用户标签controller
	userLabelController := controller.NewUserLabelController()

	// TODO 用户自身标签查看
	userLabelRoutes.GET("/show", userLabelController.Show)

	// TODO 查看指定用户的标签
	userLabelRoutes.GET("/show/:id", userLabelController.ShowLabel)

	// TODO 创建用户标签
	userLabelRoutes.POST("/creat", userLabelController.Create)

	// TODO 删除用户标签
	userLabelRoutes.DELETE("/delete", userLabelController.Delete)

	// TODO 删除指定用户的标签
	userLabelRoutes.DELETE("/delete/:id", userLabelController.DeleteLabel)

	return r
}
