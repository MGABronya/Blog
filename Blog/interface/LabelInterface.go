// @Title  LabelInterface
// @Description  该文件用于封装标签相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// LabelInterface			定义了标签相关方法
type LabelInterface interface {
	Create(ctx *gin.Context) // 创建标签
	Show(ctx *gin.Context)   // 查看标签
	Delete(ctx *gin.Context) // 删除标签
}
