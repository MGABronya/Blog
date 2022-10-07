// @Title  guestbookRoutes
// @Description  程序的留言板相关路由均集中在这个文件里
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:50
package routes

import (
	"Blog/controller"
	"ginEssential/middleware"

	"github.com/gin-gonic/gin"
)

// @title    GuestbookRoutes
// @description   给gin引擎挂上留言板相关的路由监听
// @auth      MGAronya（张健）             2022-9-16 10:52
// @param     r *gin.Engine			gin引擎
// @return    r *gin.Engine			gin引擎
func GuestbookRoutes(r *gin.Engine) *gin.Engine {

	// TODO 留言板的路由分组
	guestbookRoutes := r.Group("/guestbook")

	// TODO 添加中间件
	guestbookRoutes.Use(middleware.AuthMiddleware())

	// TODO 创建后台controller
	guestbookController := controller.NewGuestbookController()

	// TODO 用户查看自己的留言板
	guestbookRoutes.GET("", guestbookController.Show)

	// TODO 查看指定用户的留言板
	guestbookRoutes.GET("/:id", guestbookController.ShowOthers)

	// TODO 用户给某个用户留言
	guestbookRoutes.POST("/:id", guestbookController.Create)

	// TODO 用户更新留言
	guestbookRoutes.PUT("/:id", guestbookController.Update)

	// TODO 用户删除自己留言板中的某条记录
	guestbookRoutes.DELETE("/:id", guestbookController.Delete)

	return r
}
