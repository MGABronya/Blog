// @Title  friendRoutes
// @Description  程序的好友相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    FriendRoutes
// @description   给gin引擎挂上好友相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func FriendRoutes(r *gin.Engine) *gin.Engine {
	// TODO 好友的路由分组
	friendRoutes := r.Group("/friend")

	// TODO 添加中间件
	friendRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建好友controller
	friendController := controller.NewFriendController()

	// TODO 查看好友列表
	friendRoutes.GET("/show", friendController.Show)

	// TODO 查看好友申请列表
	friendRoutes.GET("/show/applied", friendController.ShowApplied)

	// TODO 查看申请列表
	friendRoutes.GET("/show/applying", friendController.ShowApplying)

	// TODO 发送好友申请
	friendRoutes.PUT("/applying/:id", friendController.Applying)

	// TODO 接受好友申请
	friendRoutes.PUT("/applied/:id", friendController.Applied)

	// TODO 拒绝好友申请
	friendRoutes.PUT("/refused/:id", friendController.Refused)

	// TODO 删除好友
	friendRoutes.DELETE("/delete/:id", friendController.Delete)

	// TODO 获取好友圈文章
	friendRoutes.GET("/articles", friendController.ShowArticles)

	// TODO 获取好友圈帖子
	friendRoutes.GET("/posts", friendController.ShowPosts)

	// TODO 获取好友圈前端文件
	friendRoutes.GET("/zipfiles", friendController.ShowZipfiles)

	return r
}
