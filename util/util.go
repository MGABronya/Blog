// @Title  util
// @Description  收集各种需要使用的工具函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package util

import (
	"archive/zip"
	"context"
	"ginEssential/common"
	"io"

	"os"
	"path"
	"path/filepath"
)

// @title    Unzip
// @description  用于解压缩文件
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     zipPath, dstDir string		压缩文件的路径和目标路径
// @return    error    查看是否发生错误
func Unzip(zipPath, dstDir string) error {
	// TODO 打开压缩文件夹
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dstDir); err != nil {
			return err
		}
	}
	return nil
}

// @title    Unzip
// @description  用于解压缩文件
// @auth      MGAronya（张健）             2022-9-16 10:29
// @param     file *zip.File, dstDir string		压缩文件的指针和目标路径
// @return    error    查看是否发生错误
func unzipFile(file *zip.File, dstDir string) error {
	// TODO 为解压文件建立路径
	filePath := path.Join(dstDir, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// TODO 打开文件夹
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// TODO 创建文件
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// save the decompressed file content
	_, err = io.Copy(w, rc)
	return err
}

var ctx context.Context = context.Background()

// @title    GetH
// @description   在redis中的一个哈希中获取值
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, H string, key string        k表示选用第几个库，H为哈希，key为在H中的key
// @return   string     	  返回对应的value
func GetH(k int, H string, key string) string {
	client := common.GetRedisClient(k)
	level, _ := client.HGet(ctx, H, key).Result()
	return level
}

// @title    SETH
// @description   在redis中的一个哈希中设置值
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, H string, key string, value string       k表示选用第几个库，H为哈希，key为在H中的key，value为要设置的对应值
// @return   void
func SetH(k int, H string, key string, value string) {
	client := common.GetRedisClient(k)
	client.HSet(ctx, H, key, value)
}

// @title    GetS
// @description   在redis中的一个集合中获取值
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, S string        k表示选用第几个库，S为集合
// @return   string     	  返回对应的value
func GetS(k int, S string) []string {
	client := common.GetRedisClient(k)
	value, _ := client.SMembers(ctx, S).Result()
	return value
}

// @title    SetS
// @description   在redis中的一个集合中设置值
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, S string, value string       k表示选用的第几个库，S为集合，value为要设置的对应值
// @return   void
func SetS(k int, S string, value string) {
	client := common.GetRedisClient(k)
	client.SAdd(ctx, S, value)
}

// @title    RemS
// @description   在redis中的一个集合中删除值
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, S string, value string       k表示选用的第几个库，S为集合，value为要删除的对应值
// @return   void
func RemS(k int, S string, value string) {
	client := common.GetRedisClient(k)
	client.SRem(ctx, S, value)
}

// @title    IsS
// @description   在redis中的一个集合中查找某个元素是否存在
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, S string, value string       k表示选用的第几个库，S为集合，value为要查找的对应值
// @return   bool 表示value是否在S中
func IsS(k int, S string, value string) bool {
	client := common.GetRedisClient(k)
	flag, _ := client.SIsMember(ctx, S, value).Result()
	return flag
}

// @title    MembersS
// @description   在redis中的一个集合中查找所有元素
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, S string      k表示选用的第几个库，S为集合
// @return   []string		表示该集合的所有元素
func MembersS(k int, S string) []string {
	client := common.GetRedisClient(k)
	es, _ := client.SMembers(ctx, S).Result()
	return es
}

// @title    CardS
// @description  查看redis中一个集合中元素的个数
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, Key string      k表示选用的第几个库，Key为要查看的集合
// @return   int	表示该集合元素的个数
func CardS(k int, Key string) int {
	client := common.GetRedisClient(k)
	cnt, _ := client.SCard(ctx, Key).Result()
	return int(cnt)
}

// @title    Del
// @description  删除redis中的一个键
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, Key string      k表示选用的第几个库，Key为要删除的键
// @return   void, void
func Del(k int, Key string) {
	client := common.GetRedisClient(k)
	client.Del(ctx, Key)
}

// @title    DelH
// @description  删除redis中的一个哈希下的键
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    k int, h string, key string      k表示选用的第几个库，h表示哪个哈希表，key为要删除的键
// @return   void, void
func DelH(k int, h string, key string) {
	client := common.GetRedisClient(k)
	client.HDel(ctx, h, key)
}

// @title    PostThread
// @description  查看帖子是否可以跟帖
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    userIdA string, postId string, userIdB string     userIdA表示当前用户，postId是帖子的id，userIdB是帖子的作者id
// @return   bool  返回值表示是否可以跟帖
func PostThread(userIdA string, postId string, userIdB string) bool {
	level := GetH(3, "W", postId)
	// TODO 如果是作者本人，返回true
	if userIdA == userIdB {
		return true
	}
	// TODO 全开放等级
	if level == "1" {
		return true
	}
	// TODO 限制等级
	if level == "3" {
		return false
	}
	// TODO 查看是否为好友
	return IsS(4, "Fr"+userIdA, userIdB)
}

// @title    Zipfile
// @description  查看前端文件是否可以下载
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    userIdA string, zipfileId string, userIdB string     userIdA表示当前用户，zipfileId是前端文件的id，userIdB是前端文件的作者id
// @return   bool  返回值表示是否可以下载
func Zipfile(userIdA string, zipfileId string, userIdB string) bool {
	level := GetH(2, "D", zipfileId)
	// TODO 如果是作者本人，返回true
	if userIdA == userIdB {
		return true
	}
	// TODO 全开放等级
	if level == "1" {
		return true
	}
	// TODO 限制等级
	if level == "3" {
		return false
	}
	// TODO 查看是否为好友
	return IsS(4, "Fr"+userIdA, userIdB)
}

// @title    ZipfileComment
// @description  查看前端文件是否可以评论
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    userIdA string, zipfileId string, userIdB string     userIdA表示当前用户，zipfileId是前端文件的id，userIdB是前端文件的作者id
// @return   bool  返回值表示是否可以评论
func ZipfileComment(userIdA string, zipfileId string, userIdB string) bool {
	level := GetH(2, "W", zipfileId)
	// TODO 如果是作者本人，返回true
	if userIdA == userIdB {
		return true
	}
	// TODO 全开放等级
	if level == "1" {
		return true
	}
	// TODO 限制等级
	if level == "3" {
		return false
	}
	// TODO 查看是否为好友
	return IsS(4, "Fr"+userIdA, userIdB)
}
