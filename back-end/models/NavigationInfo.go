package models

import (
	"Navigation-Web/dao"
	"gorm.io/gorm"
	"log"
)

type NavigationInfo struct {
	gorm.Model
	Title string `json:"title"`
	Url   string `json:"url"`
	Logo  string `json:"logo"`
}

func GetAll() (navigationInfos []NavigationInfo, err error) {
	result := dao.DB.Find(&navigationInfos)
	if result.Error != nil {
		return nil, err
	}
	return navigationInfos, nil
}
func GetOne(id int) (navigationInfo NavigationInfo, err error) {
	log.Println(id)
	result := dao.DB.First(&navigationInfo, id)
	if result.Error != nil {
		return navigationInfo, err
	}
	return navigationInfo, nil
}

func Create(info *NavigationInfo) (err error) {
	result := dao.DB.Create(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func Update(info *NavigationInfo) (err error) {
	result := dao.DB.Save(&info)
	if result.Error != nil {
		return err
	}
	return nil
}
func Delete(info *NavigationInfo) (err error) {
	result := dao.DB.Delete(info)
	if result.Error != nil {
		return err
	}
	return nil
}
