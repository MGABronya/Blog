// @Title  UserFindRoutes
// @Description  程序的搜寻用户相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    UserFindRoutes
// @description   给gin引擎挂上搜寻用户相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func UserFindRoutes(r *gin.Engine) *gin.Engine {

	// TODO 搜寻用户的路由分组
	userFindRoutes := r.Group("/user")

	// TODO 添加中间件
	userFindRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建搜寻用户controller
	userFindController := controller.NewUserFindController()

	// TODO 通过名称查找用户
	userFindRoutes.GET("/name/:id", userFindController.NameFind)

	// TODO 通过email查找用户
	userFindRoutes.GET("/email/:id", userFindController.EmailFind)

	return r
}
