// @Title  SystemRoutes
// @Description  程序的后台相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    SystemRoutes
// @description   给gin引擎挂上后台相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func SystemRoutes(r *gin.Engine) *gin.Engine {

	// TODO 后台管理的路由分组
	systemRoutes := r.Group("/system")

	// TODO 添加中间件
	systemRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建后台controller
	systemController := controller.NewSystemController()

	// TODO 用户设置权限等级
	systemRoutes.GET("/permission/:id/:level", systemController.Permission)

	// TODO 用户查看自身权限等级
	systemRoutes.GET("/permission", systemController.ShowPermission)

	return r
}
