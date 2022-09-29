// @Title  LikeInterface
// @Description  该文件用于封装点赞相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// LikeInterface			定义了点赞相关方法
type LikeInterface interface {
	Create(ctx *gin.Context)   // 创建点赞
	Show(ctx *gin.Context)     // 查看点赞
	Delete(ctx *gin.Context)   // 删除点赞
	LikeList(ctx *gin.Context) // 查看点赞列表
}
