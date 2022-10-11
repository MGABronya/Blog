// @Title  FileController
// @Description  该文件用于提供操作个人前端文件的各种函数
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:33
package controller

import (
	Interface "Blog/interface"
	"Blog/model"
	"Blog/util"
	"Blog/vo"
	"fmt"
	"ginEssential/common"
	gmodel "ginEssential/model"
	"ginEssential/response"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// IFileController			定义了文件类接口
type IFileController interface {
	Interface.RestInterface          // 实现增删查改功能
	FileList(ctx *gin.Context)       // 实现返回文件列表
	FileListMine(ctx *gin.Context)   // 实现查看自己的文件列表
	FileListOthers(ctx *gin.Context) //实现查看某一用户的文件列表
	Download(ctx *gin.Context)       // 实现下载文件功能
	Choose(ctx *gin.Context)         // 实现选择文件功能
	CreateImg(ctx *gin.Context)      // 实现上传文件描述图片功能
	DeleteImg(ctx *gin.Context)      // 实现删除文件描述图片功能
	ShowImg(ctx *gin.Context)        // 实现前端文件图片展示功能
}

// FileController			定义了文件工具类
type FileController struct {
	DB *gorm.DB //包含一个数据库指针
}

// @title    Create
// @description   用户个人上传前端文件
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Create(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	level := util.GetH(0, "permission", strconv.Itoa(int(user.ID)))

	if level[0] < '2' {
		response.Fail(ctx, nil, "权限等级不足")
		return
	}

	file, err := ctx.FormFile("file")

	//TODO 数据验证
	if err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	extName := path.Ext(file.Filename)

	// TODO 格式验证
	if extName != ".zip" {
		response.Fail(ctx, nil, "文件格式有误")
		return
	}

	// TODO 更新数据库信息
	newFile := model.ZipFile{
		UserId:  user.ID,
		Title:   user.Name,
		Content: "",
		Visible: 1,
	}

	if err := f.DB.Create(&newFile).Error; err != nil {
		panic(err)
	}

	file.Filename = fmt.Sprint(newFile.ID) + extName

	// TODO 将文件存入本地
	ctx.SaveUploadedFile(file, "./distzip/"+file.Filename)

	// TODO 更新热度
	util.AddZ(2, "H", newFile.ID.String(), 100)
	util.IncrByZ(4, "H", strconv.Itoa(int(user.ID)), 100)

	response.Success(ctx, gin.H{"file": newFile}, "上传文件成功")
}

// @title    Download
// @description   用户下载个人前端文件
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Download(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	id := ctx.Params.ByName("id")

	pfile := model.ZipFile{}

	// TODO 在数据库中查询file信息
	if f.DB.Where("id = ?", id).First(&pfile).Error != nil {
		response.Fail(ctx, nil, "压缩文件不存在")
		return
	}

	// TODO 查看是否有下载权限
	if !util.Zipfile(strconv.Itoa(int(user.ID)), pfile.ID.String(), strconv.Itoa(int(pfile.UserId))) {
		response.Fail(ctx, nil, "没有下载权限")
		return
	}

	// TODO 更新热度
	if !util.IsS(2, "T"+id, strconv.Itoa(int(user.ID))) {
		util.SetS(2, "T"+id, strconv.Itoa(int(user.ID)))
		util.IncrByZ(2, "H", id, 20)
		util.IncrByZ(4, "H", strconv.Itoa(int(pfile.UserId)), 20)
	}

	fileName := id + "zip"

	filePath := "./distzip/" + fileName

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.File(filePath)
}

// @title    Delete
// @description   个人删除前端文件
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Delete(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	id := ctx.Params.ByName("id")

	pfile := model.ZipFile{}

	// TODO 在数据库中查询file信息
	if f.DB.Where("id = ?", id).First(&pfile).Error != nil {
		response.Fail(ctx, nil, "压缩文件不存在")
		return
	}

	// TODO 查看文件删除权限
	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != pfile.UserId {
		response.Fail(ctx, nil, "权限不足，该文件不可删除")
		return
	}

	fileName := id + ".zip"

	filePath := "./distzip/" + fileName

	var comments []model.Comment

	// TODO 遍历每一条评论并移除热点
	f.DB.Where("file_id = ?", pfile.ID).Find(comments)

	for _, comment := range comments {
		DeleteCommentHot(&comment)
	}

	// TODO 删除相关评论
	f.DB.Where(model.Comment{FileId: pfile.ID.String()}).Delete(model.Comment{})

	// TODO 删除数据库中的相关数据
	f.DB.Delete(&pfile)

	// TODO 删除下载权限和评论权限
	util.DelH(2, "D", pfile.ID.String())
	util.DelH(2, "W", pfile.ID.String())

	// TODO 移除收藏
	for _, val := range util.MembersS(2, "aF"+id) {
		util.RemS(2, "Fa"+val, id)
	}
	util.Del(2, "aF"+id)

	// TODO 移除点赞
	util.Del(2, "fiL"+id)

	// TODO 移除标签
	for _, val := range util.MembersS(2, "aL"+id) {
		util.RemS(2, "La"+val, id)
	}
	util.Del(2, "aL"+id)

	// TODO 移除热度
	util.Del(2, "T"+id)
	util.Del(2, "C"+id)
	util.Del(2, "U"+id)
	util.Del(2, "M"+id)

	// TODO 更新热度
	util.IncrByZ(4, "H", strconv.Itoa(int(user.ID)), -util.ScoreZ(2, "H", id))
	util.RemZ(2, "H", id)

	// TODO 数据验证
	if err := os.Remove(filePath); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "该文件不存在")
		return
	}

	response.Success(ctx, nil, "删除文件成功")
}

