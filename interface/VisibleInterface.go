// @Title  VisibleInterface
// @Description  该文件用于封装可见度相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// VisibleInterface			定义了可见度相关方法
type VisibleInterface interface {
	Create(ctx *gin.Context) // 设置可见度
	Show(ctx *gin.Context)   // 查看可见度
}
