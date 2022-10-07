// @Title  searchRoutes
// @Description  程序的文本搜索相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    SearchRoutes
// @description   给gin引擎挂上文本搜索相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func SearchRoutes(r *gin.Engine) *gin.Engine {
	// TODO 文本搜索的路由分组
	searchRoutes := r.Group("/search")

	// TODO 添加中间件
	searchRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建文本搜索controller
	searchController := controller.NewSearchController()

	// TODO 按文本搜索文章
	searchRoutes.GET("/article/:text", searchController.Article)

	// TODO 按文本搜索帖子
	searchRoutes.GET("/post/:text", searchController.Post)

	// TODO 按文本搜索前端文件
	searchRoutes.GET("/zipfile/:text", searchController.Zipfile)

	// TODO 按文本和标签交集搜索文章
	searchRoutes.GET("/article/inter/:text", searchController.ArticleWithLabelInter)

	// TODO 按文本和标签交集搜索帖子
	searchRoutes.GET("/post/inter/:text", searchController.PostWithLabelInter)

	// TODO 按文本和标签交集搜索前端文件
	searchRoutes.GET("/zipfile/inter/:text", searchController.ZipfileWithLabelInter)

	// TODO 按文本和标签并集搜索文章
	searchRoutes.GET("/article/union/:text", searchController.ArticleWithLabelUnion)

	// TODO 按文本和标签并集搜索帖子
	searchRoutes.GET("/post/union/:text", searchController.PostWithLabelUnion)

	// TODO 按文本和标签并集搜索前端文件
	searchRoutes.GET("/zipfile/union/:text", searchController.ZipfileWithLabelUnion)

	return r
}