// @title    FileList
// @description   查看前端文件列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) FileList(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	usera := tuser.(gmodel.User)

	users := util.MembersS(4, "Fr"+strconv.Itoa(int(usera.ID)))

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// TODO 分页
	var zipfiles []model.ZipFile

	// TODO 查找所有分页中可见的条目
	f.DB.Where("visible = 2 and user_id in (?)", users).Or("visible = 1").Or("visible = 3 and user_id = ?", usera.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	var total int64
	f.DB.Where("visible = 2 and user_id in (?)", users).Or("visible = 1").Or("visible = 3 and user_id = ?", usera.ID).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"files": zipfiles, "total": total}, "成功")
}

// @title    FileListMine
// @description   查看用户个人前端文件列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) FileListMine(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// TODO 查找所有zipfile
	var zipfiles []model.ZipFile
	f.DB.Where("user_id = ?", user.ID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	var total int64
	f.DB.Where("visible = 1").Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"files": zipfiles, "total": total}, "成功")
}

// @title    FileListOthers
// @description   查看某一用户的前端文件列表
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) FileListOthers(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// TODO 获取path中的id
	userId := ctx.Params.ByName("id")

	userB := gmodel.User{}

	// TODO 查看用户是否在数据库中存在
	if f.DB.Where("id = ?", userId).First(&userB).Error != nil {
		response.Fail(ctx, nil, "用户不存在")
		return
	}

	// TODO 分页
	var zipfiles []model.ZipFile
	var level int8
	if strconv.Itoa(int(user.ID)) == userId {
		level = 4
	} else if util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), userId) {
		level = 3
	} else {
		level = 2
	}

	// TODO 查找所有分页中可见的条目
	f.DB.Where("user_id = ? and visible < ?", user.ID, level).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&zipfiles)

	// TODO 记录的总条数
	var total int64
	f.DB.Where("user_id = ? and visible < ?", user.ID, level).Model(model.ZipFile{}).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"comments": zipfiles, "total": total}, "成功")
}

