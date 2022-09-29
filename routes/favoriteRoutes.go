// @Title  favoriteRoutes
// @Description  程序的收藏相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    FavoriteRoutes
// @description   给gin引擎挂上收藏相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func FavoriteRoutes(r *gin.Engine) *gin.Engine {

	// TODO 文章收藏的路由分组
	favoriteArticleRoutes := r.Group("/articlefavorite")

	// TODO 添加中间件
	favoriteArticleRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建文章收藏controller
	articleFavoriteController := controller.NewArticleFavoriteController()

	// TODO 查看文章是否收藏
	favoriteArticleRoutes.GET("/show/:id", articleFavoriteController.Show)

	// TODO 创建文章收藏
	favoriteArticleRoutes.PUT("/creat/:id", articleFavoriteController.Create)

	// TODO 删除文章收藏
	favoriteArticleRoutes.DELETE("/delete/:id", articleFavoriteController.Delete)

	// TODO 查看文章收藏列表
	favoriteArticleRoutes.GET("/list/:id", articleFavoriteController.FavoriteList)

	// TODO 查看用户收藏的文章列表
	favoriteArticleRoutes.GET("/userlist", articleFavoriteController.UserFavoriteList)

	// TODO 帖子收藏的路由分组
	favoritePostRoutes := r.Group("/postfavorite")

	// TODO 添加中间件
	favoritePostRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建帖子收藏controller
	postFavoriteController := controller.NewPostFavoriteController()

	// TODO 查看帖子收藏
	favoritePostRoutes.GET("/show/:id", postFavoriteController.Show)

	// TODO 创建帖子收藏
	favoritePostRoutes.PUT("/creat/:id", postFavoriteController.Create)

	// TODO 删除帖子收藏
	favoritePostRoutes.DELETE("/delete/:id", postFavoriteController.Delete)

	// TODO 查看帖子收藏列表
	favoritePostRoutes.GET("/list/:id", postFavoriteController.FavoriteList)

	// TODO 查看用户收藏的帖子列表
	favoritePostRoutes.GET("/userlist", postFavoriteController.UserFavoriteList)

	// TODO 跟帖收藏的路由分组
	favoriteThreadRoutes := r.Group("/threadfavorite")

	// TODO 添加中间件
	favoriteThreadRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建跟帖收藏controller
	threadFavoriteController := controller.NewThreadFavoriteController()

	// TODO 查看跟帖收藏
	favoriteThreadRoutes.GET("/show/:id", threadFavoriteController.Show)

	// TODO 创建跟帖收藏
	favoriteThreadRoutes.PUT("/creat/:id", threadFavoriteController.Create)

	// TODO 删除跟帖收藏
	favoriteThreadRoutes.DELETE("/delete/:id", threadFavoriteController.Delete)

	// TODO 查看跟帖收藏列表
	favoriteThreadRoutes.GET("/list/:id", threadFavoriteController.FavoriteList)

	// TODO 查看用户收藏的跟帖列表
	favoriteThreadRoutes.GET("/userlist", threadFavoriteController.UserFavoriteList)

	// TODO 前端文件收藏的路由分组
	favoriteFileRoutes := r.Group("/zipfilefavorite")

	// TODO 添加中间件
	favoriteFileRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建前端文件收藏controller
	fileFavoriteController := controller.NewFileFavoriteController()

	// TODO 查看前端文件收藏
	favoriteFileRoutes.GET("/show/:id", fileFavoriteController.Show)

	// TODO 创建前端文件收藏
	favoriteFileRoutes.PUT("/creat/:id", fileFavoriteController.Create)

	// TODO 删除前端文件文件收藏
	favoriteFileRoutes.DELETE("/delete/:id", fileFavoriteController.Delete)

	// TODO 查看前端文件收藏列表
	favoriteFileRoutes.GET("/list/:id", fileFavoriteController.FavoriteList)

	// TODO 查看用户收藏的前端文件列表
	favoriteFileRoutes.GET("/userlist", fileFavoriteController.UserFavoriteList)

	return r
}
