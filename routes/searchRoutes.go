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

	// TODO 创建文章文本搜索controller
	articleSearchController := controller.NewArticleSearchController()

	// TODO 按文本搜索文章
	searchRoutes.GET("/article/:text", articleSearchController.Show)

	// TODO 按文本搜索指定用户的文章
	searchRoutes.GET("/article/:text/:id", articleSearchController.ShowUser)

	// TODO 按文本和标签交集搜索文章
	searchRoutes.GET("/article/inter/:text", articleSearchController.ShowWithLabelInter)

	// TODO 按文本和标签交集搜索指定用户的文章
	searchRoutes.GET("/article/inter/:text/:id", articleSearchController.ShowWithLabelInterUser)

	// TODO 按文本和标签并集搜索文章
	searchRoutes.GET("/article/union/:text", articleSearchController.ShowWithLabelUnion)

	// TODO 按文本和标签并集搜索指定用户的文章
	searchRoutes.GET("/article/union/:text/:id", articleSearchController.ShowWithLabelUnionUser)

	// TODO 创建文章文本搜索controller
	postSearchController := controller.NewPostSearchController()

	// TODO 按文本搜索帖子
	searchRoutes.GET("/post/:text", postSearchController.Show)

	// TODO 按文本搜索指定用户的帖子
	searchRoutes.GET("/post/:text/:id", postSearchController.ShowUser)

	// TODO 按文本和标签交集搜索帖子
	searchRoutes.GET("/post/inter/:text", postSearchController.ShowWithLabelInter)

	// TODO 按文本和标签交集搜索指定用户的帖子
	searchRoutes.GET("/post/inter/:text/:id", postSearchController.ShowWithLabelInterUser)

	// TODO 按文本和标签并集搜索帖子
	searchRoutes.GET("/post/union/:text", postSearchController.ShowWithLabelUnion)

	// TODO 按文本和标签并集搜索指定用户的帖子
	searchRoutes.GET("/post/union/:text/:id", postSearchController.ShowWithLabelUnionUser)

	// TODO 创建前端文件文本搜索controller
	fileSearchController := controller.NewFileSearchController()

	// TODO 按文本搜索帖子
	searchRoutes.GET("/zipfile/:text", fileSearchController.Show)

	// TODO 按文本搜索指定用户的帖子
	searchRoutes.GET("/zipfile/:text/:id", fileSearchController.ShowUser)

	// TODO 按文本和标签交集搜索帖子
	searchRoutes.GET("/zipfile/inter/:text", fileSearchController.ShowWithLabelInter)

	// TODO 按文本和标签交集搜索指定用户的帖子
	searchRoutes.GET("/zipfile/inter/:text/:id", fileSearchController.ShowWithLabelInterUser)

	// TODO 按文本和标签并集搜索帖子
	searchRoutes.GET("/zipfile/union/:text", fileSearchController.ShowWithLabelUnion)

	// TODO 按文本和标签并集搜索指定用户的帖子
	searchRoutes.GET("/zipfile/union/:text/:id", fileSearchController.ShowWithLabelUnionUser)

	return r
}
