// @Title  HotInterface
// @Description  该文件用于封装热度相关方法
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package Interface

import "github.com/gin-gonic/gin"

// HotInterface			定义了热度相关方法
type HotInterface interface {
	Ranking(ctx *gin.Context)   // 热度排行
	Recomment(ctx *gin.Context) // 推荐
}
