// @Title  routes
// @Description  程序的路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package main

import (
	"Blog/routes"

	"github.com/gin-gonic/gin"
)

// @title    CollectRoute
// @description   给gin引擎挂上路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func CollectRoute(r *gin.Engine) *gin.Engine {

	// TODO 上传个人前端路由
	r = routes.FileRoute(r)

	// TODO 后台管理路由
	r = routes.SystemRoutes(r)

	// TODO 标签路由
	r = routes.LabelRoutes(r)

	// TODO 点赞路由
	r = routes.LikeRoutes(r)

	// TODO 收藏路由
	r = routes.FavoriteRoutes(r)

	// TODO 好友路由
	r = routes.FriendRoutes(r)

	// TODO 设置可见路由
	r = routes.VisibleRoutes(r)

	return r
}
