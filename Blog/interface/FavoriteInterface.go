// @Title  FavoriteInterface
// @Description  该文件用于封装收藏相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// FavoriteInterface			定义了收藏相关方法
type FavoriteInterface interface {
	Create(ctx *gin.Context)           // 创建收藏
	Show(ctx *gin.Context)             // 查看收藏
	Delete(ctx *gin.Context)           // 删除收藏
	FavoriteList(ctx *gin.Context)     // 查看收藏列表
	UserFavoriteList(ctx *gin.Context) // 查看用户的收藏列表
}
