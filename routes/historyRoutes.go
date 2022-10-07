// @Title  historyRoutes
// @Description  程序的历史记录相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    HistoryRoutes
// @description   给gin引擎挂上历史记录相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func HistoryRoutes(r *gin.Engine) *gin.Engine {

	// TODO 文章收藏的路由分组
	historyRoutes := r.Group("/history")

	// TODO 添加中间件
	historyRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建文章收藏controller
	historyController := controller.NewHistoryController()

	// TODO 设置文章历史记录
	historyRoutes.POST("/article/:id", historyController.ArticleCreate)

	// TODO 设置帖子历史记录
	historyRoutes.POST("/post/:id", historyController.PostCreate)

	// TODO 设置前端文件历史记录
	historyRoutes.POST("/zipfile/:id", historyController.ZipfileCreate)

	// TODO 查看文章历史列表
	historyRoutes.GET("/article", historyController.AriticleShow)

	// TODO 查看帖子历史列表
	historyRoutes.GET("/post", historyController.PostShow)

	// TODO 查看前端文件历史列表
	historyRoutes.GET("/zipfile", historyController.ZipfileShow)

	// TODO 删除文章历史记录
	historyRoutes.DELETE("/article/:id", historyController.ArticleDelete)

	// TODO 删除帖子历史记录
	historyRoutes.DELETE("/post/:id", historyController.PostDelete)

	// TODO 删除前端文件历史记录
	historyRoutes.DELETE("/zipfile/:id", historyController.ZipfileDelete)

	// TODO 清空文章历史记录
	historyRoutes.DELETE("/article/all", historyController.ArticleDeleteAll)

	// TODO 清空帖子历史记录
	historyRoutes.DELETE("/post/all", historyController.PostDeleteAll)

	// TODO 清空前端文件历史记录
	historyRoutes.DELETE("/zipfile/all", historyController.ZipfileDeleteAll)

	return r
}
