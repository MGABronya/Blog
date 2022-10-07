// @Title  main
// @Description  程序的入口，读取配置，调用初始化函数以及运行路由
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:48
package main

import (
	"Blog/controller"
	"Blog/model"
	"Blog/util"
	"fmt"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// @title    main
// @description   程序入口，完成一些初始化工作后将开始监听
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func main() {
	InitConfig()
	common.InitDB()
	client0 := common.InitRedis(0)
	defer client0.Close()
	client1 := common.InitRedis(1)
	defer client1.Close()
	client2 := common.InitRedis(2)
	defer client2.Close()
	client3 := common.InitRedis(3)
	defer client3.Close()
	client4 := common.InitRedis(4)
	defer client4.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")

	if port != "" {
		panic(r.Run(":" + port))
	}

	// TODO 定时下降热度
	go func() {
		for {
			log.Println("cleanLimitQueue start, hot downing...")
			// TODO 执行功能
			HotDown()
			now := time.Now()
			// TODO 计算下一个3:30
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 3, 30, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()

	// TODO 定时生成用户简报
	go func() {
		for {
			log.Println("cleanLimitQueue start, powerpoint making...")
			// TODO 执行功能
			PowerPoint()
			now := time.Now()
			// TODO 计算下一个4:00
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 4, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()

	// TODO 定时生成每日推荐
	go func() {
		for {
			log.Println("cleanLimitQueue start, recommend making...")
			// TODO 执行功能
			Recomment()
			now := time.Now()
			// TODO 计算下一个4:30
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 4, 30, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()

	panic(r.Run())
}

// @title    InitConfig
// @description   读取配置文件并完成初始化
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	// TODO 如果发生错误，终止程序
	if err != nil {
		panic(err)
	}
}

// @title    HotDown
// @description   完成一轮的热度下降
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func HotDown() {
	// TODO 文章热度下降
	articles := util.RangeWithScoreZ(1, "H", 0, -1)
	for _, article := range articles {
		articleId := fmt.Sprint(article.Member)
		util.AddZ(1, "H", articleId, article.Score*0.99)
	}
	// TODO 帖子热度下降
	posts := util.RangeWithScoreZ(3, "H", 0, -1)
	for _, post := range posts {
		postId := fmt.Sprint(post.Member)
		util.AddZ(3, "H", postId, post.Score*0.99)
	}
	// TODO 前端文件热度下降
	zipfiles := util.RangeWithScoreZ(2, "H", 0, -1)
	for _, zipfile := range zipfiles {
		zipfileId := fmt.Sprint(zipfile.Member)
		util.AddZ(2, "H", zipfileId, zipfile.Score*0.99)
	}
	// TODO 用户热度下降
	users := util.RangeWithScoreZ(4, "H", 0, -1)
	for _, user := range users {
		userId := fmt.Sprint(user.Member)
		down := 0.0
		for i := 7; i >= 0; i-- {
			if controller.HotLevel[i] < user.Score {
				down = user.Score - controller.HotLevel[i]
				break
			}
		}
		util.AddZ(4, "H", userId, user.Score*0.99-down*0.01)
	}
}

// @title    PowerPoint
// @description   完成用户报表更新
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func PowerPoint() {
	var users []gmodel.User
	db := common.GetDB()
	db.Find(&users)
	for _, user := range users {
		labels := util.RangeWithScoreZ(4, "L"+strconv.Itoa(int(user.ID)), 0, -1)
		// TODO 创建日报
		for _, label := range labels {
			powerpoint := model.PowerPoint{
				UserId: user.ID,
				Label:  fmt.Sprint(label.Member),
				Score:  label.Score,
			}
			// TODO 插入数据
			if err := db.Create(&powerpoint).Error; err != nil {
				panic(err)
			}
		}
	}
}

// @title    Recomment
// @description   每日推荐
// @auth      MGAronya（张健）             2022-9-16 10:49
// @param     void			没有入参
// @return    void			没有回参
func Recomment() {
	var users []gmodel.User
	db := common.GetDB()
	db.Find(&users)
	for _, user := range users {
		// TODO 清空推荐
		userId := strconv.Itoa(int(user.ID))
		util.Del(1, "R"+userId)
		util.Del(2, "R"+userId)
		util.Del(3, "R"+userId)
		util.Del(4, "R"+userId)
		articleUnion := make(map[string]float64)
		postUnion := make(map[string]float64)
		zipfileUnion := make(map[string]float64)
		userUnion := make(map[string]float64)
		// TODO 找出并集
		labels := util.RangeWithScoreZ(4, "L"+userId, 0, -1)
		for _, label := range labels {
			articles := util.MembersS(1, "La"+fmt.Sprint(label.Member))
			for _, article := range articles {
				articleUnion[article] += label.Score
			}
			posts := util.MembersS(3, "La"+fmt.Sprint(label.Member))
			for _, post := range posts {
				postUnion[post] += label.Score
			}
			zipfiles := util.MembersS(2, "La"+fmt.Sprint(label.Member))
			for _, zipfile := range zipfiles {
				zipfileUnion[zipfile] += label.Score
			}
			users := util.MembersS(4, "La"+fmt.Sprint(label.Member))
			for _, user := range users {
				userUnion[user] += label.Score
			}
		}
		// TODO 加入文章热点信息
		articles := util.RangeZ(1, "H", 0, 9)
		for _, article := range articles {
			articleUnion[article] += 0.01
		}

		// TODO 加入帖子热点信息
		posts := util.RangeZ(3, "H", 0, 9)
		for _, post := range posts {
			postUnion[post] += 0.01
		}

		// TODO 加入前端文件热点信息
		zipfiles := util.RangeZ(2, "H", 0, 9)
		for _, zipfile := range zipfiles {
			zipfileUnion[zipfile] += 0.01
		}

		// TODO 加入用户热点信息
		users := util.RangeZ(4, "H", 0, 9)
		for _, user := range users {
			userUnion[user] += 0.01
		}

		// TODO 文章推荐
		for article := range articleUnion {
			var a gmodel.Article
			// TODO 查看是否有权限
			db.Where("id = ?", article).First(&a)
			if a.UserId != user.ID && (a.Visible == 3 || (a.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(a.UserId))))) {
				continue
			}
			score := util.ScoreZ(1, "H", article)
			util.AddZ(1, "R"+userId, article, score*(1.0+0.01*articleUnion[article]))
		}

		// TODO 帖子推荐
		for post := range postUnion {
			var p gmodel.Post
			// TODO 查看是否有权限
			db.Where("id = ?", post).First(&p)
			if p.UserId != user.ID && (p.Visible == 3 || (p.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(p.UserId))))) {
				continue
			}
			score := util.ScoreZ(3, "H", post)
			util.AddZ(3, "R"+userId, post, score*(1.0+0.01*postUnion[post]))
		}

		// TODO 前端文件推荐
		for zipfile := range zipfileUnion {
			var z model.ZipFile
			// TODO 查看是否有权限
			db.Where("id = ?", zipfile).First(z)
			if z.UserId != user.ID && (z.Visible == 3 || (z.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(z.UserId))))) {
				continue
			}
			score := util.ScoreZ(2, "H", zipfile)
			util.AddZ(2, "R"+userId, zipfile, score*(1.0+0.01*zipfileUnion[zipfile]))
		}

		// TODO 用户推荐
		for user := range userUnion {
			score := util.ScoreZ(4, "H", user)
			util.AddZ(4, "R"+userId, user, score*(1.0+0.01*articleUnion[user]))
		}

	}

}
