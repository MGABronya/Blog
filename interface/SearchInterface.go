// @Title  SearchInterface
// @Description  该文件用于封装搜索相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// SeartchInterface			定义了搜索相关方法
type SeartchInterface interface {
	Show(ctx *gin.Context)                   // 实现搜索
	ShowWithLabelInter(ctx *gin.Context)     // 实现带标签交集搜索
	ShowWithLabelUnion(ctx *gin.Context)     // 实现带标签并集搜索
	ShowUser(ctx *gin.Context)               // 实现指定用户的搜索
	ShowWithLabelInterUser(ctx *gin.Context) // 实现指定用户的带标签交集搜索
	ShowWithLabelUnionUser(ctx *gin.Context) // 实现指定用户的带标签并集搜索
}
