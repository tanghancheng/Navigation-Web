package dto

import (
	"Navigation-Web/dao"

	"gorm.io/gorm"
)

type PageInfo struct {
	PageSize int         `form:"pagesize" json:"pagesize" ` //每页条数
	PageNum  int         `form:"pagenum" json:"pagenum" `   //当前页
	Total    int64       `json:"total"`                     //总条数
	Data     interface{} `json:"data"`                      //分页数据
}

func NewPageInfo() *PageInfo {
	pageinfo := PageInfo{
		PageSize: 6,
		PageNum:  1,
	}
	return &pageinfo
}
func (pageinfo *PageInfo) GetPageDB() (db *gorm.DB) {
	return dao.DB.Limit(int(pageinfo.PageSize)).
		Offset((int(pageinfo.PageNum) - 1) * int(pageinfo.PageSize))
}
