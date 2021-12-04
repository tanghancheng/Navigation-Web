package models

import (
	"Navigation-Web/dao"
	"gorm.io/gorm"
)

type NavigationInfo struct {
	gorm.Model
	Title string `json:"title"`
	Url   string `json:"url"`
	Logo  string `json:"Logo"`
}

var db *gorm.DB

func init() {
	db = dao.InitMysqlConn()
}
func (n *NavigationInfo) GetAll() (navigationInfos []NavigationInfo, err error) {
	db.AutoMigrate(&NavigationInfo{})
	result := db.Find(&navigationInfos)
	if result.Error != nil {
		return nil, err
	}
	return navigationInfos, nil
}
