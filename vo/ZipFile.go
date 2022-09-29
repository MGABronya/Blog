// @Title  ZipFile
// @Description  接收更新前端请求时的文件信息
// @Author  MGAronya（张健）
// @Update  MGAronya（张健）  2022-9-16 0:47
package vo

// ZipFileRequest			接收更新前端请求时的文件信息
type ZipFileRequest struct {
	Title    string `json:"title"`     // 标题
	Remark   string `json:"remark"`    // 备注
	ResLong  string `json:"res_long"`  // 备用长文本
	ResShort string `json:"res_short"` // 备用短文本
}