// @title    Show
// @description   查看前端文件
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Show(ctx *gin.Context) {

	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	file := model.ZipFile{}

	// TODO 查看前端文件是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&file).Error != nil {
		response.Fail(ctx, nil, "前端文件不存在")
		return
	}

	// TODO 查看是否有权限
	if file.UserId != user.ID && (file.Visible == 3 || (file.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(file.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	response.Success(ctx, gin.H{"file": file}, "成功")
}

// @title    Choose
// @description   选择前端文件
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Choose(ctx *gin.Context) {

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	file := model.ZipFile{}

	fileName := Id + ".zip"

	filePath := "./distzip/" + fileName

	// TODO 查看文件是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&file).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	// TODO 查看是否有权限
	if file.UserId != user.ID && (file.Visible == 3 || (file.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(file.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	os.RemoveAll("./dist/" + user.Name)

	os.Mkdir("./dist/"+user.Name, os.ModePerm)

	if err := util.Unzip(filePath, "./dist/"+user.Name); err != nil {
		response.Fail(ctx, nil, "前端文件解压失败")
		return
	}

	// TODO 更新热度
	util.RemS(2, "U"+util.GetH(4, "F", strconv.Itoa(int(user.ID))), strconv.Itoa(int(user.ID)))
	util.SetH(4, "F", strconv.Itoa(int(user.ID)), Id)
	if !util.IsS(2, "U"+Id, strconv.Itoa(int(user.ID))) {
		util.SetS(2, "U"+Id, strconv.Itoa(int(user.ID)))
		util.IncrByZ(2, "H", Id, 30)
		util.IncrByZ(4, "H", strconv.Itoa(int(file.UserId)), 1)
	}
	response.Success(ctx, nil, "选择成功")
}

// @title    Update
// @description   文件描述信息更新
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) Update(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	var requestZipfile vo.ZipFileRequest
	// TODO 数据验证
	if err := ctx.ShouldBind(&requestZipfile); err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	var zipFile model.ZipFile

	// TODO 查看文章是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&zipFile).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	if user.ID != zipFile.UserId {
		response.Fail(ctx, nil, "请勿非法修改他人信息")
		return
	}

	// TODO 更新文件信息
	zipFile.Title = requestZipfile.Title
	zipFile.Content = requestZipfile.Content

	f.DB.Save(&zipFile)

	response.Success(ctx, gin.H{"file": zipFile}, "更新成功")
}

// @title    CreateImg
// @description   文件信息描述图片上传
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) CreateImg(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	file, err := ctx.FormFile("file")

	//TODO 数据验证
	if err != nil {
		log.Print(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	var zipFile model.ZipFile

	// TODO 查看文件是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&zipFile).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	if user.ID != zipFile.UserId {
		response.Fail(ctx, nil, "请勿非法修改他人信息")
		return
	}

	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	// TODO 格式验证
	if _, ok := allowExtMap[extName]; !ok {
		response.Fail(ctx, nil, "文件格式有误")
		return
	}

	var fileImg model.FileImg

	fileImg.FileId = Id
	fileImg.UserId = user.ID

	file.Filename = fmt.Sprint(fileImg.ID) + extName

	// TODO 将文件存入本地
	ctx.SaveUploadedFile(file, "./FileImg/"+file.Filename)

	response.Success(ctx, gin.H{"fileImg": fileImg}, "更新成功")
}

// @title    DeleteImg
// @description   文件描述信息图片删除
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) DeleteImg(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	var fileImg model.FileImg

	// TODO 查看文章是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&fileImg).Error != nil {
		response.Fail(ctx, nil, "图片不存在")
		return
	}

	if util.GetH(0, "permission", strconv.Itoa(int(user.ID)))[0] <= '4' && user.ID != fileImg.UserId {
		response.Fail(ctx, nil, "请勿非法修改他人信息")
		return
	}

	f.DB.Delete(&fileImg)

	response.Success(ctx, nil, "删除成功")
}

// @title    ShowImg
// @description   文件描述信息图片展示
// @auth      MGAronya（张健）       2022-9-16 12:31
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func (f FileController) ShowImg(ctx *gin.Context) {
	tuser, _ := ctx.Get("user")
	user := tuser.(gmodel.User)

	// TODO 获取path中的id
	Id := ctx.Params.ByName("id")

	file := model.ZipFile{}

	// TODO 查看前端文件是否在数据库中存在
	if f.DB.Where("id = ?", Id).First(&file).Error != nil {
		response.Fail(ctx, nil, "文件不存在")
		return
	}

	// TODO 查看是否有权限
	if file.UserId != user.ID && (file.Visible == 3 || (file.Visible == 2 && !util.IsS(4, "Fr"+strconv.Itoa(int(user.ID)), strconv.Itoa(int(file.UserId))))) {
		response.Fail(ctx, nil, "权限不足，不能查看")
		return
	}

	// TODO 记录的总条数
	var total int64
	// TODO 查找图片
	var fileImgs []model.FileImg
	f.DB.Where("file_id = ?", Id).Order("created_at desc").Find(&fileImgs).Count(&total)

	// TODO 返回数据
	response.Success(ctx, gin.H{"fileImgs": fileImgs, "total": total}, "成功")
}

// @title    NewFileController
// @description   新建一个前端文件的控制器
// @auth      MGAronya（张健）       2022-9-16 12:15
// @param    ctx *gin.Context       接收一个上下文
// @return   void
func NewFileController() IFileController {
	db := common.GetDB()
	db.AutoMigrate(model.ZipFile{})
	return FileController{DB: db}
}
