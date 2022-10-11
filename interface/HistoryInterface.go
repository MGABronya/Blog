// @Title  HistoryInterface
// @Description  该文件用于封装历史记录相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// HistoryInterface			定义了历史记录相关方法
type HistoryInterface interface {
	Create(ctx *gin.Context)    // 创建历史记录
	Show(ctx *gin.Context)      // 查看历史记录
	Delete(ctx *gin.Context)    // 删除历史记录
	DeleteAll(ctx *gin.Context) // 删除所有历史记录
}
